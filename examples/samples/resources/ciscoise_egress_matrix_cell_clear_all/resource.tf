terraform {
  required_providers {
    ciscoise = {
      version = "0.6.20-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}
resource "ciscoise_egress_matrix_cell_clear_all" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {

  }
}