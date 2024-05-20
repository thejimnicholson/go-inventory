// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	flag "github.com/spf13/pflag"

	"go-inventory/internal/ansible"
	"go-inventory/internal/inventory"
	"go-inventory/internal/ssh"
)

func printAsJSON(v interface{}) {
    value, err := json.MarshalIndent(v, "", "  ")
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    fmt.Println(string(value))
}

func main() {
    hostPtr := flag.String("host", "", "Specify the host name")
    listPtr := flag.Bool("list", false, "List all hosts")
    groupPtr := flag.String("group", "", "Specify the group name")
    typePtr := flag.String("type", "", "Specify the type")
    sshPtr := flag.Bool("ssh", false, "Print ssh config file")

    flag.Parse()

    hosts, err := inventory.LoadFromFile("./data/host_db.yaml")

    if err != nil {
        log.Fatalf("error: %v", err)
    }

    if *listPtr {
        inventory, err := ansible.Generate(hosts)
        if err != nil {
            log.Fatalf("error: %v", err)
        }
        fmt.Println(inventory)
    }

    if *sshPtr {
        config, err := ssh.Generate(hosts, false)
        if err != nil {
            log.Fatalf("error: %v", err)
        }
        fmt.Println(config)
    }

    if *hostPtr != "" {
        host := inventory.GetHostByName(*hostPtr)
        printAsJSON(host)
    }
    if *groupPtr != "" {
        hosts := inventory.GetHostsByGroup(*groupPtr)
        printAsJSON(hosts)
    }

    if *typePtr != "" {
        hosts := inventory.GetHostsByType(*typePtr)
        printAsJSON(hosts)
    }

}