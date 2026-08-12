# Changelog

All notable changes to this provider will be documented in this file.

## Unreleased

### Added

- Initial release under the `kumar1236/virtualbox` Terraform Registry address.
- OVA and OVF import support.
- Cloud-init `user_data` support through VirtualBox guest properties.

### Fixed

- Respect `status = "poweroff"` during VM creation, OVA import, linked-clone creation, and updates, allowing provisioning media to be attached before first boot.
- Apply `user_data` to linked clones and assign the Terraform resource ID before readiness checks.
