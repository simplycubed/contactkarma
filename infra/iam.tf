data "google_project" "project" {
  provider = google-beta
}

locals {
  google_contacts_permissions = toset([
    "roles/iam.serviceAccountTokenCreator",
    "roles/iam.serviceAccountUser",
    "roles/run.invoker"
  ])
  api_permissions = toset([
    "roles/datastore.user",
    "roles/pubsub.publisher"
  ])
  jobs_permissions = toset([
    "roles/datastore.user",
    "roles/pubsub.publisher"
  ])
  gw_permissions = toset([
    "roles/iam.serviceAccountTokenCreator",
    "roles/run.invoker"
  ])
}

# API - Cloud Run
resource "google_service_account" "api_sa" {
  project      = var.project_id
  account_id   = "api-sa"
  display_name = "api-sa"
}

resource "google_project_iam_member" "api_sa" {
  for_each = local.api_permissions
  project  = var.project_id
  role     = each.key
  member   = "serviceAccount:${google_service_account.api_sa.email}"
}

# API Invoker - API Gateway
resource "google_service_account" "api_gw" {
  project      = var.project_id
  account_id   = "apiinvoker"
  display_name = "apiinvoker"
}

resource "google_project_iam_member" "api_gw" {
  for_each = local.gw_permissions
  project  = var.project_id
  role     = each.key
  member   = "serviceAccount:${google_service_account.api_gw.email}"
}

# Contacts Jobs - Cloud Run
resource "google_service_account" "contacts_jobs_sa" {
  project      = var.project_id
  account_id   = "contacts-jobs-sa"
  display_name = "contacts-jobs-sa"
}

resource "google_project_iam_member" "contacts_jobs_sa" {
  for_each = local.jobs_permissions
  project  = var.project_id
  role     = each.key
  member   = "serviceAccount:${google_service_account.contacts_jobs_sa.email}"
}

resource "google_kms_crypto_key_iam_member" "contacts_jobs_sa" {
  crypto_key_id = google_kms_crypto_key.people.id
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:${google_service_account.contacts_jobs_sa.email}"
}

# Cloud Build
resource "google_project_iam_member" "cloudbuild" {
  project = var.project_id
  role    = "roles/secretmanager.secretAccessor"
  member  = "serviceAccount:${data.google_project.project.number}@cloudbuild.gserviceaccount.com"
}

# Compute - Cloud Run
resource "google_project_iam_member" "compute" {
  project = var.project_id
  role    = "roles/secretmanager.secretAccessor"
  member  = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
}

# Functions
resource "google_project_iam_member" "functions" {
  project = var.project_id
  role    = "roles/secretmanager.secretAccessor"
  member  = "serviceAccount:${var.project_id}@appspot.gserviceaccount.com"
}

# Google Contacts PubSub
resource "google_service_account" "contacts" {
  account_id   = "google-contacts"
  display_name = "Google Contacts"
}

resource "google_project_iam_member" "contacts" {
  for_each = local.google_contacts_permissions
  project  = var.project_id
  role     = each.key
  member   = "serviceAccount:${google_service_account.contacts.email}"
}

# PUSH
# Subscription SA can subscribe to topic
resource "google_pubsub_topic_iam_member" "push_contacts" {
  topic  = google_pubsub_topic.push_contact_source.id
  role   = "roles/pubsub.subscriber"
  member = "serviceAccount:${google_service_account.contacts.email}"
}

# Default PubSub SA publish permissions on DL topic
resource "google_pubsub_topic_iam_member" "push_contact_source_dl_pub" {
  topic  = google_pubsub_topic.push_contact_source_dead_letter.id
  role   = "roles/pubsub.publisher"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

# Default PubSub SA subscribe permissions on push_contacts_subscription
resource "google_pubsub_subscription_iam_member" "push_contact_source_dl_sub" {
  subscription = google_pubsub_subscription.push_contact_source.id
  role         = "roles/pubsub.subscriber"
  member       = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

# PULL
# Subscription SA can subscribe to topic
resource "google_pubsub_topic_iam_member" "pull_contacts" {
  topic  = google_pubsub_topic.pull_contacts.id
  role   = "roles/pubsub.subscriber"
  member = "serviceAccount:${google_service_account.contacts.email}"
}

# Default PubSub SA publish permissions on DL topic
resource "google_pubsub_topic_iam_member" "pull_contacts_dl_pub" {
  topic  = google_pubsub_topic.pull_contacts_dead_letter.id
  role   = "roles/pubsub.publisher"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

# Default PubSub SA subscribe permissions on pull_contacts subscription
resource "google_pubsub_subscription_iam_member" "pull_contacts_sub" {
  subscription = google_pubsub_subscription.pull_contacts.id
  role         = "roles/pubsub.subscriber"
  member       = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

# Default PubSub SA publish permissions on DL topic
resource "google_pubsub_topic_iam_member" "pull_contact_source_dl_pub" {
  topic  = google_pubsub_topic.pull_contact_source_dead_letter.id
  role   = "roles/pubsub.publisher"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

# Default PubSub SA subscribe permissions on pull_contacts subscription
resource "google_pubsub_subscription_iam_member" "pull_contact_source_sub" {
  subscription = google_pubsub_subscription.pull_contact_source.id
  role         = "roles/pubsub.subscriber"
  member       = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}
