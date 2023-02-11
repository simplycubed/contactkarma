resource "google_compute_region_network_endpoint_group" "api_gw" {
  name                  = "api-gw"
  provider              = google-beta
  network_endpoint_type = "SERVERLESS"
  region                = var.region
  serverless_deployment {
    platform = "apigateway.googleapis.com"
    resource = google_api_gateway_gateway.gw.gateway_id
  }

  lifecycle {
    create_before_destroy = true
  }
}

module "lb-http" {
  source  = "./modules/serverless_negs"
  project = var.project_id
  name    = "api-gw"

  ssl                             = true
  managed_ssl_certificate_domains = ["api.${var.base_domain}"]
  https_redirect                  = true

  backends = {
    default = {
      description             = null
      enable_cdn              = false
      custom_request_headers  = ["Access-Control-Request-Method:GET"]
      custom_response_headers = ["Access-Control-Origin:'*'"]
      security_policy         = null
      log_config = {
        enable      = true
        sample_rate = 1.0
      }
      iap_config = {
        enable               = false
        oauth2_client_id     = null
        oauth2_client_secret = null
      }
      groups = [
        {
          group = google_compute_region_network_endpoint_group.api_gw.id
        }
      ]
    }
  }
}

resource "google_compute_region_network_endpoint_group" "web" {
  count                 = var.env == "dev" ? 1 : 0
  name                  = "web"
  network_endpoint_type = "SERVERLESS"
  region                = var.region
  cloud_run {
    service = module.web[0].name
  }
  lifecycle {
    create_before_destroy = true
  }
}

module "lb-web" {
  count   = var.env == "dev" ? 1 : 0
  source  = "./modules/serverless_negs"
  project = var.project_id
  name    = "web"

  ssl                             = true
  managed_ssl_certificate_domains = [var.base_domain]
  https_redirect                  = true

  backends = {
    default = {
      description             = null
      enable_cdn              = false
      custom_request_headers  = ["Access-Control-Request-Method:GET"]
      custom_response_headers = ["Access-Control-Origin:'*'"]
      security_policy         = null
      log_config = {
        enable      = true
        sample_rate = 1.0
      }
      iap_config = {
        enable               = true
        oauth2_client_id     = google_iap_client.iap_client.client_id
        oauth2_client_secret = google_iap_client.iap_client.secret

      }
      groups = [
        {
          group = google_compute_region_network_endpoint_group.web[0].id
        }
      ]
    }
  }
}


resource "google_compute_region_network_endpoint_group" "app" {
  count                 = var.env == "dev" ? 1 : 0
  name                  = "app"
  network_endpoint_type = "SERVERLESS"
  region                = var.region
  cloud_run {
    service = module.app[0].name
  }
  lifecycle {
    create_before_destroy = true
  }
}

module "lb-app" {
  count   = var.env == "dev" ? 1 : 0
  source  = "./modules/serverless_negs"
  project = var.project_id
  name    = "app"

  ssl                             = true
  managed_ssl_certificate_domains = ["app.${var.base_domain}"]
  https_redirect                  = true

  backends = {
    default = {
      description             = null
      enable_cdn              = false
      custom_request_headers  = ["Access-Control-Request-Method:GET"]
      custom_response_headers = ["Access-Control-Origin:'*'"]
      security_policy         = null
      log_config = {
        enable      = true
        sample_rate = 1.0
      }
      iap_config = {
        enable               = true
        oauth2_client_id     = google_iap_client.iap_client.client_id
        oauth2_client_secret = google_iap_client.iap_client.secret

      }
      groups = [
        {
          group = google_compute_region_network_endpoint_group.app[0].id
        }
      ]
    }
  }
}