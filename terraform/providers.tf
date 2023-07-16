terraform {
  required_providers {
    scalyr = {
      source  = "terraform.local/local/scalyr"
      version = "1.0.0"
    }
  }
}

provider "scalyr" {
  scalyr_app_url = "https://app.scalyr.com"
  scalyr_api_key = ""
  client_timeout = 60
}
