#!/bin/sh
# Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# https://github.com/golang/go/issues/11339
# https://github.com/golang/go/issues/12683
# https://groups.google.com/forum/#!topic/golang-dev/738LXeykFsM

set -e

PWD=`pwd`

go fmt github.com/chai2010/ptyy/cmd/yjyy

# ---------------------------------------------------------
# iPhoneSimulator: 386/amd64
# ---------------------------------------------------------

export CGO_ENABLED=1
export GOARCH=386
#export CC=$PWD/clangwrap.sh
#export CXX=$PWD/clangwrap.sh

go build -buildmode=c-archive -o libyjyy_386.a github.com/chai2010/ptyy/cmd/yjyy

export CGO_ENABLED=1
export GOARCH=amd64
#export CC=$PWD/clangwrap.sh
#export CXX=$PWD/clangwrap.sh

go build -buildmode=c-archive -o libyjyy_adm64.a github.com/chai2010/ptyy/cmd/yjyy

# ---------------------------------------------------------
# arm64
# ---------------------------------------------------------

export CGO_ENABLED=1
export GOARCH=arm64
export CC=$PWD/clangwrap.sh
export CXX=$PWD/clangwrap.sh

go build -buildmode=c-archive -o libyjyy_arm64.a github.com/chai2010/ptyy/cmd/yjyy

# ---------------------------------------------------------
# armv7
# ---------------------------------------------------------

export CGO_ENABLED=1
export GOARCH=arm
export GOARM=7
export CC=$PWD/clangwrap.sh
export CXX=$PWD/clangwrap.sh

go build -buildmode=c-archive -o libyjyy_armv7.a github.com/chai2010/ptyy/cmd/yjyy

# ---------------------------------------------------------
# Merge Arm64 and ArmV7
# ---------------------------------------------------------

# Make universal library
lipo libyjyy_386.a libyjyy_adm64.a libyjyy_arm64.a libyjyy_armv7.a -create -output libyjyy.a
rm   libyjyy_386.a libyjyy_adm64.a libyjyy_arm64.a libyjyy_armv7.a
rm   libyjyy_386.h libyjyy_adm64.h libyjyy_arm64.h libyjyy_armv7.h

# ---------------------------------------------------------
# END
# ---------------------------------------------------------
