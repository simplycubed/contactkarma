locals {
  github_verification = var.env == "prod" ? "7c3820522b" : "120f7bdad4"
}

module "dns" {
  source     = "terraform-google-modules/cloud-dns/google"
  type       = "public"
  project_id = var.project_id
  name       = "site-dns"
  domain     = "${var.base_domain}."
  recordsets = [
    {
      name    = "www"
      type    = "A"
      ttl     = 60
      records = var.env == "dev" ? [module.lb-web[0].external_ip, ] : ["199.36.158.100"]
    },
    {
      name    = ""
      type    = "A"
      ttl     = 60
      records = var.env == "dev" ? [module.lb-web[0].external_ip, ] : ["199.36.158.100"]
    },
    {
      name    = "app"
      type    = "A"
      ttl     = 60
      records = var.env == "dev" ? [module.lb-app[0].external_ip, ] : ["199.36.158.100"]
    },
    {
      name = "api"
      type = "A"
      ttl  = 60
      records = [
        module.lb-http.external_ip,
      ]
    },
    {
      name = "contacts"
      type = "CNAME"
      ttl  = 60
      records = [
        "ghs.googlehosted.com.",
      ]
    },
    {
      name = "contacts-jobs"
      type = "CNAME"
      ttl  = 60
      records = [
        "ghs.googlehosted.com.",
      ]
    },
    {
      name = "options-service"
      type = "CNAME"
      ttl  = 60
      records = [
        "ghs.googlehosted.com.",
      ]
    },
    {
      name = ""
      type = "MX"
      ttl  = 300
      records = [
        "1 aspmx.l.google.com.",
        "5 alt1.aspmx.l.google.com.",
        "5 alt2.aspmx.l.google.com.",
        "10 alt3.aspmx.l.google.com.",
        "10 alt4.aspmx.l.google.com."
      ]
    },
    {
      name = "google._domainkey"
      type = "TXT"
      ttl  = 300
      records = [
        "\"v=DKIM1; k=rsa; p=MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoWuaYBkXd7ad+Onk68625A0X2jbRoorD3YMDs2CrE7eU0x122/PakAUZKJkCDdyyFnwdmZnOrSRMilChYTmEi/fSL//InqZAk/gjIAj3tcs4PN2HupJsH5qz654Q6nvhR/mIQVlZxfJxFexDGJaGlvW8pjLq80JaGgb9I4WuIRP1CKB74Y+xc/cz5Q7EqeX4I\" \"k7SZsWhvPP1PycH/ZBpXn0iaHlaY2sPer3QPdOeAiKw5Mn6PkjekivmDYDWOrazWBKVupaRDo6JlVQf2d1rqd1hWs23mwVQeFsrbv1FTf2Lf0QR0uGXt9OjkvgX3JZyadldMsFbx7l0HxS3Gsls+wIDAQAB\""
      ]
    },
    {
      name = "_github-challenge-simplycubed-org"
      type = "TXT"
      ttl  = 300
      records = [
        local.github_verification
      ]
    },
    {
      name = ""
      type = "TXT"
      ttl  = 300
      records = [
        "\"v=spf1 include:_spf.google.com include:_spf.firebasemail.com include:sendgrid.net ~all\""
      ]
    },
    {
      name = ""
      type = "SPF"
      ttl  = 300
      records = [
        "\"v=spf1 include:_spf.google.com include:_spf.firebasemail.com include:sendgrid.net ~all\""
      ]
    },
    {
      name = "_dmarc"
      type = "TXT"
      ttl  = 300
      records = [
        "\"v=DMARC1;\" \"p=quarantine;\" \"rua=mailto:dmarc@${var.base_domain};\" \"ruf=mailto:dmarc@${var.base_domain};\" \"fo=1:d:s\""
      ]
    },
    {
      name = "25830355"
      type = "CNAME"
      ttl  = 60
      records = [
        "sendgrid.net.",
      ]
    },
    {
      name = "em3263"
      type = "CNAME"
      ttl  = 60
      records = [
        "u25830355.wl106.sendgrid.net.",
      ]
    },
    {
      name = "s1._domainkey"
      type = "CNAME"
      ttl  = 60
      records = [
        "s1.domainkey.u25830355.wl106.sendgrid.net.",
      ]
    },
    {
      name = "s2._domainkey"
      type = "CNAME"
      ttl  = 60
      records = [
        "s2.domainkey.u25830355.wl106.sendgrid.net.",
      ]
    },
    {
      name = "url8408"
      type = "CNAME"
      ttl  = 60
      records = [
        "sendgrid.net.",
      ]
    },
  ]
}
