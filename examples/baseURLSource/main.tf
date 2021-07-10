terraform {
  required_providers {
    fram = {
      version = "0.1"
      source  = "darkedges.com/forgerock/fram"
    }
  }
}

data "fram_baseurlsource" "all" {}

# Returns Base URL Source
output "baseurlsource" {
  value = data.fram_baseurlsource.all
}
