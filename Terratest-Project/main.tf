provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "example" {
  name     = var.resource_group_name
  location = var.location
  
}

output "resource_group_name" {
  value = azurerm_resource_group.example.name
}

output "azurerm_resource_group_location" {
  value = azurerm_resource_group.example.location
}

resource "azurerm_virtual_network" "example1" {
  name                = var.resource_name1
  location            = var.location
  resource_group_name = azurerm_resource_group.example.name
  address_space       = ["10.0.0.0/16"]
  dns_servers         = ["10.0.0.4", "10.0.0.5"]
  subnet {
    name           = "subnet1"
    address_prefix = "10.0.1.0/24"
  }

  subnet {
    name           = "subnet2"
    address_prefix = "10.0.2.0/24"
  }

  tags = {
    environment = "Production"
  }
}

output "azurerm_virtual_network_sgn" {
    value = azurerm_virtual_network.example1.resource_group_name
}

output "azurerm_virtual_network_name" {
    value = azurerm_virtual_network.example1.name
}

output "azurerm_virtual_network_id" {
  value = azurerm_virtual_network.example1.id
}
