terraform {
  required_providers {
    fram = {
      version = "0.1"
      source  = "darkedges.com/forgerock/fram"
    }
  }
}

provider "fram" {
  base_url = "https://fram.darkedges.com/openam"
  username = "amadmin"
  password = "Passw0rd"
  realm    = "/root"
}

module "psl" {
  source = "./baseurlsource"
}

output "psl" {
  value = module.psl.baseurlsource
}

resource "fram_baseurlsource" "test" {
  source = "FIXED_VALUE"
  fixed_value = "https://fram.darkedges.com"
  context_path = "/openam"
}

output "test" {
  value = fram_baseurlsource.test
}