---
page_title: "Add DNS entry at Cloudflare."
---

#  Add DNS entry at Cloudflare.

You can use a DNS record to always point to the latest version of your content. You can then use HTTP domains to make your content easy available.

Cloudflare is only used as an example, any DNS provider can used. 

-> See [DNSLink](https://docs.ipfs.io/concepts/dnslink/) for a more detailed description.

```hcl
provider "cloudflare" {
  email   = "ipfs-dns@example.com"
  api_key = "ABCDEF0123456789"
}

data "cloudflare_zones" "demo" {
  filter {
    name   = "example.com"
    status = "active"
    paused = false
  }
}

resource "cloudflare_record" "demo" {
  zone_id = lookup(data.cloudflare_zones.demo.zones[0], "id")
  name    = "ipfs-demo"
  value   = "www.cloudflare-ipfs.com"
  type    = "CNAME"
  ttl     = 600
}

resource "cloudflare_record" "demo_dnslink" {
  zone_id = lookup(data.cloudflare_zones.demo.zones[0], "id")
  name    = "_dnslink.ipfs-demo"
  value   = "dnslink=/ipns/${ipfs_publish.demo.name}"
  type    = "TXT"
  ttl     = 600
}
```


