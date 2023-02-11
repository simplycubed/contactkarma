resource "google_kms_key_ring" "google" {
  name     = "google"
  location = "global"
}

resource "google_kms_crypto_key" "people" {
  name            = "people"
  key_ring        = google_kms_key_ring.google.id
  rotation_period = "100000s"

  lifecycle {
    prevent_destroy = false
  }
}
