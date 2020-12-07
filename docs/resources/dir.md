# ipfs_dir

Add a directory to IPFS.

## Example Usage

```hcl
resource "ipfs_dir" {
    local_path = "./path/to/directory/"
}
```

## Argument Reference

* `local_path` - Local filesystem path of the directory to be added.

## Attribute Reference

* `cid` - Content identifier of the added content.
