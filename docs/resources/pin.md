# ipfs_pin

Pin objects to local storage.

## Example Usage

```hcl
resource "ipfs_pin" "example" {
    cid = "Qm..."
}
```

## Argument Reference

* `cid` - Content identifier to be pinned.
