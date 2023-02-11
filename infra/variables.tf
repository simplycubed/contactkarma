
variable "ac_token" {
  description = "ActiveCampaign Token"
  sensitive   = true
  type        = string
}

variable "ac_base_url" {
  description = "ActiveCampaign Base URL"
  type        = string
}

variable "base_domain" {
  description = "Base domain for DNS records"
  type        = string
}

variable "cypress_app_key" {
  description = "Cypress Dashboard Record Key"
  type        = string
}

variable "cypress_web_key" {
  description = "Cypress Dashboard Record Key"
  type        = string
}

variable "cypress_test_user_email" {
  description = "Cypress Test User Email"
  type        = string
}

variable "cypress_test_user_pass" {
  description = "Cypress Test User Password"
  type        = string
}

variable "cypress_test_user_new_pass" {
  description = "Cypress Test User New Password"
  type        = string
}

# Base image, will update it with other images in Cloud Build
variable "default_image" {
  description = "Default Cloud Run image used to create instance using Terraform"
  type        = string
  default     = null
}

variable "env" {
  description = "Environment"
  type        = string
}

variable "firebase_app_id" {
  description = "Firebase App ID"
  type        = string
}

variable "firebase_api_key" {
  description = "Firebase API Key"
  type        = string
}

variable "firebase_auth_domain" {
  description = "Firebase Auth Domain"
  type        = string
}

variable "firebase_storage_bucket" {
  description = "Firebase Storage Bucket"
  type        = string
}

variable "firebase_messaging_sender_id" {
  description = "Firebase Messaging Sender ID"
  type        = string
}

variable "firebase_measurement_id" {
  description = "Firebase Measurement ID"
  type        = string
}

# CANNOT SET TO SENSITIVE
variable "google_client_id" {
  description = "Google Client ID"
  type        = string
  sensitive   = false
}

# CANNOT SET TO SENSITIVE
variable "google_client_secret" {
  description = "Google Client Secret"
  type        = string
  sensitive   = false
}

variable "iap_brand_name" {
  description = "OAuth IAP brand name for gke endpoints"
}

variable "iap_domain" {
  description = "Domain used for the environment"
}

variable "project_id" {
  description = "Project ID"
  type        = string
}

variable "region" {
  description = "Region for gcloud resources"
  type        = string
}

variable "stripe_pk" {
  description = "Stripe Private Key"
  type        = string
  sensitive   = false
}

# CANNOT SET TO SENSITIVE
variable "typesense_api_key" {
  description = "Typesense API Key"
  type        = string
  sensitive   = false
}

# CANNOT SET TO SENSITIVE
variable "typesense_host" {
  description = "Typesense Host"
  type        = string
  sensitive   = false
}
