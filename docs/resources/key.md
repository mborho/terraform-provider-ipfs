# ipfs_key

Create a new keypair.

## Example Usage

```hcl
resource "ipfs_key" {
    name = "new-key-name"
    type = "rsa"
    size = 2048
}
```

## Argument Reference

* `name` *string* - Name of the key.
* `type` - Type of key. Can be `rsa` or `ed25519`, default is `rsa`.
* `size` - Size of key, default is `2048`.
