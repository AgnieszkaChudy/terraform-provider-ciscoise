terraform {
  required_providers {
    ciscoise = {
      version = "0.6.22-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_system_config_version" "example" {
  provider = ciscoise
}

output "ciscoise_system_config_version_example" {
  value = data.ciscoise_system_config_version.example.item
}
