// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef YJYY_H_
#define YJYY_H_

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

extern char* YjyyGetDataVersion();
extern char* YjyySearch(char* query, int32_t limits);

#ifdef __cplusplus
}
#endif

#endif // YJYY_H_
