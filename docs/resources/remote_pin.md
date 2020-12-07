# ipfs_remote_pin

Pin objects to a remote pinning service.

Implements the [IPFS pinning service API Spec](https://ipfs.github.io/pinning-services-api-spec/).

## Example Usage

```hcl
resource "ipfs_remote_pin" "example" {
  service = "vendor-name-from-provider"
  cid     = ipfs_file.example.cid
  name    = "name.txt"
  origins = [
    "/ip6/2a03:b0c0:3:d0::3281:e001/udp/4001/quic/p2p/12D3KooWNJGCBznrEnRngbvoE1gPzoW8sdiNE3kB1mQXYndzHYuP",
    "/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ"
  ]
  meta = {
    foo = "bla"
    baz = "bar"
  }
}
```

## Argument Reference

* `service` - Name of the service, same as in provider setup.
* `cid` - Content identifier of the content to be pinned.
* `name` - Name of the content to be pinned.
* `origins` - List of IPFS multi-addresses for service to grab content from.
* `meta` - Map of meta informations to be saved at service.

## Attribute Reference

* `request_id` - Id of the pin at the pinning service.
* `status` - Status of the pin at the pinning service.
* `delegates` - List of IPFS multi-addresses of pinning services nodes to connect to.

*info* data from service not supported by now.
