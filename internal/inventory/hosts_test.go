// hosts_test.go
package inventory

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestLoadFromFile(t *testing.T) {
    // Create a temporary file
    tmpfile, err := ioutil.TempFile("", "example")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpfile.Name()) // clean up

    // Write a test YAML to the file
    text := []byte(`
- name: testHost
  alias: testAlias
  ip: 192.168.1.1
  ssh:
    user: testUser
    key_path: default
  meta:
    type: testType
    host: testHost
    id: testID
  groups:
    - testGroup1
    - testGroup2
  dns:
    type: testType
    domain: testDomain
`)
    if _, err := tmpfile.Write(text); err != nil {
        t.Fatal(err)
    }
    if err := tmpfile.Close(); err != nil {
        t.Fatal(err)
    }

    // Call LoadFromFile
    hosts, err := LoadFromFile(tmpfile.Name())
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    // Check the returned hosts
    if len(hosts) != 1 {
        t.Fatalf("Expected 1 host, got %d", len(hosts))
    }

    host := hosts[0]
    if host.Name != "testHost" {
        t.Errorf("Expected name to be 'testHost', got '%s'", host.Name)
    }
    // ... repeat for other fields
}

func TestListAllHosts(t *testing.T) {
    // Set up some hosts
    hosts = []Host{
        {
            Name: "testHost1",
            IP:   "192.168.1.1",
            SSH: SSH{
                User:    "testUser1",
                KeyPath: "default",
            },
            Alias: "testAlias1",
        },
        {
            Name: "testHost2",
            IP:   "192.168.1.2",
            SSH: SSH{
                User:    "testUser2",
                KeyPath: "default",
            },
            Alias: "testAlias2",
        },
    }

    // Call ListAllHosts
    listedHosts := ListAllHosts()

    // Check that the returned slice matches the hosts
    if !reflect.DeepEqual(listedHosts, hosts) {
        t.Errorf("Expected hosts to be %v, got %v", hosts, listedHosts)
    }
}

func TestGetHostByName(t *testing.T) {
    // Set up a host
    host := Host{
        Name: "testHost",
        IP:   "192.168.1.1",
        SSH: SSH{
            User:    "testUser",
            KeyPath: "default",
        },
        Alias: "testAlias",
    }
    hosts = []Host{host}

    // Call GetHostByName
    gotHost := GetHostByName("testHost")

    // Check that the returned host matches the expected host
    if !reflect.DeepEqual(gotHost, &host) {
        t.Errorf("Expected host to be %v, got %v", &host, gotHost)
    }
}

func TestGetHostsByGroup(t *testing.T) {
    // Set up some hosts
    hosts = []Host{
        {
            Name:   "testHost1",
            IP:     "192.168.1.1",
            SSH: SSH{
                User:    "testUser1",
                KeyPath: "default",
            },
            Alias:  "testAlias1",
            Groups: []string{"testGroup"},
        },
        {
            Name:   "testHost2",
            IP:     "192.168.1.2",
            SSH: SSH{
                User:    "testUser2",
                KeyPath: "default",
            },
            Alias:  "testAlias2",
            Groups: []string{"otherGroup"},
        },
    }

    // Call GetHostsByGroup
    gotHosts := GetHostsByGroup("testGroup")

    // Check that the returned hosts match the expected hosts
    expectedHosts := []Host{hosts[0]}
    if !reflect.DeepEqual(gotHosts, expectedHosts) {
        t.Errorf("Expected hosts to be %v, got %v", expectedHosts, gotHosts)
    }
}

func TestGetHostsByType(t *testing.T) {
    // Set up some hosts
    hosts = []Host{
        {
            Name:   "testHost1",
            IP:     "192.168.1.1",
            SSH: SSH{
                User:    "testUser1",
                KeyPath: "default",
            },
            Alias:  "testAlias1",
            Meta: Meta{
                Type: "testType",
            },
        },
        {
            Name:   "testHost2",
            IP:     "192.168.1.2",
            SSH: SSH{
                User:    "testUser2",
                KeyPath: "default",
            },
            Alias:  "testAlias2",
            Meta: Meta{
                Type: "otherType",
            },
        },
    }

    // Call GetHostsByType
    gotHosts := GetHostsByType("testType")

    // Check that the returned hosts match the expected hosts
    expectedHosts := []Host{hosts[0]}
    if !reflect.DeepEqual(gotHosts, expectedHosts) {
        t.Errorf("Expected hosts to be %v, got %v", expectedHosts, gotHosts)
    }
}