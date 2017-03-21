package main

import (
	"fmt"
	"os/exec"
)

func exec_cmd(cmd string) []byte {
	fmt.Println("Executing:", cmd)
	out, _ := exec.Command("/bin/sh", "-c", cmd).Output()
	return out
}

func main() {
	// Storage for container names
	docker_containers := [2]string{"net_demo", "net_demo_2"}

	// Run commands to create 2 docker containers with names net_demo
	// and net_demo_2 using a busy box image
	for i := 0; i < 2; i++ {
		fmt.Println("Creating container: ", docker_containers[i])
		cmd := fmt.Sprintf("docker run --name %s --network none -td busybox", docker_containers[i])
		out := exec_cmd(cmd)
		fmt.Printf("%s", out)
	}

	// Create the netns runtime directory if not already present
	exec_cmd("mkdir -p /var/run/netns")

	// Get the process id of each of the containers & mount their
	// namespaces
	// The right way of doing this is through netlink - https://github.com/milosgajdos83/tenus
	c1_pid := exec_cmd(fmt.Sprintf("docker inspect --format '{{.State.Pid}}' %s", docker_containers[0]))
	c2_pid := exec_cmd(fmt.Sprintf("docker inspect --format '{{.State.Pid}}' %s", docker_containers[1]))

	exec_cmd(fmt.Sprintf("ln -s /proc/%s/ns/net /var/run/netns/%s", c1_pid, docker_containers[0]))
	exec_cmd(fmt.Sprintf("ln -s /proc/%s/ns/net /var/run/netns/%s", c2_pid, docker_containers[1]))

	// Create veth pair
	exec_cmd("ip link add veth0 type veth peer name veth1")

	// Set ends of pair to each container
	exec_cmd(fmt.Sprintf("ip link set veth0 netns %s", docker_containers[0]))
	exec_cmd(fmt.Sprintf("ip link set veth1 netns %s", docker_containers[1]))

	// Set Ip addresses for the links
	exec_cmd(fmt.Sprintf("ip netns exec %s ifconfig veth0 192.168.122.1/24", docker_containers[0]))
	exec_cmd(fmt.Sprintf("ip netns exec %s ifconfig veth1 192.168.122.2/24", docker_containers[1]))

	// Run test from the container.
	exec_cmd(fmt.Sprintf("docker exec %s ping -c4 192.168.122.2", docker_containers[0]))
}
