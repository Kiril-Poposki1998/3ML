package terraform

var Main = `
terraform {
  required_providers {
	{{ .provider_name }} = {
		source = "{{ .remote_repo }}"	
		version = "{{ .provider_version }}"
	}
  }
}

{{- .additional_info -}}
`

var DO_Additional = `
variable "do_token" {}

provider "digitalocean" {
  token = var.do_token
}

resource "digitalocean_droplet" "web" {
  image   = "ubuntu-20-04-x64"
  name    = "web-1"
  region  = "nyc2"
  size    = "s-1vcpu-1gb"
  backups = true
  backup_policy {
    plan    = "weekly"
    weekday = "TUE"
    hour    = 8
  }
}

output "IP" {
  value = digitalocean_droplet.web.ipv4_address
}
`

var AWS_Additional = `
provider "aws" {
  region     = "us-west-2"
  access_key = "my-access-key"
  secret_key = "my-secret-key"
}

resource "aws_ec2_host" "web" {
  instance_type     = "t2.medium"
  availability_zone = "us-west-2a"
  host_recovery     = "on"
  auto_placement    = "on"
}

output "IP" {
  value = aws_ec2_host.web.id
}
`

var Azure_Additional = `
provider "azurerm" {
  resource_provider_registrations = "none"
  features {}
}

resource "azurerm_linux_virtual_machine" "web" {
  name                = "example-machine"
  resource_group_name = azurerm_resource_group.example.name
  location            = azurerm_resource_group.example.location
  size                = "Standard_F2"
  admin_username      = "adminuser"
  network_interface_ids = [
    azurerm_network_interface.example.id,
  ]
}
`

var GCP_additional = `
provider "google" {
  project     = "my-project-id"
  region      = "us-central1"
}

resource "google_compute_instance" "web" {
  name         = "web-instance"
  machine_type = "e2-medium"
  zone         = "us-central1-a"
  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2022-04-lts"
    }
  }
  network_interface {
    network = "default"
    access_config {
    }
  }
}

output "IP" {
  value = google_compute_instance.web.network_interface[0].access_config[0].nat_ip
}
`
