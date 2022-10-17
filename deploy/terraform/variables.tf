variable "location" {
  type        = string
  description = "Azure Location for deploying resources"
}

variable "subscription_id" {
  type        = string
  description = "Runtime Subscription id"
}

variable "infra_subscription_id" {
  type        = string
  description = "Infrastructure Subscription id"
}

variable "infra_resource_group_name" {
  type        = string
  description = "Infrastructure Resource Group name"
}

variable "infra_storage_account_name" {
  type        = string
  description = "Storage account name for infrastructure"
}

variable "infra_container_name" {
  type        = string
  description = "Storage Account Container name for infrastructure"
}

variable "infra_container_key" {
  type        = string
  description = "Storage Account Container key for infrastructure"
  default     = "state.tfstate"
}

variable "infra_container_registry_name" {
  type        = string
  description = "Name of the azure container registry"
}

variable "app_name" {
  type        = string
  description = "Name of the application. e.g. core"
}

variable "environment" {
  type        = string
  description = "Name of the environment. e.g. staging"
}

variable "address_space" {
  type        = string
  description = "Address space for the virtual networgitk"
}

variable "runtime_subnet_address_space" {
  type        = string
  description = "Address space for the runtime environment"
}

variable "postgres_subnet_address_space" {
  type        = string
  description = "Address space for the postgres subnet"
}

variable "postgresql_server_name" {
  type        = string
  description = "Unique name for the postgresql server"
}

variable "postgres_sku_name" {
  type        = string
  description = "Name of the postgres SKU"
}
