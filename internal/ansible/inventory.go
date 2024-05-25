// ansible_inventory.go
package ansible

import (
	"encoding/json"
	"fmt"
	"go-inventory/internal/inventory"
)

func Generate(hosts []inventory.Host) (string, error) {
    inventory := make(map[string]map[string][]string)

    for _, host := range hosts {
        if !host.SSH.Skip {
            for _, group := range host.Groups {
                if _, ok := inventory[group]; !ok {
                    inventory[group] = make(map[string][]string)
                }
                inventory[group]["hosts"] = append(inventory[group]["hosts"], host.Name)
            }
        }
    }

    inventoryJSON, err := json.MarshalIndent(inventory, "", "  ")
    if err != nil {
        return "", err
    }

    return string(inventoryJSON), nil
}

func HostData(name string) (string, error) {
    // Find the host with the given name
    host := inventory.GetHostByName(name)
    if host == nil {
        return "", fmt.Errorf("no host with name %s", name)
    }

    // Create a map with the host data formatted for Ansible
    data := make(map[string]interface{})
    data["ansible_host"] = host.IP
    data["ansible_ssh_private_key_file"] = host.SSH.KeyPath
    data["ansible_ssh_user"] = host.SSH.User
    data["name"] = host.Name
    data["alias"] = host.Alias
    data["groups"] = host.Groups
    data["meta"] = host.Meta
    data["dns"] = host.DNS

    // Convert the map to JSON
    jsonData, err := json.MarshalIndent(data, "", "  ") 
    if err != nil {
        return "", fmt.Errorf("failed to marshal host data to JSON: %w", err)
    }

    return string(jsonData), nil
}