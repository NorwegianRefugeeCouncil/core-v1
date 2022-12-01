terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.33"
    }
    azapi = {
      source  = "Azure/azapi"
      version = "~> 1.1"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.4"
    }
  }
  backend "azurerm" {
  }
}

provider "azapi" {
  alias = "runtime"
}

provider "azurerm" {
  alias           = "runtime"
  subscription_id = var.subscription_id
  features {
  }
}

provider "azurerm" {
  alias           = "infra"
  subscription_id = var.infra_subscription_id
  features {
  }
}

