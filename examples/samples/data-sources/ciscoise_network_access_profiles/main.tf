terraform {
  required_providers {
    ciscoise = {
      version = "0.6.4-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_network_access_profiles" "example" {
  provider = ciscoise
}

output "ciscoise_network_access_profiles_example" {
  value = data.ciscoise_network_access_profiles.example.items
}
