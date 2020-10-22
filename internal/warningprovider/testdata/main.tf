terraform {
  required_providers {
    corner = {
      source = "hashicorp/corner-warning"
    }
  }
}

data "corner_warning_only" "test" {
}

resource "corner_warning_only" "test" {
  set_id = "foo"
}
