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

type SSH struct {
    Skip    bool   `json:"skip" yaml:"skip"`
    User    string `json:"user" yaml:"user"`
    KeyPath string `json:"key_path" yaml:"key_path"`
}

type Meta struct {
    Type   string   `json:"type" yaml:"type"`
    Host   string   `json:"host" yaml:"host"`
    ID     string   `json:"id" yaml:"id"`
}

type DNS struct {
    Type   string `json:"type" yaml:"type"`
    Domain string `json:"domain" yaml:"domain"`
}

type Host struct {
    Name   string   `json:"name" yaml:"name"`
    Alias  string   `json:"alias" yaml:"alias"`
    IP     string   `json:"ip" yaml:"ip"`
    SSH    SSH      `json:"ssh" yaml:"ssh"`
    Meta   Meta     `json:"meta" yaml:"meta"`
    Groups []string `json:"groups" yaml:"groups"`
    DNS    DNS      `json:"dns" yaml:"dns"`
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

func GetAllGroups() []string {
    mu.Lock()
    defer mu.Unlock()

    if hosts == nil {
        return nil
    }

    groupSet := make(map[string]struct{})
    for _, host := range hosts {
        for _, group := range host.Groups {
            groupSet[group] = struct{}{}
        }
    }

    groups := make([]string, 0, len(groupSet))
    for group := range groupSet {
        groups = append(groups, group)
    }

    return groups
}

func GetAllHostNames() *[]string {
    defer mu.Unlock()

    if hosts == nil {
        return nil
    }

    var hostnames []string

    for _, host := range hosts {
        hostnames = append(hostnames, host.Name)
    }
    return &hostnames
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
        if host.Meta.Type == hostType {
            hostsByType = append(hostsByType, host)
        }
    }
    return hostsByType
}
