#!/bin/sh
# Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# https://github.com/golang/go/issues/11339
# https://github.com/golang/go/issues/12683
# https://groups.google.com/forum/#!topic/golang-dev/738LXeykFsM

set -e

PWD=`pwd`

go fmt

# ---------------------------------------------------------
# iPhoneSimulator: 386/amd64
# ---------------------------------------------------------

export CGO_ENABLED=1
export GOARCH=386
#export CC=$PWD/clangwrap.sh
#export CXX=$PWD/clangwrap.sh

go build -buildmode=c-archive -o libgo_386.a

export CGO_ENABLED=1
export GOARCH=amd64
#export CC=$PWD/clangwrap.sh
#export CXX=$PWD/clangwrap.sh

go build -buildmode=c-archive -o libgo_adm64.a

# ---------------------------------------------------------
# arm64
# ---------------------------------------------------------

export CGO_ENABLED=1
export GOARCH=arm64
export CC=$PWD/clangwrap.sh
export CXX=$PWD/clangwrap.sh

go build -buildmode=c-archive -o libgo_arm64.a

# ---------------------------------------------------------
# armv7
# ---------------------------------------------------------

export CGO_ENABLED=1
export GOARCH=arm
export GOARM=7
export CC=$PWD/clangwrap.sh
export CXX=$PWD/clangwrap.sh

go build -buildmode=c-archive -o libgo_armv7.a

# ---------------------------------------------------------
# Merge Arm64 and ArmV7
# ---------------------------------------------------------

# Make universal library
lipo libgo_386.a libgo_adm64.a libgo_arm64.a libgo_armv7.a -create -output libgo.a
rm   libgo_386.a libgo_adm64.a libgo_arm64.a libgo_armv7.a
rm   libgo_386.h libgo_adm64.h libgo_arm64.h libgo_armv7.h

# ---------------------------------------------------------
# END
# ---------------------------------------------------------
