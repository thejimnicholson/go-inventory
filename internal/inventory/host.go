// inventory/host.go
package inventory

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
    hosts []Host
    mu    sync.Mutex
)

type Host struct {
    Name    string `json:"name" yaml:"name"`
    Alias   string `json:"alias" yaml:"alias"`
    IP      string `json:"ip" yaml:"ip"`
    Ansible struct {
        Skip    bool   `json:"skip" yaml:"skip"`
        User    string `json:"user" yaml:"user"`
        KeyPath string `json:"key_path" yaml:"key_path"`
    } `json:"ansible" yaml:"ansible"`
    Type   string   `json:"type" yaml:"type"`
    Groups []string `json:"groups" yaml:"groups"`
    DNS    struct {
        Type   string `json:"type" yaml:"type"`
        Domain string `json:"domain" yaml:"domain"`
    } `json:"dns" yaml:"dns"`
}

func LoadFromFile(filename string) ([]Host, error) {
    mu.Lock()
    defer mu.Unlock()

    if hosts != nil {
        return hosts, nil
    }

    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("could not read file: %w", err)
    }

    err = yaml.Unmarshal(data, &hosts)
    if err != nil {
        return nil, fmt.Errorf("could not unmarshal data: %w", err)
    }

    return hosts, nil
}

func ListAllHosts() []Host {
    mu.Lock()
    defer mu.Unlock()

    if hosts == nil {
        return nil
    }

    // Return a copy of hosts to prevent modification of the original data
    return append([]Host(nil), hosts...)
}

func GetHostByName(name string) *Host {
    mu.Lock()
    defer mu.Unlock()
    for _, host := range hosts {
        if host.Name == name {
            return &host
        }
    }
    return nil
}

func GetHostsByGroup(groupName string) []Host {
    mu.Lock()
    defer mu.Unlock()
    var hostsByGroup []Host
    for _, host := range hosts {
        for _, group := range host.Groups {
            if group == groupName {
                hostsByGroup = append(hostsByGroup, host)
                break
            }
        }
    }
    return hostsByGroup
}

func GetHostsByType(hostType string) []Host {
    mu.Lock()
    defer mu.Unlock()
    var hostsByType []Host
    for _, host := range hosts {
        if host.Type == hostType {
            hostsByType = append(hostsByType, host)
        }
    }
    return hostsByType
}
