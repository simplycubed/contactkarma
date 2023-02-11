# Pull Contacts
resource "google_pubsub_topic" "pull_contacts" {
  name = "pull-contacts"
}

resource "google_pubsub_topic" "pull_contacts_dead_letter" {
  name                       = "pull-contacts-dead-letter"
  message_retention_duration = "604800s"
}

resource "google_pubsub_subscription" "pull_contacts" {
  name  = "pull-contacts"
  topic = google_pubsub_topic.pull_contacts.name
  push_config {
    push_endpoint = "${module.contacts_jobs.url}/pull-contacts"
    attributes = {
      x-goog-version = "v1"
    }
    oidc_token {
      service_account_email = google_service_account.contacts.email
    }
  }
  enable_exactly_once_delivery = true
  expiration_policy {
    ttl = "" // never expires
  }
  dead_letter_policy {
    dead_letter_topic     = google_pubsub_topic.pull_contacts_dead_letter.id
    max_delivery_attempts = 100
  }
  retry_policy {
    minimum_backoff = "10s"
    maximum_backoff = "600s"
  }
}

resource "google_pubsub_subscription" "pull_contacts_dead_letter" {
  name  = "pull-contacts-dead-letter"
  topic = google_pubsub_topic.pull_contacts_dead_letter.name
  expiration_policy {
    ttl = "" // never expires
  }
}

# Pull Contacts Job Scheduler
resource "google_cloud_scheduler_job" "job" {
  name        = "pull-contacts-job"
  description = "Pull Contacts Job"
  schedule    = "0 0 * * *"

  pubsub_target {
    topic_name = google_pubsub_topic.pull_contacts.id
    attributes = {
      x-goog-version = "v1"
    }
  }
}

# Pull Contact Source
resource "google_pubsub_topic" "pull_contact_source" {
  name = "pull-contact-source"
}

resource "google_pubsub_topic" "pull_contact_source_dead_letter" {
  name                       = "pull-contact-source-dead-letter"
  message_retention_duration = "604800s"
}

resource "google_pubsub_subscription" "pull_contact_source" {
  name  = "pull-contact-source"
  topic = google_pubsub_topic.pull_contact_source.name
  push_config {
    push_endpoint = "${module.contacts_jobs.url}/pull-contact-source"
    attributes = {
      x-goog-version = "v1"
    }
    oidc_token {
      service_account_email = google_service_account.contacts.email
    }
  }
  enable_exactly_once_delivery = true
  expiration_policy {
    ttl = "" // never expires
  }
  dead_letter_policy {
    dead_letter_topic     = google_pubsub_topic.pull_contact_source_dead_letter.id
    max_delivery_attempts = 100
  }
  retry_policy {
    minimum_backoff = "10s"
    maximum_backoff = "600s"
  }
}

resource "google_pubsub_subscription" "pull_contact_source_dead_letter" {
  name  = "pull-contact-source-dead-letter-subscription"
  topic = google_pubsub_topic.pull_contact_source_dead_letter.name
  expiration_policy {
    ttl = "" // never expires
  }
}

# Push Contact Source
resource "google_pubsub_topic" "push_contact_source" {
  name = "push-contact-source"
}

resource "google_pubsub_topic" "push_contact_source_dead_letter" {
  name                       = "push-contact-source-dead-letter"
  message_retention_duration = "604800s"
}

resource "google_pubsub_subscription" "push_contact_source" {
  name  = "push-contact-source"
  topic = google_pubsub_topic.push_contact_source.name
  push_config {
    push_endpoint = "${module.contacts_jobs.url}/push-contact-source"
    attributes = {
      x-goog-version = "v1"
    }
    oidc_token {
      service_account_email = google_service_account.contacts.email
    }
  }
  enable_exactly_once_delivery = true
  expiration_policy {
    ttl = "" // never expires
  }
  dead_letter_policy {
    dead_letter_topic     = google_pubsub_topic.push_contact_source_dead_letter.id
    max_delivery_attempts = 100
  }
  retry_policy {
    minimum_backoff = "10s"
    maximum_backoff = "600s"
  }
}

resource "google_pubsub_subscription" "push_contact_source_dead_letter" {
  name  = "push-contact-source-dead-letter"
  topic = google_pubsub_topic.push_contact_source_dead_letter.name
  expiration_policy {
    ttl = "" // never expires
  }
}
