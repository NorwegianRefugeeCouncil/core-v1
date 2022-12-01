resource "azurerm_cdn_frontdoor_custom_domain" "backend" {
  provider                 = azurerm.runtime
  name                     = "backend"
  cdn_frontdoor_profile_id = azurerm_cdn_frontdoor_profile.fd.id
  dns_zone_id              = azurerm_dns_zone.dns.id
  host_name                = var.backend_host_name

  tls {
    certificate_type    = "ManagedCertificate"
    minimum_tls_version = "TLS12"
  }
}
