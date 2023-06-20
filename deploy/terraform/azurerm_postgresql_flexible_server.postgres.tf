resource "azurerm_postgresql_flexible_server" "postgres" {
  provider                     = azurerm.runtime
  name                         = random_id.postgresql_id.hex
  resource_group_name          = azurerm_resource_group.rg.name
  location                     = azurerm_resource_group.rg.location
  version                      = var.postgres_version
  delegated_subnet_id          = azurerm_subnet.postgres_subnet.id
  private_dns_zone_id          = azurerm_private_dns_zone.postgres_dns.id
  administrator_login          = random_pet.postgres_admin_username.id
  administrator_password       = random_password.postgres_admin_password.result
  zone                         = var.postgres_availability_zone
  storage_mb                   = var.postgres_storage_mb
  sku_name                     = var.postgres_sku_name
  backup_retention_days        = var.postgres_backup_retention_days
  geo_redundant_backup_enabled = var.postgres_geo_redundant_backup_enabled
  dynamic "high_availability" {
    for_each = var.postgres_enable_high_availability ? [var.postgres_standby_availability_zone] : []
    content {
      mode                      = "ZoneRedundant"
      standby_availability_zone = high_availability.value
    }
  }

  depends_on = [
    azurerm_private_dns_zone_virtual_network_link.vnet_link
  ]

  lifecycle {
    ignore_changes = [
      administrator_password,
      zone,
      high_availability.0.standby_availability_zone,
    ]
  }
}

resource "azurerm_postgresql_flexible_server_configuration" "extensions" {
  provider  = azurerm.runtime
  server_id = azurerm_postgresql_flexible_server.postgres.id
  name      = "azure.extensions"
  value     = "uuid-ossp"
}