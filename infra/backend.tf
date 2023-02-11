terraform {
  cloud {
    organization = "simplycubed"
    hostname     = "app.terraform.io"

    workspaces {
      tags = ["contactkarma"]
    }
  }
}
