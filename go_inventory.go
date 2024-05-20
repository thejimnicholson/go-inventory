package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Host struct {
	Name    string `yaml:"name"`
	Alias   string `yaml:"alias"`
	IP      string `yaml:"ip"`
	Ansible struct {
		Skip    bool   `yaml:"skip"`
		User    string `yaml:"user"`
		KeyPath string `yaml:"key_path"`
	} `yaml:"ansible"`
	Type   string   `yaml:"type"`
	Groups []string `yaml:"groups"`
	DNS    struct {
		Type   string `yaml:"type"`
		Domain string `yaml:"domain"`
	} `yaml:"dns"`
}

func main() {
	hostPtr := flag.String("host", "", "Specify the host name")
	listPtr := flag.Bool("list", false, "List all hosts")
	groupPtr := flag.String("group", "", "Specify the group name")
	typePtr := flag.String("type", "", "Specify the type")

	flag.Parse()

	data, err := os.ReadFile("host_db.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var hosts []Host

	err = yaml.Unmarshal([]byte(data), &hosts)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if *listPtr {
		fmt.Println(hosts)
		os.Exit(0)
	}

	if *hostPtr != "" {
		for _, host := range hosts {
			if host.Name == *hostPtr {
				fmt.Println(host)
				os.Exit(0)
			}
		}
		fmt.Println("Host not found")
		os.Exit(1)
	}

	if *groupPtr != "" {
		for _, host := range hosts {
			if strings.Contains(strings.Join(host.Groups, " "), *groupPtr) {
				fmt.Println(host)
			}
		}
	}

	if *typePtr != "" {
		for _, host := range hosts {
			if strings.Contains(host.Type, *typePtr) {
				fmt.Println(host)
			}
		}
	}
}
