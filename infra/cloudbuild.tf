#
# app
resource "google_cloudbuild_trigger" "deploy_app" {
  name = "deploy-app"
  github {
    owner = "simplycubed"
    name  = "app"
    push {
      tag    = var.env == "prod" ? "^production-v\\d+\\.\\d+\\.\\d+$" : null
      branch = var.env == "dev" ? "^main$" : null
    }
  }
  substitutions = {
    _ENV                          = var.env
    _FIREBASE_API_KEY             = var.firebase_api_key
    _FIREBASE_APP_ID              = var.firebase_app_id
    _FIREBASE_AUTH_DOMAIN         = var.firebase_auth_domain
    _FIREBASE_API_URL             = "https://api.${var.base_domain}"
    _FIREBASE_APP_URL             = "https://app.${var.base_domain}"
    _FIREBASE_MEASUREMENT_ID      = var.firebase_measurement_id
    _FIREBASE_MESSAGING_SENDER_ID = var.firebase_messaging_sender_id
    _FIREBASE_STORAGE_BUCKET      = var.firebase_storage_bucket
    _FIREBASE_WEB_URL             = "https://${var.base_domain}"
    _STRIPE_PK                    = var.stripe_pk
  }
  filename = var.env == "prod" ? "cloudbuild.main.yaml" : "cloudbuild.dev.yaml"
  tags     = ["managed by terraform"]
}

resource "google_cloudbuild_trigger" "build_app" {
  count = var.env == "prod" ? 0 : 1
  name  = "build-app"
  github {
    owner = "simplycubed"
    name  = "app"
    pull_request {
      branch = ".*"
    }
  }
  substitutions = {
    _CYPRESS_RECORD_KEY           = var.cypress_app_key
    _ENV                          = var.env
    _FIREBASE_API_KEY             = var.firebase_api_key
    _FIREBASE_APP_ID              = var.firebase_app_id
    _FIREBASE_AUTH_DOMAIN         = var.firebase_auth_domain
    _FIREBASE_API_URL             = "https://api.${var.base_domain}"
    _FIREBASE_APP_URL             = "https://app.${var.base_domain}"
    _FIREBASE_MEASUREMENT_ID      = var.firebase_measurement_id
    _FIREBASE_MESSAGING_SENDER_ID = var.firebase_messaging_sender_id
    _FIREBASE_STORAGE_BUCKET      = var.firebase_storage_bucket
    _FIREBASE_WEB_URL             = "https://${var.base_domain}"
    _STRIPE_PK                    = var.stripe_pk
    _CYPRESS_TEST_USER_EMAIL      = var.cypress_test_user_email
    _CYPRESS_TEST_USER_PASS       = var.cypress_test_user_pass
    _CYPRESS_TEST_USER_NEW_PASS   = var.cypress_test_user_new_pass
  }
  filename = "cloudbuild.pr.yaml"
  tags     = ["managed by terraform"]
}

#
# contacts
#
resource "google_cloudbuild_trigger" "deploy_contacts" {
  name = "deploy-contacts"
  github {
    owner = "simplycubed"
    name  = "contacts"
    push {
      tag    = var.env == "prod" ? "^production-v\\d+\\.\\d+\\.\\d+$" : null
      branch = var.env == "dev" ? "^main$" : null
    }
  }
  substitutions = {
    _ENV = var.env
  }
  filename = "cloudbuild.main.yaml"
  tags     = ["managed by terraform"]
}

resource "google_cloudbuild_trigger" "build_contacts" {
  count = var.env == "prod" ? 0 : 1
  name  = "build-contacts"
  github {
    owner = "simplycubed"
    name  = "contacts"
    pull_request {
      branch = ".*"
    }
  }
  substitutions = {
    _ENV = var.env
  }
  filename = "cloudbuild.pr.yaml"
  tags     = ["managed by terraform"]
}

#
# cypress
#
resource "google_cloudbuild_trigger" "push_cypress_base_image" {
  name = "push-cypress-base-image"
  github {
    owner = "simplycubed"
    name  = "cypress"
    push {
      branch = "^main$"
    }
  }
  filename = "cloudbuild.yaml"
  tags     = ["managed by terraform"]
}

#
# docker-compose
#
resource "google_cloudbuild_trigger" "push_docker_compose_base_image" {
  name = "push-docker-compose-base-image"
  github {
    owner = "simplycubed"
    name  = "docker-compose"
    push {
      branch = "^main$"
    }
  }
  filename = "cloudbuild.yaml"
  tags     = ["managed by terraform"]
}

#
# firebase
#
resource "google_cloudbuild_trigger" "push_firebase_base_image" {
  name = "push-firebase-base-image"
  github {
    owner = "simplycubed"
    name  = "firebase"
    push {
      branch = "^main$"
    }
  }
  filename = "cloudbuild.yaml"
  tags     = ["managed by terraform"]
}

#
# firebase-emulators
#
resource "google_cloudbuild_trigger" "push_firebase_emulators_base_image" {
  name = "push-firebase-emulators-base-image"
  github {
    owner = "simplycubed"
    name  = "firebase-emulators"
    push {
      branch = "^main$"
    }
  }
  filename = "cloudbuild.yaml"
  tags     = ["managed by terraform"]
}

