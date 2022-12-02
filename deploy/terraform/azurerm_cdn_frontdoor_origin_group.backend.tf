resource "azurerm_cdn_frontdoor_origin_group" "backend" {
  provider                 = azurerm.runtime
  name                     = "fdog-${random_id.app_id.hex}"
  cdn_frontdoor_profile_id = azurerm_cdn_frontdoor_profile.fd.id
  session_affinity_enabled = false
  load_balancing {
    sample_size                 = 4
    successful_samples_required = 3
  }
}
