resource "azurerm_dns_zone" "dns" {
  provider            = azurerm.runtime
  name                = var.dns_zone_name
  resource_group_name = azurerm_resource_group.rg.name
}

resource "azurerm_dns_a_record" "dns_a_record" {
  provider            = azurerm.runtime
  name                = "@"
  zone_name           = azurerm_dns_zone.dns.name
  resource_group_name = azurerm_resource_group.rg.name
  ttl                 = 3600
  target_resource_id  = azurerm_cdn_frontdoor_endpoint.backend.id
}

resource "azurerm_dns_txt_record" "example" {
  name                = "_dnsauth"
  zone_name           = azurerm_dns_zone.dns.name
  resource_group_name = azurerm_resource_group.rg.name
  ttl                 = 3600

  record {
    value = var.dns_text_record_value
  }
}