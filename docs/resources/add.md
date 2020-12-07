# ipfs_add

Add a file to IPFS.

## Example Usage

```hcl
resource "ipfs_add" {
    path = "./path/to/file"
}
```

## Argument Reference


* `path` - The path to a file to be added to ipfs.

## Attribute Reference

* `cid` - Content identifier of the added content.
