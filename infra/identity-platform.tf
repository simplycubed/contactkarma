# Configuration
# https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/identity_platform_default_supported_idp_config

# resource "google_identity_platform_default_supported_idp_config" "facebook" {
#   enabled       = true
#   idp_id        = "facebook.com"
#   client_id     = var.facebook_client_id
#   client_secret = var.facebook_client_secret
# }

resource "google_identity_platform_default_supported_idp_config" "google" {
  enabled       = true
  idp_id        = "google.com"
  client_id     = var.google_client_id
  client_secret = var.google_client_secret
}

# resource "google_identity_platform_default_supported_idp_config" "linkedin" {
#   enabled       = true
#   idp_id        = "linkedin.com"
#   client_id     = var.linkedin_client_id
#   client_secret = var.linkedin_client_secret
# }

resource "google_identity_platform_tenant" "tenant" {
  display_name             = "tenant"
  allow_password_signup    = true
  enable_email_link_signin = true
}

# resource "google_identity_platform_default_supported_idp_config" "twitter" {
#   enabled       = true
#   idp_id        = "gtwitteroogle.com"
#   client_id     = var.twitter_client_id
#   client_secret = var.twitter_client_secret
# }
