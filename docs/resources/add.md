# ipfs_add

Add a file to IPFS.

## Example Usage

```hcl
resource "ipfs_add" {
    local_path = "./path/to/file"
}
```

## Argument Reference


* `local_path` - local filesystem path to be added to ipfs.

## Attribute Reference

* `cid` - Content identifier of the added content.
