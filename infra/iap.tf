resource "google_iap_web_iam_member" "access_iap_policy" {
  provider = google-beta
  project  = var.project_id
  role     = "roles/iap.httpsResourceAccessor"
  member   = "group:project-team@${var.iap_domain}"
}

resource "google_iap_client" "iap_client" {
  provider     = google-beta
  display_name = "IAP Auth"
  brand        = var.iap_brand_name
}
