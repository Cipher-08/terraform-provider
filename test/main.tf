terraform {
  required_providers {
    yourprovider = {
      source = "github.com/Cipher-08/yourprovider"
      version = "0.1.0"
    }
  }
}

provider "yourprovider" {}

data "yourprovider_data" "example" {
  input = "Hello, OpenTofu!"
}

output "result" {
  value = data.yourprovider_data.example.output
}
