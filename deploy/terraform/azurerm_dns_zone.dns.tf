resource "azurerm_dns_zone" "dns" {
  provider            = azurerm.runtime
  name                = var.dns_zone_name
  resource_group_name = azurerm_resource_group.rg.name
}

resource "azurerm_dns_a_record" "dns_a_record" {
  name                = "dns-a-record"
  zone_name           = azurerm_dns_zone.dns.name
  resource_group_name = azurerm_resource_group.rg.name
  ttl                 = 3600
  target_resource_id  = azurerm_cdn_frontdoor_origin.backend.id
}