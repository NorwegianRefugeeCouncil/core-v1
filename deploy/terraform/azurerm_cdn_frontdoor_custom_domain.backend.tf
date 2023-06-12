resource "azurerm_cdn_frontdoor_custom_domain" "backend" {
  provider                 = azurerm.runtime
  name                     = "backend"
  cdn_frontdoor_profile_id = azurerm_cdn_frontdoor_profile.fd.id
  dns_zone_id              = data.azurerm_dns_zone.dns.id
  host_name                = var.backend_host_name

  tls {
    certificate_type    = "ManagedCertificate"
    minimum_tls_version = "TLS12"
  }
}

resource "azurerm_dns_txt_record" "backend" {
  name                = "_dnsauth.${var.environment}-${var.app_name}"
  zone_name           = data.azurerm_dns_zone.dns.name
  resource_group_name = var.dns_resource_group_name
  ttl                 = 3600

  record {
    value = azurerm_cdn_frontdoor_custom_domain.backend.validation_token
  }
}

resource "azurerm_dns_cname_record" "backend" {
  depends_on = [azurerm_cdn_frontdoor_route.backend]

  name                = "${var.environment}-${var.app_name}" 
  zone_name           = data.azurerm_dns_zone.dns.name
  resource_group_name = var.dns_resource_group_name
  ttl                 = 3600
  record              = azurerm_cdn_frontdoor_endpoint.backend.host_name
}