// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "contact karma service",
    "title": "Contact Karma Service",
    "version": "1.0.0"
  },
  "paths": {
    "/contact-source-clean-up": {
      "post": {
        "summary": "clean up job after deleting a contact source",
        "operationId": "contact-source-clean-up",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/PubsubMessage"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Triggered",
            "schema": {
              "$ref": "#/definitions/JobSuccess"
            }
          },
          "400": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/JobError"
            }
          }
        }
      }
    },
    "/pull-contact-source": {
      "post": {
        "summary": "sync all contacts from a contact source",
        "operationId": "pull-contact-source",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/PubsubMessage"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Triggered",
            "schema": {
              "$ref": "#/definitions/JobSuccess"
            }
          },
          "400": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/JobError"
            }
          }
        }
      }
    },
    "/pull-contacts": {
      "post": {
        "summary": "publish pull-contact-source for contact-sources",
        "operationId": "pull-contacts",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/PubsubMessage"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Triggered",
            "schema": {
              "$ref": "#/definitions/JobSuccess"
            }
          },
          "400": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/JobError"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ContactSourceDeleted": {
      "type": "object",
      "title": "ContactSourceDeleted",
      "properties": {
        "contactSourceId": {
          "type": "string"
        },
        "removeContactsFromUnified": {
          "type": "boolean"
        },
        "source": {
          "type": "string"
        },
        "userId": {
          "type": "string"
        }
      }
    },
    "JobError": {
      "type": "object",
      "title": "JobError",
      "properties": {
        "description": {
          "type": "string"
        },
        "error": {
          "type": "string"
        }
      }
    },
    "JobSuccess": {
      "type": "object",
      "title": "message",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Message": {
      "type": "object",
      "title": "Message",
      "required": [
        "messageId"
      ],
      "properties": {
        "attributes": {
          "type": "object"
        },
        "data": {
          "type": "string",
          "format": "byte"
        },
        "deliveryAttempt": {
          "type": "integer"
        },
        "messageId": {
          "type": "string"
        },
        "orderingKey": {
          "type": "string"
        },
        "publishTime": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "PubsubMessage": {
      "type": "object",
      "title": "PubsubMessage",
      "properties": {
        "message": {
          "type": "object",
          "$ref": "#/definitions/Message"
        },
        "subscription": {
          "type": "string"
        }
      }
    },
    "PullContactsRequest": {
      "type": "object",
      "title": "PullContactsRequest",
      "properties": {
        "contactSourceId": {
          "type": "string"
        },
        "userId": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "contact karma service",
    "title": "Contact Karma Service",
    "version": "1.0.0"
  },
  "paths": {
    "/contact-source-clean-up": {
      "post": {
        "summary": "clean up job after deleting a contact source",
        "operationId": "contact-source-clean-up",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/PubsubMessage"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Triggered",
            "schema": {
              "$ref": "#/definitions/JobSuccess"
            }
          },
          "400": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/JobError"
            }
          }
        }
      }
    },
    "/pull-contact-source": {
      "post": {
        "summary": "sync all contacts from a contact source",
        "operationId": "pull-contact-source",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/PubsubMessage"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Triggered",
            "schema": {
              "$ref": "#/definitions/JobSuccess"
            }
          },
          "400": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/JobError"
            }
          }
        }
      }
    },
    "/pull-contacts": {
      "post": {
        "summary": "publish pull-contact-source for contact-sources",
        "operationId": "pull-contacts",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/PubsubMessage"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Triggered",
            "schema": {
              "$ref": "#/definitions/JobSuccess"
            }
          },
          "400": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/JobError"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ContactSourceDeleted": {
      "type": "object",
      "title": "ContactSourceDeleted",
      "properties": {
        "contactSourceId": {
          "type": "string"
        },
        "removeContactsFromUnified": {
          "type": "boolean"
        },
        "source": {
          "type": "string"
        },
        "userId": {
          "type": "string"
        }
      }
    },
    "JobError": {
      "type": "object",
      "title": "JobError",
      "properties": {
        "description": {
          "type": "string"
        },
        "error": {
          "type": "string"
        }
      }
    },
    "JobSuccess": {
      "type": "object",
      "title": "message",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Message": {
      "type": "object",
      "title": "Message",
      "required": [
        "messageId"
      ],
      "properties": {
        "attributes": {
          "type": "object"
        },
        "data": {
          "type": "string",
          "format": "byte"
        },
        "deliveryAttempt": {
          "type": "integer"
        },
        "messageId": {
          "type": "string"
        },
        "orderingKey": {
          "type": "string"
        },
        "publishTime": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "PubsubMessage": {
      "type": "object",
      "title": "PubsubMessage",
      "properties": {
        "message": {
          "type": "object",
          "$ref": "#/definitions/Message"
        },
        "subscription": {
          "type": "string"
        }
      }
    },
    "PullContactsRequest": {
      "type": "object",
      "title": "PullContactsRequest",
      "properties": {
        "contactSourceId": {
          "type": "string"
        },
        "userId": {
          "type": "string"
        }
      }
    }
  }
}`))
}
