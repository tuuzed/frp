#!/bin/sh
export CGO_ENABLED=1

##
# Android
##
export GOOS="android"
export TOOL_CHAINS_HOME="/opt/Android/sdk/ndk-bundle/toolchains/llvm/prebuilt/linux-x86_64/bin"

export CC=$TOOL_CHAINS_HOME/armv7a-linux-androideabi21-clang
export CXX=$TOOL_CHAINS_HOME/armv7a-linux-androideabi21-clang++
export GOARCH="arm"
go build -ldflags="-s -w" -buildmode=c-shared -o dists/android/armeabi-v7a/libsharedfrp.so client.go server.go common.go

export CC=$TOOL_CHAINS_HOME/aarch64-linux-android21-clang
export CXX=$TOOL_CHAINS_HOME/aarch64-linux-android21-clang++
export GOARCH="arm64"
go build -ldflags="-s -w" -buildmode=c-shared -o dists/android/arm64-v8a/libsharedfrp.so client.go server.go common.go

export CC=$TOOL_CHAINS_HOME/i686-linux-android21-clang
export CXX=$TOOL_CHAINS_HOME/i686-linux-android21-clang++
export GOARCH="386"
go build -ldflags="-s -w" -buildmode=c-shared -o dists/android/x86/libsharedfrp.so client.go server.go common.go

export CC=$TOOL_CHAINS_HOME/x86_64-linux-android21-clang
export CXX=$TOOL_CHAINS_HOME/x86_64-linux-android21-clang++
export GOARCH="amd64"
go build -ldflags="-s -w" -buildmode=c-shared -o dists/android/x86_64/libsharedfrp.so client.go server.go common.go

##
# Linux
##
export GOOS=linux
export CC=/usr/bin/gcc
export CXX=/usr/bin/g++

export GOARCH=amd64
go build -ldflags="-s -w" -buildmode=c-shared -o dists/linux/x86_64/libsharedfrp.so client.go server.go common.go

export GOARCH=386
go build -ldflags="-s -w" -buildmode=c-shared -o dists/linux/x86/libsharedfrp.so client.go server.go common.go

export GOARCH=arm
go build -ldflags="-s -w" -buildmode=c-shared -o dists/linux/arm/libsharedfrp.so client.go server.go common.go

export GOARCH=arm64
go build -ldflags="-s -w" -buildmode=c-shared -o dists/linux/arm64/libsharedfrp.so client.go server.go common.go
