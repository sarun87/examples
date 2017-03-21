# Files
## container_netns.c
Source code to demonstrate containers and network namespaces.
Compile code using gcc as follows

> gcc container_netns.c

Run the program without any command line arguments to create a container
with the same network namespace as the calling process.

> ./a.out

Run the program with "net" argument to create a container with a dedicated
network namespace for the container.

> ./a.out net

Output of the "ip a" command is recorded as below:

> /tmp/eg_parent_ipa.out -> Parent process's output of ip a command

> /tmp/eg_container_ipa.out -> Child process's (container) output of ip a command

Notice the difference in the o/p of "ip a" by the container in both cases

[Reference](https://www.redhat.com/archives/libvir-list/2008-January/msg00444.html)

## docker_point2point.sh
Script file to demonstrate and understand networking with containers.
The script creates two docker containers without networking followed by
adding the ends of the veth pair to each of the containers.
To run the script just execute the script as follows

> ./docker_point2point.sh

To explore the internals, use "ip netns", "ip netns exec" and "docker exec"
commands. You could even install tcpdump within the container & capture
packets.

## docker_point2point.go
Source code to demonstrate and understand networking with containers (in go).
First pass implementation of just directly executing the same commands that
the shell script executes.
Next steps would be to use netlink directly to perform linux networking operations.
Compile the code in go as follows

> go build docker_point2point.go

Run the program as follows

> ./docker_point2point
