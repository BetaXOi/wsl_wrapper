#!/bin/bash

port=$1
user=$2
gw=$(ip route show default |awk '{print $3}')
echo ssh ningbo@$gw -p $port

#tmux new-session "ssh -o KexAlgorithms=+diffie-hellman-group14-sha1,diffie-hellman-group1-sha1 -o HostKeyAlgorithms=+ssh-rsa ningbo@$gw -p $port"
tmux new-session "/home/nb/bin/termtunnel ssh -o KexAlgorithms=+diffie-hellman-group14-sha1,diffie-hellman-group1-sha1 -o HostKeyAlgorithms=+ssh-rsa $user@$gw -p $port"


