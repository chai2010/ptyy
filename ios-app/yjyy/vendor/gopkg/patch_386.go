// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 针对iOS模拟器link时缺少的函数
// 属于临时解决方案

package main

/*
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <time.h>

size_t fwrite$UNIX2003(const void* a, size_t b, size_t c, FILE* d) {
    return fwrite(a, b, c, d);
}

char* strerror$UNIX2003(int errnum) {
    return strerror(errnum);
}

time_t mktime$UNIX2003(struct tm * a) {
    return mktime(a);
}
double strtod$UNIX2003(const char * a, char ** b) {
    return strtod(a, b);
}

int setenv$UNIX2003(const char* envname, const char* envval, int overwrite) {
    return setenv(envname, envval, overwrite);
}
int unsetenv$UNIX2003(const char* name) {
    return unsetenv(name);
}

*/
import "C"
