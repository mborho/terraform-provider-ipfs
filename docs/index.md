# IPFS Provider

This provider supports Terraform 0.12.x and later. It expects a running IPFS node on the local machine.

[IPFS pinning service API ](https://ipfs.github.io/pinning-services-api-spec/) is implemented, though no vendor support at the moment.

## Example Usage

```hcl
terraform {
  required_providers {
    ipfs = {
      source  = "mborho/ipfs"
      version = "=> 0.1.0"
    }
}

provider "ipfs" {
    node = "<http address of ipfs node, default is localhost:5001>"

    remote_pin_service {
      name            = "dev"
      endpoint        = "https://pinning-service-api.example.com/api/v1"
      token           = var.pinning_api_token
      skip_ssl_verify = false
  }
}
```

## Argument Reference

* **node** *string* - Server address of the IPFS node, default is **localhost:5001**
* **remote_pin_service** *map* - Configuration of remote pinning service, can be defined *multiple* times. 
  * **name** *string* - Identifier name of pinning service, unique for this provider.
  * **endpoint** *string* - API endpoint of remote pinning service.
  * **token** *string* - Token for authentication.
  * **skip_ssl_verify** *bool* - Skip SSL verification, default is **false**

## Environment Variables

**HTTP_PROXY** is supported.
