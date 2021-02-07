$env:CGO_ENABLED="1"

##
# Android
##
$env:GOOS="android"
$TOOL_CHAINS_HOME="C:\opt\Android\sdk\ndk-bundle\toolchains\llvm\prebuilt\windows-x86_64\bin"

$env:CC="$TOOL_CHAINS_HOME\armv7a-linux-androideabi21-clang.cmd"
$env:CXX="$TOOL_CHAINS_HOME\armv7a-linux-androideabi21-clang++.cmd"
$env:GOARCH="arm"
go build -ldflags="-s -w" -buildmode=c-shared -o dists/android/armeabi-v7a/libsharedfrp.so client.go server.go common.go

$env:CC="$TOOL_CHAINS_HOME\aarch64-linux-android21-clang.cmd"
$env:CXX="$TOOL_CHAINS_HOME\aarch64-linux-android21-clang++.cmd"
$env:GOARCH="arm64"
go build -ldflags="-s -w" -buildmode=c-shared -o dists/android/arm64-v8a/libsharedfrp.so client.go server.go common.go

$env:CC="$TOOL_CHAINS_HOME\i686-linux-android21-clang.cmd"
$env:CXX="$TOOL_CHAINS_HOME\i686-linux-android21-clang++.cmd"
$env:GOARCH="386"
go build -ldflags="-s -w" -buildmode=c-shared -o dists/android/x86/libsharedfrp.so client.go server.go common.go

$env:CC="$TOOL_CHAINS_HOME\x86_64-linux-android21-clang.cmd"
$env:CXX="$TOOL_CHAINS_HOME\x86_64-linux-android1-clang++.cmd"
$env:GOARCH="amd64"
go build -ldflags="-s -w" -buildmode=c-shared -o dists/android/x86_64/libsharedfrp.so client.go server.go common.go

##
# Windows
##
$env:GOOS="windows"
$TOOL_CHAINS_HOME="C:\opt\mingw-w64\i686-8.1.0-win32-dwarf-rt_v6-rev0\mingw32\bin"
$env:CC="$TOOL_CHAINS_HOME\gcc.exe"
$env:CXX="$TOOL_CHAINS_HOME\g++.exe"
$env:GOARCH="386"
go build -ldflags="-s -w" -buildmode=c-shared -o dists/windows/x86/sharedfrp.dll client.go server.go common.go

$TOOL_CHAINS_HOME="C:\opt\mingw-w64\x86_64-8.1.0-win32-seh-rt_v6-rev0\mingw64\bin"
$env:CC="$TOOL_CHAINS_HOME\gcc.exe"
$env:CXX="$TOOL_CHAINS_HOME\g++.exe"
$env:GOARCH="amd64"
go build -ldflags="-s -w" -buildmode=c-shared -o dists/windows/x86_64/sharedfrp.dll client.go server.go common.go
