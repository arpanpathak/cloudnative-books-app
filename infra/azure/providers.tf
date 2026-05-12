terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.100"
    }
    local = {
      source  = "hashicorp/local"
      version = "~> 2.0"
    }
  }
}

provider "azurerm" {
  # The "features" block is required for AzureRM provider 2.x and 3.x
  features {}
  subscription_id = "83965f97-3369-41e6-aff5-763219a67144"
}
