---
page_title: "Publish an object as IPNS name."
---

# Publish an object as IPNS name.


```hcl

provider "ipfs" {}

// Add directory
resource "ipfs_dir" "demo" {
  path = "./path/to/dir/"
}

// Pin directory
resource "ipfs_pin" "demo_dir" {
  cid = ipfs_dir.demo.cid
}

// Create own key, don't use the default key from the node.
resource "ipfs_key" "demo" {
  name = "demo-key"
}

// publish directory with own key (ipns)
resource "ipfs_publish" "demo" {
  cid = ipfs_dir.demo.cid
  key = ipfs_key.demo.name
}
```
