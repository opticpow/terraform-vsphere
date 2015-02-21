package main

import (
    "testing"
    "fmt"
    "github.com/hashicorp/terraform/helper/resource"
    "github.com/hashicorp/terraform/terraform"
)

// Requirements:
// base VM name: basic
// folder: <root>
// VM tools installed
func TestAccVirtualMachine_basic(t *testing.T) {
    resource.Test(t, resource.TestCase{
        PreCheck:  func() { testAccPreCheck(t) },
        Providers: testAccProviders,
        Steps: []resource.TestStep{
            resource.TestStep{
                Config: testAccVirtualMachine_basic,
                Check: resource.ComposeTestCheckFunc(
                testAccCheckVirtualMachineState("name", "basic-1"),
                ),
            },
        },
        CheckDestroy: testAccCheckVirtualMachineDestroy,
    })
}

func testAccCheckVirtualMachineState(key, value string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        rs, ok := s.RootModule().Resources["vsphere_virtual_machine.app"]
        if !ok {
            return fmt.Errorf("Not found: %s", "vsphere_virtual_machine.app")
        }

        if rs.Primary.ID == "" {
            return fmt.Errorf("No ID is set")
        }

        p := rs.Primary
        if p.Attributes[key] != value {
            return fmt.Errorf("%s != %s (actual: %s)", key, value, p.Attributes[key])
        }
        if p.Attributes["ip_address"] == "" {
            return fmt.Errorf("IP address is not set")
        }

        return nil
    }
}

func testAccCheckVirtualMachineDestroy(s *terraform.State) error {
    return fmt.Errorf("####### %s", s)
}

const testAccVirtualMachine_basic = `
resource "vsphere_virtual_machine" "app" {
    name =  "basic-1"
    image = "basic"
}`
