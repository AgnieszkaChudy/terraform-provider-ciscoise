terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.3"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_device_administration_service_names" "example" {
  provider = ciscoise
}

output "ciscoise_device_administration_service_names_example" {
  value = data.ciscoise_device_administration_service_names.example.items
}
