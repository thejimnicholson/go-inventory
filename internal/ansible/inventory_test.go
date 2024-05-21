// ansible_test.go
package ansible

import (
	"go-inventory/internal/inventory"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	// Set up some hosts
	hosts := []inventory.Host{
		{
			Name: "testHost1",
			IP:   "192.168.1.1",
			SSH: inventory.SSH{
				User:    "testUser1",
				KeyPath: "default",
				Skip:    false,
			},
			Alias:  "testAlias1",
			Groups: []string{"testGroup1", "testGroup2"},
		},
		{
			Name: "testHost2",
			IP:   "192.168.1.2",
			SSH: inventory.SSH{
				User:    "testUser2",
				KeyPath: "default",
				Skip:    true,
			},
			Alias:  "testAlias2",
			Groups: []string{"testGroup1"},
		},
	}

	// Call Generate
	got, err := Generate(hosts)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check that the returned string matches the expected JSON
	expected := `{
  "testGroup1": {
    "hosts": [
      "testHost1"
    ]
  },
  "testGroup2": {
    "hosts": [
      "testHost1"
    ]
  }
}`
	if got != expected {
		t.Errorf("Expected JSON to be %s, got %s", expected, got)
	}
}

func TestHostData(t *testing.T) {
	// Set up a host
	host := `---
- name: testHost
  ip: 192.168.1.1
  ssh:
    user: testUser
    keyPath: default
  alias: testAlias
  groups: ["testGroup"]
  meta: {}
  dns: {}`

	// Write the host data to a YAML file
	file, err := os.CreateTemp("", "testHostData")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(host))
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	file.Close()

	// Load the hosts from the YAML file
	_, err = inventory.LoadFromFile(file.Name())
	if err != nil {
		t.Fatalf("Failed to load hosts from file: %v", err)
	}

	// Call HostData
	got, err := HostData("testHost")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check that the returned string matches the expected JSON
  expected := `{"alias":"testAlias","ansible_host":"192.168.1.1","ansible_ssh_private_key_file":"","ansible_ssh_user":"testUser","dns":{"type":"","domain":""},"groups":["testGroup"],"meta":{"type":"","host":"","id":""},"name":"testHost"}`
  if got != expected {
		t.Errorf("Expected JSON to be %s, got %s", expected, got)
	}
}
