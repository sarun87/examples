#!/bin/bash

. util.sh

DEMO_RUN_FAST=1

clear

# set kubeconfig
run "export KUBECONFIG=/Users/arunsriraman/kubeconfigs/arun-calico-01.yml"

# Deploy NGINX Server
# kubectl run nginx --replicas=2 --image=nginx
# kubectl expose deployment nginx --port=80 --type=LoadBalancer --name=nginx-service 

# Deploy wordpress
# helm init
# helm install --name=wordpress stable/wordpress

# Get LB IPs / DNS Names
NGINX_LB=$(kubectl get svc nginx-service -o custom-columns=:status.loadBalancer.ingress[0].hostname --no-headers)
WP_LB=$(kubectl get svc wordpress-server-wordpress -o custom-columns=:status.loadBalancer.ingress[0].hostname --no-headers)

desc "Pods & services in default ns"
run "kubectl get pods -n default"
run "kubectl get svc -n default -o wide"

clear

desc "Trying to curl  wordpress server"
run "curl -m 8 $WP_LB"

desc "Try curling nginx server"
run "curl -m 8 $NGINX_LB"

clear

desc "NetworkPolicy for denyall in default ns"
run "cat deny_all_defaultns.yml"

desc "Apply denyall network policy"
run "kubectl apply -f deny_all_defaultns.yml"

clear

desc "Trying to curl  wordpress server"
run "curl -m 5 $WP_LB"

desc "Trying to curl nginx server"
run "curl -m 5 $NGINX_LB"

clear

desc "Lets allow access to nginx alone"
run "cat allow_nginx.yml"

run "kubectl apply -f allow_nginx.yml"

clear

desc "Curl wordpress server (we shouldn't be able to)"
run "curl -m 5 $WP_LB"

desc "Curl nginx"
run "curl -m 8 $NGINX_LB"

desc "Remove all rules"
run "kubectl delete -f ."


