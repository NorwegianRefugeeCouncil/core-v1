resource "azurerm_subnet" "runtime_subnet" {
  provider             = azurerm.runtime
  name                 = "runtime"
  resource_group_name  = azurerm_resource_group.rg.name
  virtual_network_name = azurerm_virtual_network.vnet.name
  address_prefixes     = [var.runtime_subnet_address_space]
  delegation {
    name = "appService"
    service_delegation {
      name = "Microsoft.Web/serverFarms"
      actions = [
        "Microsoft.Network/virtualNetworks/subnets/action",
      ]
    }
  }
}

resource "azurerm_network_security_group" "runtime_nsg" {
  provider            = azurerm.runtime
  name                = "runtime-nsg"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
}

# Outbound traffic to the Azure Instance Metadata Service (IMDS) using its well-known IP address
# This is required to use Managed Identities
resource "azurerm_network_security_rule" "outbound_rule_imds" {
  provider                    = azurerm.runtime
  name                        = "allow-outbound"
  priority                    = 100
  direction                   = "Outbound"
  access                      = "Allow"
  protocol                    = "Tcp"
  source_port_range           = "*"
  destination_port_range      = "80"
  source_address_prefix       = "*"
  destination_address_prefix  = "169.254.169.254"
  network_security_group_name = azurerm_network_security_group.runtime_nsg.name
  resource_group_name         = azurerm_resource_group.rg.name
}

resource "azurerm_subnet_network_security_group_association" "runtume_association" {
  provider                  = azurerm.runtime
  subnet_id                 = azurerm_subnet.runtime_subnet.id
  network_security_group_id = azurerm_network_security_group.runtime_nsg.id
}