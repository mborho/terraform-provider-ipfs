terraform-provider-ipfs
========================

![release](https://github.com/mborho/terraform-provider-ipfs/workflows/release/badge.svg)

This provider supports Terraform 0.13.x and later. *(0.12.x if building manually*. 

It expects a running IPFS node on the local machine.

[IPFS pinning service API ](https://ipfs.github.io/pinning-services-api-spec/) is implemented, though no vendor support at the moment.

## Building the provider

*Requirements*: [Terraform](https://www.terraform.io/downloads.html) 0.12+


```sh
$ git clone git@github.com:mborho/terraform-provider-ipfs.git
$ cd terraform-provider-ipfs
$ go install
```

If **terraform init** can't find the provider, copy the installed binary for your system in one of these 3 places:

1. in the standard user plugin-dir for terraform: 
   * Linux/Mac:	**~/.terraform.d/plugins** 
   * Windows: **%APPDATA%\terraform.d\plugins**
2. in your local terraform project folder under the directory **.terraform/plugins/**
3. in **/usr/local/bin/** or somewhere else in your *$PATH*.

See [terraform.io/docs/configuration/providers.html#third-party-plugins](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins) for more infos.

## Documentation

See [registry.terraform.io/providers/mborho/ipfs/latest/docs](https://registry.terraform.io/providers/mborho/ipfs/latest/docs) for documentation and provider setup.

### Resources:

* [ipfs_add](https://registry.terraform.io/providers/mborho/ipfs/latest/docs/resources/add)
* [ipfs_dir](https://registry.terraform.io/providers/mborho/ipfs/latest/docs/resources/dir)
* [ipfs_file](https://registry.terraform.io/providers/mborho/ipfs/latest/docs/resources/file)
* [ipfs_key](https://registry.terraform.io/providers/mborho/ipfs/latest/docs/resources/key)
* [ipfs_pin](https://registry.terraform.io/providers/mborho/ipfs/latest/docs/resources/pin)
* [ipfs_remote_pin](https://registry.terraform.io/providers/mborho/ipfs/latest/docs/resources/remote-pin)
* [ipfs_swarm_connect](https://registry.terraform.io/providers/mborho/ipfs/latest/docs/resources/swarm-connect)
* [ipfs_publish](https://registry.terraform.io/providers/mborho/ipfs/latest/docs/resources/publish)

### Guides

* [Add DNS entry at Cloudflare](https://registry.terraform.io/providers/mborho/ipfs/latest/docs/guides/dns)
* [Publish an object as IPNS name.](https://registry.terraform.io/providers/mborho/ipfs/latest/docs/guides/publish)

## License

[Mozilla Public License 2.0](./LICENSE)