#
# golang
#
resource "google_cloudbuild_trigger" "push_golang_base_image" {
  name = "push-golang-base-image"
  github {
    owner = "simplycubed"
    name  = "golang"
    push {
      branch = "^main$"
    }
  }
  filename = "cloudbuild.yaml"
  tags     = ["managed by terraform"]
}

#
# integrations
#
resource "google_cloudbuild_trigger" "deploy_integrations" {
  name = "deploy-integrations"
  github {
    owner = "simplycubed"
    name  = "integrations"
    push {
      tag    = var.env == "prod" ? "^production-v\\d+\\.\\d+\\.\\d+$" : null
      branch = var.env == "dev" ? "^main$" : null
    }
  }
  substitutions = {
    _ENV               = var.env
    _TYPESENSE_API_KEY = var.typesense_api_key
    _TYPESENSE_HOST    = var.typesense_host
  }
  filename = "cloudbuild.yaml"
  tags     = ["managed by terraform"]
}

#
# options-service
#
resource "google_cloudbuild_trigger" "build_options_service" {
  count = var.env == "prod" ? 0 : 1
  name  = "build-options-service"
  github {
    owner = "simplycubed"
    name  = "options-service"
    pull_request {
      branch = ".*"
    }
  }
  substitutions = {
    _ENV = var.env
  }
  filename = "cloudbuild.pr.yaml"
  tags     = ["managed by terraform"]
}

resource "google_cloudbuild_trigger" "deploy_options_service" {
  name = "deploy-options-service"
  github {
    owner = "simplycubed"
    name  = "options-service"
    push {
      tag    = var.env == "prod" ? "^production-v\\d+\\.\\d+\\.\\d+$" : null
      branch = var.env == "dev" ? "^main$" : null
    }
  }
  substitutions = {
    _ENV = var.env
  }
  filename = "cloudbuild.main.yaml"
  tags     = ["managed by terraform"]
}

#
# nginx
#
resource "google_cloudbuild_trigger" "push_nginx_base_image" {
  name = "push-nginx-base-image"
  github {
    owner = "simplycubed"
    name  = "nginx"
    push {
      branch = "^main$"
    }
  }
  filename = "cloudbuild.yaml"
  tags     = ["managed by terraform"]
}

#
# node
#
resource "google_cloudbuild_trigger" "push_node_base_image" {
  name = "push-node-base-image"
  github {
    owner = "simplycubed"
    name  = "node"
    push {
      branch = "^main$"
    }
  }
  filename = "cloudbuild.yaml"
  tags     = ["managed by terraform"]
}

#
# sample-service
#
resource "google_cloudbuild_trigger" "build_sample_service" {
  name = "build-sample-service"
  github {
    owner = "simplycubed"
    name  = "sample-service"
    pull_request {
      branch = "^main$"
    }
  }
  substitutions = {
    _ENV = var.env
  }
  filename = "cloudbuild.yaml"
  tags     = ["managed by terraform"]
}

#
# web
resource "google_cloudbuild_trigger" "deploy_web" {
  name = "deploy-web"
  github {
    owner = "simplycubed"
    name  = "web"
    push {
      tag    = var.env == "prod" ? "^production-v\\d+\\.\\d+\\.\\d+$" : null
      branch = var.env == "dev" ? "^main$" : null
    }
  }
  substitutions = {
    _ENV                          = var.env
    _FIREBASE_API_KEY             = var.firebase_api_key
    _FIREBASE_APP_ID              = var.firebase_app_id
    _FIREBASE_AUTH_DOMAIN         = var.firebase_auth_domain
    _FIREBASE_API_URL             = "https://api.${var.base_domain}"
    _FIREBASE_APP_URL             = "https://app.${var.base_domain}"
    _FIREBASE_MEASUREMENT_ID      = var.firebase_measurement_id
    _FIREBASE_MESSAGING_SENDER_ID = var.firebase_messaging_sender_id
    _FIREBASE_STORAGE_BUCKET      = var.firebase_storage_bucket
    _FIREBASE_WEB_URL             = "https://${var.base_domain}"
  }
  filename = var.env == "prod" ? "cloudbuild.main.yaml" : "cloudbuild.dev.yaml"
  tags     = ["managed by terraform"]
}

resource "google_cloudbuild_trigger" "build_web" {
  count = var.env == "prod" ? 0 : 1
  name  = "build-web"
  github {
    owner = "simplycubed"
    name  = "web"
    pull_request {
      branch = ".*"
    }
  }
  substitutions = {
    _CYPRESS_RECORD_KEY           = var.cypress_web_key
    _ENV                          = var.env
    _FIREBASE_API_KEY             = var.firebase_api_key
    _FIREBASE_APP_ID              = var.firebase_app_id
    _FIREBASE_AUTH_DOMAIN         = var.firebase_auth_domain
    _FIREBASE_API_URL             = "https://api.${var.base_domain}"
    _FIREBASE_APP_URL             = "https://app.${var.base_domain}"
    _FIREBASE_MEASUREMENT_ID      = var.firebase_measurement_id
    _FIREBASE_MESSAGING_SENDER_ID = var.firebase_messaging_sender_id
    _FIREBASE_STORAGE_BUCKET      = var.firebase_storage_bucket
    _FIREBASE_WEB_URL             = "https://${var.base_domain}"
  }
  filename = "cloudbuild.pr.yaml"
  tags     = ["managed by terraform"]
}
