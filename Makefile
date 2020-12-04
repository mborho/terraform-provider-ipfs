build:
	echo "Building binaries for provider..."
	echo "Use 'make v1.1.1' to build specific version."
	GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64/terraform-provider-ipfs
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin_amd64/terraform-provider-ipfs
	GOOS=windows GOARCH=amd64 go build -o bin/windows_amd64/terraform-provider-ipfs

v%:
	echo "Building binaries for provider in version $@"
	GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64/terraform-provider-ipfs_$@
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin_amd64/terraform-provider-ipfs_$@
	GOOS=windows GOARCH=amd64 go build -o bin/windows_amd64/terraform-provider-ipfs_$@

