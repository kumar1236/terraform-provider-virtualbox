# Changelog

All notable changes to this provider will be documented in this file.

## Unreleased

## v2.1.2 - 2026-08-13

### Added

- Initial release under the `kumar1236/virtualbox` Terraform Registry address.
- OVA and OVF import support.
- Cloud-init `user_data` support through VirtualBox guest properties.

### Fixed

- Respect `status = "poweroff"` during VM creation, OVA import, linked-clone creation, and updates, allowing provisioning media to be attached before first boot.
- Apply `user_data` to linked clones and assign the Terraform resource ID before readiness checks.
- Keep NAT port-forwarding updates idempotent by replacing rules by name.
- Preserve configured NAT forwarding and NIC options during refresh.
- Read NIC hardware types and memory values back into state without causing perpetual diffs.
- Recognize VirtualBox machine-not-found errors reported through stderr.
- Reduce unnecessary guest-readiness waits when a VM has no immediately available address.
