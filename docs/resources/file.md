# ipfs_file

Interact with IPFS objects representing Unix filesystems.

## Example Usage

```hcl
resource "ipfs_file" "example" {
    file = "./local/path/to/file.txt"
    path = "/ipfs-unixfs-path/filename.txt"
}
```

## Argument Reference

* `file`- Path to the file to be added.
* `path`- Path in the IPFS local filespace.

## Attribute Reference

* `cid` - Content identifier of the added content.
