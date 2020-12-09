# ipfs_publish

Publish IPNS names.

## Example Usage

```hcl
resource "ipfs_publish" {
    path = "Qm..."
    key = "my-new-key"
}
```

## Argument Reference

* `path` - IPFS path of the object to be published. Normal CID identifier will be expanded with */ipfs/* prefix, if missing.
* `key` - Name of the key to be used or a valid PeerID, default is `self`.

## Attribute Reference


* `name` - Name under the content was published, `/ipns/...`
* `value` - Published IPFS path.
