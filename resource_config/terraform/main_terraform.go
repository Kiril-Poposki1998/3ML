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
`

var AWS_Additional = `
provider "aws" {
  region     = "us-west-2"
  access_key = "my-access-key"
  secret_key = "my-secret-key"
}
`

var Azure_Additional = `
provider "azurerm" {
  resource_provider_registrations = "none"
  features {}
}
`

var GCP_additional = `
provider "google" {
  project     = "my-project-id"
  region      = "us-central1"
}
`
