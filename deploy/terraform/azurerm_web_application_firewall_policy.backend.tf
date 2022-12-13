resource "azurerm_web_application_firewall_policy" "backend" {
  name                = "waf${random_id.app_id.hex}"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location

  policy_settings {
    file_upload_limit_in_mb     = 500
    max_request_body_size_in_kb = 128
  }

  managed_rules {
    managed_rule_set {
      version = "3.2"
      type    = "OWASP"
    }
  }
}
