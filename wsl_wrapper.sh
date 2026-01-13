#!/bin/bash

url=$1
url=${url##ssh://}
auth=$(echo $url | awk -F@ '{print $1}')
site=$(echo $url | awk -F@ '{print $2}')
user=${auth%%:*}
pass=${auth##*:}
host=${site%%:*}
port=${site##*:}


export PATH=$HOME/bin:$PATH

#tmux new-session "/home/nb/bin/termtunnel ssh -o KexAlgorithms=+diffie-hellman-group14-sha1,diffie-hellman-group1-sha1 -o HostKeyAlgorithms=+ssh-rsa $user@localhost -p $port"
#echo /home/nb/bin/termtunnel ssh -o KexAlgorithms=+diffie-hellman-group14-sha1,diffie-hellman-group1-sha1 -o HostKeyAlgorithms=+ssh-rsa $user@localhost -p $port
#/home/nb/bin/termtunnel ssh -o KexAlgorithms=+diffie-hellman-group14-sha1,diffie-hellman-group1-sha1 -o HostKeyAlgorithms=+ssh-rsa $user@localhost -p $port

tmux has-session -t 4a
if [ $? -eq 0  ]; then
	echo tmux new-window -t 4a -n $site sshpass -p $pass ssh -oStrictHostKeyChecking=no -oKexAlgorithms=+diffie-hellman-group14-sha1,diffie-hellman-group1-sha1 -oHostKeyAlgorithms=+ssh-rsa -p$port $user@$host
	tmux new-window -t 4a -n $site sshpass -p $pass ssh -oStrictHostKeyChecking=no -oKexAlgorithms=+diffie-hellman-group14-sha1,diffie-hellman-group1-sha1 -oHostKeyAlgorithms=+ssh-rsa -p$port $user@$host
else
	echo tmux new-session -s 4a -n $site sshpass -p $pass ssh -oStrictHostKeyChecking=no -oKexAlgorithms=+diffie-hellman-group14-sha1,diffie-hellman-group1-sha1 -oHostKeyAlgorithms=+ssh-rsa -p$port $user@$host
	tmux new-session -s 4a -n $site sshpass -p $pass ssh -oStrictHostKeyChecking=no -oKexAlgorithms=+diffie-hellman-group14-sha1,diffie-hellman-group1-sha1 -oHostKeyAlgorithms=+ssh-rsa -p$port $user@$host
fi

sleep 10
