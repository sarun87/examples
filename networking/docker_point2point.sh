#!/bin/bash

# Program to set up two busy box containers and connect them
# with a veth pair. basic point2point connectivity test
# Author: Arun Sriraman <sarun87[at]gmail[dot]com>

# Copyright (c) 2017 Arun Sriraman.
# 
# This program is free software: you can redistribute it and/or modify  
# it under the terms of the GNU General Public License as published by  
# the Free Software Foundation, version 3.
#
# This program is distributed in the hope that it will be useful, but 
# WITHOUT ANY WARRANTY; without even the implied warranty of 
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU 
# General Public License for more details.
#
# You should have received a copy of the GNU General Public License 
# along with this program. If not, see <http://www.gnu.org/licenses/>.
#

echo "Delete containers if they are already present"
docker kill net_demo
docker kill net_demo_2
docker rm net_demo
docker rm net_demo_2
unlink /var/run/netns/net_demo
unlink /var/run/netns/net_demo_2

echo "Creating container: net_demo"
docker run --name net_demo --network none -td busybox

echo "Creating container: net_demo_2"
docker run --name net_demo_2 --network none -td busybox

echo "Setting up namespace runtime dir"
mkdir -p /var/run/netns

c1_pid="$(docker inspect --format '{{.State.Pid}}' net_demo)"
c2_pid="$(docker inspect --format '{{.State.Pid}}' net_demo_2)"

echo "Setting up container namespaces"
ln -s /proc/$c1_pid/ns/net /var/run/netns/net_demo
ln -s /proc/$c2_pid/ns/net /var/run/netns/net_demo_2

echo "Create and add veths to the containers"
ip link add veth0 type veth peer name veth1
ip link set veth0 netns net_demo
ip link set veth1 netns net_demo_2

echo "Assign IPs to veths"
ip netns exec net_demo ifconfig veth0 192.168.122.1/24
ip netns exec net_demo_2 ifconfig veth1 192.168.122.2/24

echo "Testing setup"
docker exec net_demo ping -c4 192.168.122.2

