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

The public Terraform Registry requires this repository to remain public and to
use the `terraform-provider-virtualbox` naming convention. Before the first
release, create a dedicated RSA or DSA GPG signing key (the Registry does not
accept ECC signing keys), add its public key to the Terraform Registry, and
configure these GitHub Actions repository secrets:

- `GPG_PRIVATE_KEY`: ASCII-armored private signing key
- `PASSPHRASE`: passphrase for that signing key

To publish a release:

1. Run `go test ./...` and `go build ./...`.
2. Update `CHANGELOG.md`, commit the release changes, and push `main`.
3. Create and push a semantic-version tag, for example `v2.1.2`.
4. Confirm the `release` GitHub Actions workflow succeeds and publishes the
   checksums, detached signature, manifest, and platform archives.
5. For the first release under this namespace, sign in to the Terraform
   Registry with GitHub and register `kumar1236/terraform-provider-virtualbox`.

Do not create the GitHub release manually. Pushing the tag triggers the release
workflow, which signs the checksums and publishes the completed draft.
