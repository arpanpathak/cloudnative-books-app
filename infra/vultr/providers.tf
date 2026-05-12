terraform {
  required_providers {
    vultr = {
      source  = "vultr/vultr"
      version = "~> 2.19"
    }
  }
}

provider "vultr" {
  # The API Key is automatically picked up from the VULTR_API_KEY environment variable.
  # Security Best Practice: Never hardcode your API key in Terraform files.
}
