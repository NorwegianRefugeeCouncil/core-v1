resource "azurerm_linux_web_app" "app" {
  provider                  = azurerm.runtime
  location                  = var.location
  name                      = "${var.app_name}-${var.environment}-${random_id.app_id.hex}"
  resource_group_name       = azurerm_resource_group.rg.name
  service_plan_id           = azurerm_service_plan.plan.id
  https_only                = true
  virtual_network_subnet_id = azurerm_subnet.runtime_subnet.id
  identity {
    type         = "SystemAssigned, UserAssigned"
    identity_ids = [azurerm_user_assigned_identity.app.id]
  }
  site_config {
    ftps_state                                    = "Disabled"
    container_registry_use_managed_identity       = true
    container_registry_managed_identity_client_id = azurerm_user_assigned_identity.app.client_id
    remote_debugging_enabled                      = false
    websockets_enabled                            = false
    http2_enabled                                 = true
    application_stack {
      docker_image     = var.container_image
      docker_image_tag = var.container_image_tag
    }
    app_command_line = "serve"
  }
  app_settings = {
    # See https://learn.microsoft.com/en-us/azure/app-service/configure-custom-container?pivots=container-linux#configure-port-number
    WEBSITES_PORT               = var.port
    oidc_AUTHENTICATION_SECRET  = var.oidc_client_secret
    CORE_DB_DSN                 = "postgres://${random_pet.postgres_admin_username.id}:${random_password.postgres_admin_password.result}@${azurerm_postgresql_flexible_server.postgres.fqdn}/${azurerm_postgresql_flexible_server_database.db.name}?sslmode=require"
    CORE_DB_DRIVER              = "postgres"
    CORE_LISTEN_ADDRESS         = ":${var.port}"
    # See https://learn.microsoft.com/en-us/azure/app-service/configure-authentication-customize-sign-in-out?source=recommendations#sign-out-of-a-session
    CORE_LOGOUT_URL             = "https://${var.app_name}-${var.environment}-${random_id.app_id.hex}.azurewebsites.net/.auth/logout"
    CORE_LOGIN_URL              = "https://${var.app_name}-${var.environment}-${random_id.app_id.hex}.azurewebsites.net/.auth/login/oidc"
    CORE_JWT_GLOBAL_ADMIN_GROUP = var.jwt_global_admin_group
    CORE_AUTH_HEADER_NAME       = "x-ms-token-oidc-id-token"
    CORE_AUTH_HEADER_FORMAT     = "jwt"
    DOCKER_CUSTOM_IMAGE_NAME    = "${var.container_image}:${var.container_image_tag}"
    CORE_OIDC_ISSUER            = var.oidc_issuer_url
    CORE_OAUTH_CLIENT_ID        = var.oidc_client_id
    CORE_TOKEN_REFRESH_INTERVAL = "15m"
    CORE_TOKEN_REFRESH_URL      = "https://${var.app_name}-${var.environment}-${random_id.app_id.hex}.azurewebsites.net/.auth/refresh"
    CORE_LOG_LEVEL              = var.log_level
  }
  sticky_settings {
    app_setting_names = [
      "oidc_AUTHENTICATION_SECRET",
    ]
  }
  depends_on = [
    azurerm_role_assignment.acr_pull
  ]
}

# AzureRM provider does not yet support the authsettingsv2.
# For now, we rely on the azapi provider to update the
# app service with proper oidc settings.
resource azapi_update_resource app_auth {
  provider    = azapi.runtime
  type        = "Microsoft.Web/sites/config@2022-03-01"
  resource_id = "${azurerm_linux_web_app.app.id}/config/web"
  body        = jsonencode({
    properties = {
      siteAuthSettingsV2 = {
        platform = {
          enabled        = true
          runtimeVersion = "~1"
        }
        globalValidation = {
          requireAuthentication       = true
          unauthenticatedClientAction = "RedirectToLoginPage"
          redirectToProvider          = "oidc"
        }
        identityProviders = {
          customOpenIdConnectProviders = {
            oidc = {
              registration = {
                clientId         = var.oidc_client_id
                clientCredential = {
                  clientSecretSettingName = "oidc_AUTHENTICATION_SECRET"
                }
                openIdConnectConfiguration = {
                  wellKnownOpenIdConfiguration = var.oidc_well_known_url
                }
              },
              login = {
                nameClaimType = "name"
                scopes        = [
                  "openid",
                  "profile",
                  "email",
                  "groups",
                  "offline_access",
                ]
              }
            }
          }
        }
        login = {
          tokenStore = {
            enabled                    = true
            tokenRefreshExtensionHours = 72
          }
        }
        httpSettings = {
          requireHttps = true
        }
      }
    }
  })
}

