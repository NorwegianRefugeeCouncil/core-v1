resource "azurerm_dns_zone" "dns" {
  provider            = azurerm.infra
  name                = var.infra_dns_zone_name
  resource_group_name = var.infra_resource_group_name
}

resource "azurerm_dns_txt_record" "backend" {
  name                = "_dnsauth.${var.environment}-${var.app_name}"
  zone_name           = azurerm_dns_zone.dns.name
  resource_group_name = var.infra_resource_group_name
  ttl                 = 3600

  record {
    value = azurerm_cdn_frontdoor_custom_domain.backend.validation_token
  }
}

resource "azurerm_dns_cname_record" "backend" {
  depends_on = [azurerm_cdn_frontdoor_route.backend]

  name                = "${var.environment}-${var.app_name}" 
  zone_name           = azurerm_dns_zone.dns.name
  resource_group_name = var.infra_resource_group_name
  ttl                 = 3600
  record              = azurerm_cdn_frontdoor_endpoint.backend.host_name
}