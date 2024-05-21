// ssh_config_test.go
package ssh

import (
	"go-inventory/internal/inventory"
	"testing"
)

func TestGenerate(t *testing.T) {
    hosts := []inventory.Host{
        {
            Name: "testHost",
            IP:   "192.168.1.1",
            SSH: inventory.SSH {
                User:    "testUser",
                KeyPath: "default",
            },
            Alias: "testAlias",
        },
    }

    expectedOutput := `
Host testHost
    HostName 192.168.1.1
    User testUser
Host testAlias
    HostName 192.168.1.1
    User testUser
`

    output, err := Generate(hosts, false)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    if output != expectedOutput {
        t.Fatalf("Expected:\n%s\nGot:\n%s", expectedOutput, output)
    }
}