#!/bin/bash
function check_version {
	grep "VERSION" ./main.go | grep  \"$1\"
	if [[ $? -eq 1 ]]
	then
		echo "VERSION hasn't been updated"
		exit 1
	fi
}

function build_binary {
    CURRENT_BIN=$(basename "$PWD")
    OUTPUT="../build/akamai-${CURRENT_BIN}-$1"
    mkdir -p ../build

    GOOS=darwin GOARCH=amd64 go build -o ${OUTPUT}-macamd64 .
    shasum -a 256 ${OUTPUT}-macamd64 | awk '{print $1}' > ${OUTPUT}-macamd64.sig
    GOOS=linux GOARCH=amd64 go build -o ${OUTPUT}-linuxamd64 .
    shasum -a 256 ${OUTPUT}-linuxamd64 | awk '{print $1}' > ${OUTPUT}-linuxamd64.sig
    GOOS=linux GOARCH=386 go build -o ${OUTPUT}-linux386 .
    shasum -a 256 ${OUTPUT}-linux386 | awk '{print $1}' > ${OUTPUT}-linux386.sig
    GOOS=windows GOARCH=386 go build -o ${OUTPUT}-windows386.exe .
    shasum -a 256 ${OUTPUT}-windows386.exe | awk '{print $1}' > ${OUTPUT}-windows386.exe.sig
    GOOS=windows GOARCH=amd64 go build -o ${OUTPUT}-windowsamd64.exe .
    shasum -a 256 ${OUTPUT}-windowsamd64.exe | awk '{print $1}' > ${OUTPUT}-windowsamd64.exe.sig
}

if [[ -z "$1" ]]
then
	echo "Version not supplied."
	echo "Usage: build.sh <version>"
	exit 1
fi

CURRENT_BIN=$(basename "$PWD")
if [[ $CURRENT_BIN == "cli-api-gateway" ]]
then
    for CURRENT_BIN in api-gateway api-keys api-security
    do
        echo "Building ${CURRENT_BIN}"
        cd $CURRENT_BIN
        check_version $1
        build_binary $1
        cd ..
    done
    echo "Done."
else
    echo "Building ${CURRENT_BIN}"
    check_version $1
    build_binary $1
    echo "Done."
fi
