#!/bin/bash

systemArchitecture=$(uname -m)
case $systemArchitecture in
    x86_64)
        architecture="amd64"
        ;;
    arm*)
        if [ "$systemArchitecture" == "armv7l" ]; then
            architecture="arm"
            goarm="7"
        else # simplification, I only use arm64 and armv7l
            architecture="arm64"
        fi
        ;;
    *)
        echo "$systemArchitecture is not supported by this script"
        exit 1
        ;;
esac


if [ "$architecture" == "arm" ]; then
    buildEnv="GOARCH=$architecture GOARM=$goarm"
else
    buildEnv="GOARCH=$architecture go build"
fi

buildCommand="$buildEnv go build -o proxy"

eval $buildCommand