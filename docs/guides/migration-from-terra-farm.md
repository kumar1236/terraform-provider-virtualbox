---
page_title: "Migrating from terra-farm/virtualbox"
subcategory: "Guides"
description: |-
  Guide for migrating from the terra-farm/virtualbox provider to kumar1236/virtualbox.
---

# Migrating from terra-farm/virtualbox

This guide helps you migrate from the discontinued `terra-farm/virtualbox` provider to `kumar1236/virtualbox`.

## Step 1: Update Provider Source

Change your `required_providers` block:

```hcl
# Before (terra-farm)
terraform {
  required_providers {
    virtualbox = {
      source  = "terra-farm/virtualbox"
      version = "~> 0.2"
    }
  }
}

# After (kumar1236)
terraform {
  required_providers {
    virtualbox = {
      source  = "kumar1236/virtualbox"
      version = "~> 1.0"
    }
  }
}
```

## Step 2: Migrate State

Replace the provider in your Terraform state:

```bash
terraform state replace-provider terra-farm/virtualbox kumar1236/virtualbox
```

## Step 3: Review Configuration

All resource names remain the same (`virtualbox_vm`, etc.). New default values match the old hardcoded values, so existing configurations should work without changes:

| Attribute | terra-farm (hardcoded) | kumar1236/virtualbox (default) |
|-----------|----------------------|----------------------|
| `os_type` | `Linux_64` | `Linux_64` |
| `vram` | `20` | `20` |
| `firmware` | `bios` | `bios` |
| `graphics_controller` | N/A | `vmsvga` |
| `gui` | N/A (headless) | `false` (headless) |

## Step 4: Run Plan

```bash
terraform init -upgrade
terraform plan
```

The plan should show no changes for existing resources. If you see a diff on `memory` format (e.g., `"512 mib"` vs `"512mib"`), this is a normalization fix — apply it once and it will be stable.

## Step 5: Take Advantage of New Features

You now have access to 30+ new attributes on `virtualbox_vm`, plus 4 new resources and 3 data sources. See the [provider documentation](../index.md) for details.

## Breaking Changes

- **Memory format**: The provider normalizes memory values without spaces (`"512mib"` not `"512 mib"`). This may cause a one-time diff.
- **`user_data`**: Was deprecated and non-functional in terra-farm. Now works via cloud-init guest properties.
- **`url` attribute**: Still supported but deprecated. Use `image` instead.

## Getting Help

- [GitHub Issues](https://github.com/kumar1236/terraform-provider-virtualbox/issues)
- [Provider Documentation](https://registry.terraform.io/providers/kumar1236/virtualbox/latest/docs)
