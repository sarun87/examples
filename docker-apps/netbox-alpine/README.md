# Alpine based docker container with networking tools pre-installed

Image: sarun87/netbox-alpine:v1

### Tools available as part of image

bash, curl, ssh, vim, sudo

### Additional networking specific tools/utilities
nc, tcpdump, iptables, iputils

### Alpine image comes with
ip, route, ifconfig, netstat, brctl installed

-----

## Running netbox

Run  docker container as privileged mounting root file system as below for added debugging capability

### Run a specific command and exit the container

    docker run -v /:/rootfs --privileged=true --net=host --name netbox --rm sarun87/netbox-alpine:v1 <command>

### Run the container in the background and exec into the container using bash to execute commands

    docker run -d -v /:/rootfs --privileged=true --net=host --name netbox sarun87/netbox-alpine:v1 sleep 20000
    docker exec -it netbox bash


