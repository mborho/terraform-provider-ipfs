terraform-provider-ipfs
========================
This provider supports Terraform 0.12.x and later. It expects a running IPFS node on the local machine.

[IPFS pinning service API ](https://ipfs.github.io/pinning-services-api-spec/) is implemented, though no vendor support at the moment.

## Requirements


- [Terraform](https://www.terraform.io/downloads.html) 0.12+



## Building the provider


```sh
$ git clone git@github.com:mborho/terraform-provider-ipfs.git
$ cd terraform-provider-ipfs
$ go install
```

If **terraform init** can't find the provider, copy the installed binary for your system in one of these 3 places:

1. in the standard user plugin-dir for terraform: 
   * Linux/Mac:	**~/.terraform.d/plugins** 
   * Windows: **%APPDATA%\terraform.d\plugins**
2. in your local terraform project folder under the directory **.terraform/plugins/**
3. in **/usr/local/bin/** or somewhere else in your *$PATH*.

See [terraform.io/docs/configuration/providers.html#third-party-plugins](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins) for more infos.


## Resources

### Provider

```hcl
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

**Arguments:**

* **node** - Server address of the IPFS node, default is **localhost:5001**
* **remote_pin_service** - Configuration of remote pinning service, can be defined *multiple* times. 
  * **name** - Identifier name of pinning service, unique for this provider.
  * **endpoint** - API endpoint of remote pinning service.
  * **token** - Token for authentication.
  * **skip_ssl_verify** - *boolean* Skip SSL verification, default is **false**

### ipfs_add

```hcl
resource "ipfs_add" {
    path = "./path/to/file"
}
```

**Arguments:**

* **path** - Path to the file to be added.

**Attributes:**

* **cid** - Content identifier of the added content.


### ipfs_dir

```hcl
resource "ipfs_dir" {
    path = "./path/to/directory/"
}
```

**Arguments:**

* **path** - Path to the file to be added.

**Attributes:**

* **cid** - Content identifier of the added content.


### ipfs_file

```hcl
resource "ipfs_file" "example" {
    file = "./local/path/to/file.txt"
    path = "/ipfs-unixfs-path/filename.txt"
}
```

**Arguments:**

* **file** - Path to the file to be added.
* **path** - Path in the IPFS local filespace.

**Attributes:**

* **cid** - Content identifier of the added content.


### ipfs_pin

```hcl
resource "ipfs_pin" "example" {
    cid = "Qm..."
}
```

**Arguments:**

* **cid** - Content identifier of the content to be pinned.

### ipfs_remote_pin

```hcl
resource "ipfs_remote_pin" "example" {   
  service = "vendor-name-from-provider"
  cid     = ipfs_file.example.cid
  name    = "name.txt" 
  origins = [
    "/ip6/2a03:b0c0:3:d0::3281:e001/udp/4001/quic/p2p/12D3KooWNJGCBznrEnRngbvoE1gPzoW8sdiNE3kB1mQXYndzHYuP",
    "/ip4/139.59.141.250/udp/4001/quic/p2p/12D3KooWNJGCBznrEnRngbvoE1gPzoW8sdiNE3kB1mQXYndzHYuP"
  ]     
  meta = {
    foo = "bla"
    baz = "baz"
  }
}
```

**Arguments:**

* **service** - Name of the service, same as in provider setup.
* **cid** - Content identifier of the content to be pinned.
* **name** - Name of the content to be pinned.
* **origins** - List of multi-addresses for service to grab content from.
* **meta** - Map of meta informations to be saved at service.

**Attributes:**

* **request_id** - Id of the pin at the pinning service.
* **status** - Status of the pin at the pinning service.
* **delegates** - List of pinning services nodes to connect to.

*info* data from service not supported by now.

### ipfs_swarm_connect

```hcl
resource "ipfs_swarm_connect" "test" {
  addresses = ipfs_remote_pin.example.delegates
  can_fail  = true   # fail gracefully, no error when connect times out.
}  
```
**Arguments:**

* **origins** - List of multi-addresses for IPFS node to connect.
* **can_fail** - *boolean*, connection request can fail gracefully, **true** is default.

### ipfs_key

```hcl
resource "ipfs_key" {
    name = "new-key-name"
    type = "rsa|ed25519" // default rsa
    size = 2048  // default
}
```
**Arguments:**

* **name** - Name of the key.
* **type** - Type of key, default, is **rsa**.
* **size** - Size of key, default is **2048**.

### ipfs_publish

```hcl
resource "ipfs_publish" {
    cid = "Qm..."
    key = "my-new-key" // default 'self'/Node-ID
}
```

**Arguments:**

* **cid** - Content identifier of the content to be published.
* **key** - Name of the key under which the content will be published, default, is **self**.

**Attributes:**

* **path** - Published IPFS path.
* **name** - Name under the content was published, **/ipns/...**
* **value** - Published IPFS path.



## Usage Example

```hcl

provider "ipfs" {}

// Add directory
resource "ipfs_dir" "demo" {
  path = "./path/to/dir/"
}

// Pin directory
resource "ipfs_pin" "demo_dir" {
  cid = ipfs_dir.demo.cid
}

// create key
resource "ipfs_key" "demo" {
  name = "demo-key"
}

// publish directory with own key (ipns)
resource "ipfs_publish" "demo" {
  cid = ipfs_dir.demo.cid
  key = ipfs_key.demo.name
}
```

###  Add DNS entry at cloudflare

```hcl
provider "cloudflare" {
  email   = "ipfs-dns@example.com"
  api_key = "ABCDEF0123456789"
}

data "cloudflare_zones" "demo" {
  filter {
    name   = "example.com"
    status = "active"
    paused = false
  }
}

resource "cloudflare_record" "demo" {
  zone_id = lookup(data.cloudflare_zones.demo.zones[0], "id")
  name    = "ipfs-demo"
  value   = "www.cloudflare-ipfs.com"
  type    = "CNAME"
  ttl     = 600
}

resource "cloudflare_record" "demo_dnslink" {
  zone_id = lookup(data.cloudflare_zones.demo.zones[0], "id")
  name    = "_dnslink.ipfs-demo"
  value   = "dnslink=/ipns/${ipfs_publish.demo.name}"
  type    = "TXT"
  ttl     = 600
}
```

