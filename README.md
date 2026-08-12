[![CI](https://github.com/kumar1236/terraform-provider-virtualbox/actions/workflows/ci.yaml/badge.svg)](https://github.com/kumar1236/terraform-provider-virtualbox/actions/workflows/ci.yaml)

# VirtualBox provider for Terraform

Source and issue tracking are available in the [GitHub repository](https://github.com/kumar1236/terraform-provider-virtualbox).

## Usage

Until the provider is published in the Terraform Registry, build it locally
and use a Terraform CLI `dev_overrides` entry as described in
[CONTRIBUTING.md](CONTRIBUTING.md).

```tf
terraform {
  required_providers {
    virtualbox = {
      source = "kumar1236/virtualbox"
    }
  }
}

resource "virtualbox_vm" "basic" {
  name   = "basic-vm"
  image  = "https://app.vagrantup.com/ubuntu/boxes/bionic64/versions/20180903.0.0/providers/virtualbox.box"
  cpus   = 2
  memory = "1024mib"

  network_adapter {
    type = "nat"
  }
}
```

## Examples

The [`/examples`](/examples) directory contains ready-to-use configurations:

| Example | Description |
|---------|-------------|
| [basic](examples/basic/) | Simple VM with NAT networking |
| [port-forwarding](examples/port-forwarding/) | NAT with SSH and HTTP port forwarding |
| [multi-disk](examples/multi-disk/) | VM with additional NVMe data disk |
| [windows-vm](examples/windows-vm/) | Windows 11 VM with EFI and GUI |
| [complete](examples/complete/) | All features: networks, disks, snapshots, shared folders |

If you want to contribute documentation changes, see the [Contribution guide](CONTRIBUTING.md).

## Limitations

- __Experimental provider!__
- We only officially support the latest version of Go, Virtualbox and Terraform. The provider might be compatible and work with other versions
  but we do not provide any level of support for this due to lack of time.
- The defaults here are only tested with the [vagrant insecure (packer) keys](https://github.com/hashicorp/vagrant/tree/master/keys) as the login.

## Contributors

Special thanks to all contributors, and [@ccll](https://github.com/ccll) for donating the original project to the terra-farm group!

Inspired by [terraform-provider-vix](https://github.com/hooklift/terraform-provider-vix)

## License

MIT. See [LICENSE](LICENSE).
