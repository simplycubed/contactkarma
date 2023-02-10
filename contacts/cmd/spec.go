/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/spf13/cobra"
)

type GoogleBackendExtension struct {
	Address         string `json:"address" yml:"address"`
	PathTranslation string `json:"path_translation" yml:"path_translation"`
}

// specCmd represents the spec command
var specCmd = &cobra.Command{
	Use:   "spec",
	Short: "Generate api spec for infra from api.yml",
	Long:  `Generate api spec for infra`,
	Run: func(cmd *cobra.Command, args []string) {

		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		// read api spec
		doc, err := loads.Spec(wd + "/openapi/api.yml")
		if err != nil {
			log.Fatal(err)
		}

		// prepare modifications

		contactGoogleBackend := GoogleBackendExtension{
			Address:         "${contacts_url}",
			PathTranslation: "APPEND_PATH_TO_ADDRESS",
		}
		security := map[string][]string{"firebase": {}}

		// add security definitions and options for cors for infra spec
		orgSpec := doc.OrigSpec()
		for key, path := range orgSpec.Paths.Paths {

			if path.Options == nil {
				// add options
				path.Options = getOptionDefintion(path)
			}
			// add security and backend to other operations
			if path.Get != nil {
				path.Get.AddExtension("x-google-backend", contactGoogleBackend)
				path.Get.Security = []map[string][]string{}
				path.Get.Security = append(path.Get.Security, security)
				modifyFileParam(path.Get)
			}
			if path.Post != nil {
				path.Post.AddExtension("x-google-backend", contactGoogleBackend)
				path.Post.Security = []map[string][]string{}
				path.Post.Security = append(path.Post.Security, security)
				modifyFileParam(path.Post)
			}
			if path.Patch != nil {
				path.Patch.AddExtension("x-google-backend", contactGoogleBackend)
				path.Patch.Security = []map[string][]string{}
				path.Patch.Security = append(path.Patch.Security, security)
				modifyFileParam(path.Patch)
			}
			if path.Put != nil {
				path.Put.AddExtension("x-google-backend", contactGoogleBackend)
				path.Put.Security = []map[string][]string{}
				path.Put.Security = append(path.Put.Security, security)
				modifyFileParam(path.Put)
			}
			if path.Delete != nil {
				path.Delete.AddExtension("x-google-backend", contactGoogleBackend)
				path.Delete.Security = []map[string][]string{}
				path.Delete.Security = append(path.Delete.Security, security)
				modifyFileParam(path.Delete)
			}
			orgSpec.Paths.Paths[key] = path
		}

		raw, err := extendWithSecurityDefinition(orgSpec)
		if err != nil {
			log.Fatal(err)
		}

		data, err := jsonToYml(raw)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(wd+"/api.tpl", data, 0777)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func jsonToYml(raw []byte) (yml []byte, err error) {
	rawObj := map[string]interface{}{}
	err = json.Unmarshal(raw, &rawObj)
	if err != nil {
		log.Println("err", err)
		return
	}
	yml, err = yaml.Marshal(rawObj)
	if err != nil {
		return
	}
	return
}

func extendWithSecurityDefinition(orgSpec *spec.Swagger) (raw []byte, err error) {
	raw, err = orgSpec.MarshalJSON()
	if err != nil {
		return
	}

	rawObj := map[string]json.RawMessage{}
	err = json.Unmarshal(raw, &rawObj)
	if err != nil {
		return
	}

	securityDefinition, err := getSecurityDefinitions()
	if err != nil {
		return
	}

	rawObj["securityDefinitions"] = securityDefinition

	raw, err = json.Marshal(rawObj)
	return
}

// scopes are ommited when marshalled, getSecurityDefinitions return with scopes set explicitly.
func getSecurityDefinitions() (raw json.RawMessage, err error) {
	securityScheme := &spec.SecurityScheme{
		SecuritySchemeProps: spec.SecuritySchemeProps{
			Flow:             "implicit",
			AuthorizationURL: "",
			Scopes:           map[string]string{},
			Type:             "oauth2",
		},
		VendorExtensible: spec.VendorExtensible{
			Extensions: spec.Extensions{
				"x-google-issuer":    "https://securetoken.google.com/${project_name}",
				"x-google-jwks_uri":  "https://www.googleapis.com/service_accounts/v1/metadata/x509/securetoken@system.gserviceaccount.com",
				"x-google-audiences": "${project_name}",
			},
		},
	}

	securitySchemeJson, err := json.Marshal(securityScheme)
	if err != nil {
		return
	}

	securitySchemaJsonMap := map[string]json.RawMessage{}
	err = json.Unmarshal(securitySchemeJson, &securitySchemaJsonMap)
	if err != nil {
		return
	}

	// add empty scope
	emptyObjectJson, _ := json.Marshal(map[string]string{})
	securitySchemaJsonMap["scopes"] = json.RawMessage(emptyObjectJson)

	securityDefinitions := map[string]map[string]json.RawMessage{}
	securityDefinitions["firebase"] = securitySchemaJsonMap

	securityDefinitionsJson, err := json.Marshal(securityDefinitions)
	if err != nil {
		return
	}
	return securityDefinitionsJson, nil
}

func getOptionDefintion(path spec.PathItem) *spec.Operation {
	if path.Get != nil {
		return getOption(path.Get)
	}
	if path.Post != nil {
		return getOption(path.Post)
	}
	if path.Patch != nil {
		return getOption(path.Patch)
	}
	if path.Put != nil {
		return getOption(path.Put)
	}
	if path.Delete != nil {
		return getOption(path.Delete)
	}
	return nil
}

func getOption(operation *spec.Operation) *spec.Operation {
	corsGoogleBackend := GoogleBackendExtension{
		Address:         "${options_url}",
		PathTranslation: "APPEND_PATH_TO_ADDRESS",
	}
	options := spec.NewOperation("cors-" + operation.ID)
	params := []spec.Parameter{}
	for _, param := range operation.Parameters {
		if param.In == "path" || param.In == "header" {
			params = append(params, param)
		}
	}
	options.Parameters = params
	options.AddExtension("x-google-backend", corsGoogleBackend)
	options.RespondsWith(200, spec.NewResponse().WithDescription("A successful response"))
	return options
}

// https://stackoverflow.com/questions/43059304/error-deploying-endpoint-containing-parameter-of-type-file
func modifyFileParam(operation *spec.Operation) {
	for index, param := range operation.Parameters {
		// modify file params
		if param.Type == "file" {
			param.Type = "string"
			param.Format = "binary"
		}
		operation.Parameters[index] = param
	}
}

func init() {
	rootCmd.AddCommand(specCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// specCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// specCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
