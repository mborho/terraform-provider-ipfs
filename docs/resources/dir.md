# ipfs_dir

Add a directory to IPFS.

## Example Usage

```hcl
resource "ipfs_dir" {
    path = "./path/to/directory/"
}
```

## Argument Reference

* `path` - Path of the directory to be added.

## Attribute Reference

* `cid` - Content identifier of the added content.
