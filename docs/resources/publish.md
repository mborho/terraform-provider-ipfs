# ipfs_publish

Publish IPNS names.

## Example Usage

```hcl
resource "ipfs_publish" {
    cid = "Qm..."
    key = "my-new-key"
}
```

## Argument Reference

* `cid` - IPFS path of the object to be published.
* `key` - Name of the key to be used or a valid PeerID, default is `self`.

## Attribute Reference


* `path` - Published IPFS path.
* `name` - Name under the content was published, `/ipns/...`
* `value` - Published IPFS path.
