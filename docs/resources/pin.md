# ipfs_pin

Pin objects to local storage.

-> Objects linked to by the specified `cid` will get recursively pinned.

## Example Usage

```hcl
resource "ipfs_pin" "example" {
    cid = "Qm..."
}
```

## Argument Reference

* `cid` - Content identifier to be pinned.
