// ansible_inventory.go
package ansible

import (
	"encoding/json"
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