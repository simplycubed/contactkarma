locals {
  api_config_id_prefix     = "api"
  api_gateway_container_id = "api-gw"
  gateway_id               = "gw"
}

resource "google_api_gateway_api" "api_gw" {
  provider     = google-beta
  api_id       = local.api_gateway_container_id
  display_name = "API Gateway"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider             = google-beta
  api                  = google_api_gateway_api.api_gw.api_id
  api_config_id_prefix = local.api_config_id_prefix
  display_name         = "Config"

  openapi_documents {
    document {
      path     = "spec.yaml"
      contents = base64encode(data.template_file.api_config.rendered)
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_gateway" "gw" {
  provider = google-beta
  region   = var.region

  api_config = google_api_gateway_api_config.api_cfg.id

  gateway_id   = local.gateway_id
  display_name = "Gateway"

  depends_on = [google_api_gateway_api_config.api_cfg]
}

data "template_file" "api_config" {
  template = file("${path.module}/openapi/api.tpl")
  vars = {
    contacts_url = module.contacts.url
    options_url  = module.options_service.url
    project_name = "contactkarma-${var.env}"
  }
}

resource "google_api_gateway_api_config_iam_member" "member" {
  provider   = google-beta
  api        = google_api_gateway_api_config.api_cfg.api
  api_config = google_api_gateway_api_config.api_cfg.api_config_id
  role       = "roles/apigateway.viewer"
  member     = "serviceAccount:${google_service_account.api_gw.email}"
}