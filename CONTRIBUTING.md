# Contributing

## Build from source

```bash
git clone git@github.com:kumar1236/terraform-provider-virtualbox.git
cd terraform-provider-virtualbox
go test ./...
go build -o bin/terraform-provider-virtualbox .
```

On Windows, use `bin/terraform-provider-virtualbox.exe` as the output path.

## Test a local build

Create a Terraform CLI configuration file such as `dev.tfrc`:

```hcl
provider_installation {
  dev_overrides {
    "kumar1236/virtualbox" = "/absolute/path/to/terraform-provider-virtualbox/bin"
  }

  direct {}
}
```

Set `TF_CLI_CONFIG_FILE` to that file and run `terraform plan` or
`terraform apply`. The provider source in Terraform configurations must be
`kumar1236/virtualbox`.

## Documentation

Provider documentation is stored in `docs/` using the Terraform Registry
documentation layout.

## Releases

Push a semantic-version tag such as `v1.0.0`. The release workflow builds and
signs provider packages for the Terraform Registry. Repository secrets for the
GPG private key and passphrase must be configured before publishing a release.
