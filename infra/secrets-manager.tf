# ACTIVECAMPAIGN API KEY
resource "google_secret_manager_secret" "ac_token" {
  secret_id = "ac-token"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "ac_token_version" {
  secret      = google_secret_manager_secret.ac_token.id
  secret_data = var.ac_token
}

# ACTIVECAMPAIGN BASE URL
resource "google_secret_manager_secret" "ac_base_url" {
  secret_id = "ac-base-url"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "ac_base_url_version" {
  secret      = google_secret_manager_secret.ac_base_url.id
  secret_data = var.ac_base_url
}
