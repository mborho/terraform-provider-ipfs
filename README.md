terraform-provider-ipfs
========================
This provider supports Terraform 0.12.x and later. It expects a running IPFS node on the local machine.

[IPFS pinning service API ](https://ipfs.github.io/pinning-services-api-spec/) is implemented, though no vendor support at the moment.

## Requirements


- [Terraform](https://www.terraform.io/downloads.html) 0.12+



## Building the provider


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

See [docs/index.md](./docs/index.md) for provider setup.

### Resources:

* [ipfs_add](./docs/resources/add.md)
* [ipfs_dir](./docs/resources/dir.md)
* [ipfs_file](./docs/resources/file.md)
* [ipfs_key](./docs/resources/key.md)
* [ipfs_pin](./docs/resources/pin.md)
* [ipfs_remote_pin](./docs/resources/remote-pin.md)
* [ipfs_swarm_connect](./docs/resources/swarm-connect.md)
* [ipfs_publish](./docs/resources/publish.md)

### Guides

* [Add DNS entry at Cloudflare](./docs/guides/dns.md)
* [Publish an object as IPNS name.](./docs/guides/publish.md)

## License

[Mozilla Public License 2.0](./LICENSE)


