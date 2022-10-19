resource azurerm_management_lock deletion_lock {
  provider   = azurerm.runtime
  count      = var.prevent_deletion ? 1 : 0
  name       = azurerm_resource_group.rg.name
  lock_level = "CanNotDelete"
  scope      = azurerm_resource_group.rg.id
}