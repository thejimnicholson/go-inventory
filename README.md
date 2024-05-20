# go-inventory

## Background

Managing my home network has become a challenge. I realized this when I tried to change the IP addresses for some of my servers.

I have a number of projects that need to have an accurate picture of my network. 

- ansible code that acts on specific servers within my inventory.
- an ansible playbook that updates route53 with my external and internal DNS records.
- an ansible playbook that generates nginx config files for my layer 4/layer 7 load balancer. 
- an ssh config file that gets checked into my dotfiles project and shared around to various desktops.
- a guest WiFi network that isolates all my IoT devices and home automation equipment. 

In addition to all this, I have plans for several other projects, including

- separating the physical network between my homelab gear and my family's work and personal devices.
- moving internal DNS from route53 to an internal hosted solution (probably coreDNS).
- implementing DHCP servers separate from my router, which has limitations (can't pre-allocate IPs.)
- possibly implementing vlans.
- moving my load balancer, which services my kubernetes clusters hosting apps that are available via the Internet, from a VM to a physical device.
- adding new servers to my Proxmox virtualization cluster.
- Implementing a storage network.

## Approach

All of these projects consume information about various hosts in my network. Typically, I have managed that information seperately, and had trouble keeping them in sync.

But what if I didn't have to do that? What if there was a single place where I could keep all the information about things inside my home network - VMs, hardware devices, etc.?

So, I looked at open-source IPAM tools, because we use InfoBlox at work, and I figured there must be something I can use that will do what I need. 

Boy, was I mistaken.

Open source IPAM tools *are* available, and I would imagine that they are as good as many commercial offerings. (Not necessarily a high bar.) But all of them that I looked into are a lot more interested in managing data center inventory than simply providing IP address management. And most of them have shitting implementations using quirky scripting languages that are fun to code in but a real pain in the butt to maintain. 

Plus, I figure, I didn't really need that. 

What I need is two things:

- a common source of truth for information about hosts within my home network. 
- a way to use that truth to generate (dynamically or on-demand) configuration files for the tools I use. 

The answer to the first need seems to be: I need a YAML file. That turns out to be fairly easy, because some of the tools (those that rely on ansible) already are using YAML files as their inventory source.

The answer to the second is this project. 

The idea is to build a tool that can read the YAML file and output various different formats, which then can be fed into my tools. This tool can be called from automation that runs whenever the YAML is updated, and the resulting files would then be distributed to the places where they are needed.

## Specifics

This program reads a YAML file (`data/host_db.yaml`) and emits output files for various tools:

- It can produce an ssh config file (`.ssh/config`) that configures ssh connections to all the hosts that are reachable via ssh. 
- It can produce an [ansible](https://www.ansible.com) inventory that can be used to drive playbooks that run against servers within the network. 
- It will update AWS route53 with DNS records as appropriate.
- It will produce a zone file for coreDNS, which will handle internal DNS.

