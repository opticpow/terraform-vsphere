package main

import (
    "testing"
    "fmt"
    "github.com/hashicorp/terraform/helper/resource"
    "github.com/hashicorp/terraform/terraform"
    "github.com/vmware/govmomi"
    "github.com/vmware/govmomi/vim25/types"
)

// Requirements:
// base VM name: basic
// folder: <root>
// VM tools installed
func TestAccVirtualMachine_basic(t *testing.T) {
    var vm_id string
    resource.Test(t, resource.TestCase{
        PreCheck:  func() { testAccPreCheck(t) },
        Providers: testAccProviders,
        Steps: []resource.TestStep{
            resource.TestStep{
                Config: testAccVirtualMachine_basic,
                Check: resource.ComposeTestCheckFunc(
                testAccCheckVirtualMachineState(&vm_id),
                ),
            },
        },
        CheckDestroy: resource.ComposeTestCheckFunc(
            testAccCheckVirtualMachineDestroy(&vm_id),
        ),
    },
    )
}

func testAccCheckVirtualMachineState(vm_id *string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources["vsphere_virtual_machine.app"]
        if !ok {
            return fmt.Errorf("Not found: %s", "vsphere_virtual_machine.app")
        }

        p := rs.Primary
        if p.ID == "" {
            return fmt.Errorf("No ID is set")
        }
        if p.Attributes["ip_address"] == "" {
            return fmt.Errorf("IP address is not set")
        }
        *vm_id = p.ID

        return nil
    }
}

func testAccCheckVirtualMachineDestroy(vm_id *string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        client := testAccProvider.Meta().(*govmomi.Client)

        vm_mor := types.ManagedObjectReference{Type: "VirtualMachine", Value: *vm_id }
        err := client.Properties(vm_mor, []string{"summary"}, &vm_mor)
        if err == nil {
            return fmt.Errorf("Record still exists")
        }

        return nil
    }
}

const testAccVirtualMachine_basic = `
resource "vsphere_virtual_machine" "app" {
    name =  "basic-1"
    image = "basic"
}`
