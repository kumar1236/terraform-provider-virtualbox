package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestVMStatusValidation(t *testing.T) {
	validate := resourceVM().Schema["status"].ValidateFunc

	for _, status := range []string{"running", "poweroff"} {
		_, errs := validate(status, "status")
		if len(errs) != 0 {
			t.Errorf("expected status %q to be valid, got %v", status, errs)
		}
	}

	for _, status := range []string{"poweredoff", "stopped", "paused", ""} {
		_, errs := validate(status, "status")
		if len(errs) == 0 {
			t.Errorf("expected status %q to be rejected", status)
		}
	}
}

func TestApplyDesiredPowerStatePoweroffDoesNotStartVM(t *testing.T) {
	d := schema.TestResourceDataRaw(t, resourceVM().Schema, map[string]any{
		"name":   "test-vm",
		"status": "poweroff",
	})
	vm := &Machine{UUID: "test-vm-uuid"}

	// No VirtualBox executable or mock is needed: the poweroff path must not
	// execute VBoxManage at all.
	if err := applyDesiredPowerState(context.Background(), d, vm, nil); err != nil {
		t.Fatalf("poweroff should leave the VM stopped: %v", err)
	}
}
