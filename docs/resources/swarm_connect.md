# ipfs_swarm_connect

Opens new direct connections to a list of given peer addresses.

The address format is an IPFS multiaddr: `/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ`


## Example Usage

```hcl
resource "ipfs_swarm_connect" "test" {
  addresses = [
      "/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ"
  ]
  can_fail  = true
}
```

## Argument Reference

* `addresses` - List of IPFS multi-addresses to connect to.
* `can_fail` - Connection requests can fail gracefully, `true` is default.
