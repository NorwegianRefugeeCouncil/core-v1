variable "location" {
  type        = string
  description = <<EOF
The location/region where the azure resources for this environment will be created.

To view the available locations, run the following command:
az account list-locations --query "[].{Name:name}" -o table
EOF
}

variable "subscription_id" {
  type        = string
  description = <<EOF
The subscription ID where the azure resources for this environment will be created.
EOF
}

variable "infra_subscription_id" {
  type        = string
  description = <<EOF
The subscription ID where the infrastructure resources are located.

The so-called infrastructure resources are the resources that are
shared across all environments, such as container registries, key vaults, etc.
EOF
}

variable "infra_resource_group_name" {
  type        = string
  description = <<EOF
The name of the resource group where the infrastructure resources are located.

The so-called infrastructure resources are the resources that are
shared across all environments, such as container registries, key vaults, etc.
EOF
}

variable "infra_container_registry_name" {
  type        = string
  description = <<EOF
The name of the shared container registry where the container images are located.
These terraform manifests will allow the azure app service to pull the images
from this registry by creating a role assignment for the app service.
EOF
}

variable "app_name" {
  type        = string
  description = <<EOF
Name of the application.

The name of the application will be used to identify resources that belong to
this application. For example, the name of the resource group, the name of the
app service plan, the name of the app service, etc.
EOF
}

variable "environment" {
  type        = string
  description = <<EOF
The name of the environment.

The name of the environment will be used to identify resources that belong to
this environment. For example, the name of the resource group, the name of the
app service plan, the name of the app service, etc.
EOF
}

variable "address_space" {
  type        = string
  description = <<EOF
The address space of the virtual network in CIDR notation.

The virtual network will be used to host the app service as well as using
subnet delegation to the flexible postgres server.
EOF
}

variable "runtime_subnet_address_space" {
  type        = string
  description = <<EOF
The address space of the runtime subnet in CIDR notation.
The runtime subnet will be used to host the app service.
EOF
}

variable "postgres_subnet_address_space" {
  type        = string
  description = <<EOF
The address space of the postgres subnet in CIDR notation.
The postgres subnet will be used to host the flexible postgres server.
EOF
}

variable "postgres_sku_name" {
  type        = string
  description = <<EOF
The sku name of the flexible postgres server.
To view the available sku names, run the following command:
az postgres flexible-server list-skus -l <location> -o table

Note that the sku name must be prefixed with the tier name.
If the SKU name is "Standard_B1ms" and the tier is "Burstable",
the sku name must be "B_Standard_B1ms".
EOF
}

variable "postgres_geo_redundant_backup_enabled" {
  type        = bool
  default     = false
  description = <<EOF
Whether or not geo-redundant backups are enabled for the flexible postgres server.

If geo-redundant backups are enabled, the flexible postgres server will be
backed up to a secondary region. This will increase the cost of the flexible
postgres server.
EOF
}

variable "postgres_backup_retention_days" {
  type        = number
  default     = 7
  description = <<EOF
The number of days to retain backups for the flexible postgres server.
The minimum value is 7 days.
The maximum value is 35 days.
EOF
  validation {
    condition     = var.postgres_backup_retention_days >= 7 && var.postgres_backup_retention_days <= 35
    error_message = "The number of days to retain backups for the flexible postgres server must be between 7 and 35 days."
  }
}

variable "postgres_enable_high_availability" {
  type        = bool
  default     = false
  description = <<EOF
Whether or not high availability is enabled for the flexible postgres server.

If high availability is enabled, the flexible postgres server will be
hosted in a secondary region. This will increase the cost of the flexible
postgres server.

When this variable is set to true, the variable "postgres_standby_availability_zone"
must also be set.
EOF
}

variable "postgres_availability_zone" {
  type        = string
  default     = "1"
  description = <<EOF
The availability zone of the flexible postgres server.
The availability zone must be either "1", "2", or "3".
EOF
}

variable "postgres_standby_availability_zone" {
  type        = string
  default     = "2"
  description = <<EOF
The availability zone of the flexible postgres server's standby.
The availability zone must be either "1", "2", or "3".

This variable is only meaningful when the variable "postgres_enable_high_availability"
is set to true.
EOF
}

variable "postgres_version" {
  type        = string
  default     = "14"
  description = <<EOF
The version of the flexible postgres server.

To view the available versions see
https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/concepts-supported-versions
EOF
}

variable "postgres_storage_mb" {
  type        = number
  default     = 32768
  description = <<EOF
The storage size of the flexible postgres server in MB.

To view the available storage sizes see
https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/concepts-compute-storage
EOF
}

variable "oidc_client_id" {
  type        = string
  description = <<EOF
The client ID of the OIDC application.

The OIDC application is used to authenticate users to the app service.
The OIDC application must be created beforehand.
EOF
}

variable "oidc_client_secret" {
  type        = string
  description = <<EOF
The client secret of the OIDC application.

The OIDC application is used to authenticate users to the app service.
The OIDC application must be created beforehand.
EOF
}

variable "oidc_well_known_url" {
  type        = string
  description = <<EOF
The well-known URL of the OIDC issuer.

The url is usually in the form of
https://<issuer>/.well-known/openid-configuration
EOF

}

variable "oidc_issuer_url" {
  type        = string
  description = <<EOF
The URL of the OIDC issuer.
EOF
}


variable "container_image" {
  type        = string
  description = <<EOF
The container image of the application.

The container image must be located in the shared container registry, or
in a public container registry.
EOF
}

variable "container_image_tag" {
  type        = string
  description = <<EOF
The tag of the container image of the application.
EOF
}

variable "jwt_global_admin_group" {
  type        = string
  description = <<EOF
The name of the global admin group.

The global admin group is used to grant global admin access to the application.
When users authenticate to the application, the application will check if the
user is a member of the global admin group. If yes, the user will be granted
global admin access to the application.
EOF
}

variable "jwt_can_read_group" {
  type        = string
  description = <<EOF
The name of the can read group.

EOF
}

variable "jwt_can_write_group" {
  type        = string
  description = <<EOF
The name of the can write group.

EOF
}

variable "port" {
  type        = number
  description = <<EOF
The publicly accessible port of the application.
EOF
}

variable "prevent_deletion" {
  type        = bool
  default     = true
  description = <<EOF
Whether or not to prevent the deletion of the resource group.
This is implemented using an Azure Management Lock.
EOF
}

variable "log_level" {
  type        = string
  default     = "info"
  description = <<EOF
The log level of the application.
EOF
  validation {
    condition     = contains(["debug", "info", "warn", "error"], var.log_level)
    error_message = "The log level must be one of \"debug\", \"info\", \"warn\", or \"error\"."
  }
}

variable "service_plan_sku_name" {
  type        = string
  default     = "P1v2"
  description = <<EOF
THe sku name of the app service plan. https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/service_plan#sku_name
EOF
}

variable "dns_zone_name" {
  type        = string
  description = "The name of the dns zone."
}

variable "backend_host_name" {
  type        = string
  description = "The hostname of the backend."
}

variable "frontdoor_sku_name" {
  type        = string
  description = "The sku name of the frontdoor."
  default     = "Standard_AzureFrontDoor"
}

variable "action_group_webhook_url" {
  type        = string
  description = "The url of the webhook for the action group."
}

variable "dns_text_record_value" {
  type        = string
  description = "The value of the DNS txt record"
}