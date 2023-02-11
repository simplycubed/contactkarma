module "app" {
  count                 = var.env == "dev" ? 1 : 0
  source                = "simplycubed/cloud-run-cicd-updates/google"
  version               = "2.2.2"
  name                  = "app"
  image                 = "gcr.io/${var.project_id}/${var.default_image}"
  location              = var.region
  map_domains           = ["app.${var.base_domain}"]
  allow_public_access   = true
  ingress               = "all"
  service_account_email = google_service_account.api_sa.email
}

module "contacts" {
  source                = "simplycubed/cloud-run-cicd-updates/google"
  version               = "2.2.2"
  name                  = "contacts"
  image                 = "gcr.io/${var.project_id}/${var.default_image}"
  location              = var.region
  map_domains           = ["contacts.${var.base_domain}"]
  allow_public_access   = false
  ingress               = "all"
  service_account_email = google_service_account.api_sa.email
  env = [
    { key = "ALLOWED_ORIGIN", value = "*" },
    { key = "ENV", value = var.env },
    { key = "FIREBASE_URL", value = "${var.project_id}.firebaseapp.com" },
    { key = "GOOGLE_CLOUD_PROJECT", value = var.project_id },
    { key = "FRONTEND_URL", value = "https://app.${var.base_domain}" },
    { key = "GOOGLE_AUTH_CLIENT_ID", value = var.google_client_id },
    { key = "GOOGLE_AUTH_CLIENT_SECRET", value = var.google_client_secret },
    { key = "GOOGLE_AUTH_REDIRECT_URL", value = "https://app.${var.base_domain}/contact-source/connect" },
    { key = "PUSH_CONTACT_SOURCE_TOPIC", value = "push-contact-source" },
    { key = "PULL_CONTACT_SOURCE_TOPIC", value = "pull-contact-source" },
    { key = "PULL_CONTACTS_TOPIC", value = "pull-contacts" },
    { key = "CONTACTS_KMS_KEY", value = "projects/${var.project_id}/locations/global/keyRings/google/cryptoKeys/people" },
    { key = "TYPESENSE_API_KEY", value = var.typesense_api_key },
    { key = "TYPESENSE_HOST", value = var.typesense_host }
  ]
}

module "contacts_jobs" {
  source                = "simplycubed/cloud-run-cicd-updates/google"
  version               = "2.2.2"
  name                  = "contacts-jobs"
  image                 = "gcr.io/${var.project_id}/${var.default_image}"
  entrypoint            = ["/app"]
  args                  = ["jobs"]
  location              = var.region
  map_domains           = ["contacts-jobs.${var.base_domain}"]
  allow_public_access   = false
  ingress               = "all"
  service_account_email = google_service_account.contacts_jobs_sa.email
  timeout               = 600
  env = [
    { key = "ALLOWED_ORIGIN", value = "*" },
    { key = "ENV", value = var.env },
    { key = "JOB_PORT", value = "8080" },
    { key = "FIREBASE_URL", value = "${var.project_id}.firebaseapp.com" },
    { key = "GOOGLE_CLOUD_PROJECT", value = var.project_id },
    { key = "FRONTEND_URL", value = "https://app.${var.base_domain}" },
    { key = "GOOGLE_AUTH_CLIENT_ID", value = var.google_client_id },
    { key = "GOOGLE_AUTH_CLIENT_SECRET", value = var.google_client_secret },
    { key = "GOOGLE_AUTH_REDIRECT_URL", value = "https://app.${var.base_domain}/contact-source/connect" },
    { key = "PUSH_CONTACT_SOURCE_TOPIC", value = "push-contact-source" },
    { key = "PULL_CONTACT_SOURCE_TOPIC", value = "pull-contact-source" },
    { key = "PULL_CONTACTS_TOPIC", value = "pull-contacts" },
    { key = "CONTACTS_KMS_KEY", value = "projects/${var.project_id}/locations/global/keyRings/google/cryptoKeys/people" }
  ]
}

module "options_service" {
  source                = "simplycubed/cloud-run-cicd-updates/google"
  version               = "2.2.2"
  name                  = "options-service"
  image                 = "gcr.io/${var.project_id}/${var.default_image}"
  location              = var.region
  map_domains           = ["options-service.${var.base_domain}"]
  allow_public_access   = true
  ingress               = "all"
  service_account_email = google_service_account.api_sa.email
  env = [
    { key = "ALLOWED_ORIGIN", value = "*" },
    { key = "ENV", value = var.env },
    { key = "GOOGLE_CLOUD_PROJECT", value = var.project_id }
  ]
}

module "web" {
  count                 = var.env == "dev" ? 1 : 0
  source                = "simplycubed/cloud-run-cicd-updates/google"
  version               = "2.2.2"
  name                  = "web"
  image                 = "gcr.io/${var.project_id}/${var.default_image}"
  location              = var.region
  map_domains           = ["${var.base_domain}"]
  allow_public_access   = true
  ingress               = "all"
  service_account_email = google_service_account.api_sa.email
}
