# 使用Go语言开发iOS应用(Swift版)

本文加上读者对Go语言和Swift语言都有一定了解, 但是对二者混合使用不了解的同学.

## 背景

Go语言是Google公司于2010年开源的一个面向网络服务和多并发环境的编程语言，特点是简单。
但是因为简单，也就只能实现90%的性能，这是Go语言的最大优点，因为 少即是多 的道理不是每个人都能领悟的。

Swift是Apple公司于2014年发布的用来替代ObjectiveC的语言，主要面向iOS和OS X上的界面程序开发。
当然用swift来开发服务器也是大家关注的一个领域，作者看好在不远的将来Swift将逐步替代C++和Rust语言。

Go语言和Swift语言本来是风马牛不相及的两个语言，为何非一定要整到一起呢？
原因很简单，因为作者是一个Go粉，同时也算是半个Swift粉；想试水iOS开发，但是实在是受不了ObjectiveC的裹脚布语法。

补充下：本人虽然不喜欢ObjectiveC的语法，但是觉得ObjectiveC的runtime还是很强悍的。
理论上，基于ObjectiveC的runtime，可以用任何流行的编程语言来开发iOS应用，RubyMotion就是一个例子。

其实，现在流行的绝大部分语言都有一个交集，就是c语言兼容的二进制接口。
所以说，C++流行并不是C++多厉害，而是它选择几本无缝兼容了C语言的规范。

但是，完全兼容C语言的规范也有缺点，就是语言本身无法自由地发展，因为很多地方会受到C语言编程模型的限制。
C++和ObjectiveC是两个比较有代表的例子。

所以说，Swift一出世就兼容C语言的二进制接口规范，同时抱紧了ObjectiveC的runtime大腿，而去自己确实有很大优秀的特性。

但是，我们这里暂时不关心Swift和ObjectiveC的混合编程，我们只关注作为ObjectiveC子集的C语言如何与Swift混合编程。

## Swift调用C函数

Swift调用C函数的方法有多种：通过ObjectiveC桥接调用和直接调用。其实两者的原理是一样的，我个人跟喜欢选择最直接也最暴力的直接调用C函数的方式。

比如有一个C函数:

```
#include <stdio.h>

void getInput(int *output) {
    scanf("%i", output);
}
```

生成一个桥接的头文件`xxx-Bridging-Header.h`，里面包含c函数规格说明：

```c
void getInput(int *output);
```

swift就可以直接使用了:

```swift
import Foundation

var output: CInt = 0
getInput(&output)

println(output)
```

如果不用桥接文件，可以在swift中声明一个Swift函数，对应C函数:

```swift
@_silgen_name("getInput") func getInput_swift(query:UnsafePointer<CInt>)
```

为了明确区分C函数和swift函数，我们将`getInput`重新声明为`getInput_swift`，使用方法和前面一样：

```swift
import Foundation

var output: CInt = 0
getInput_swift(&output)

println(output)
```

## Swift中如何管理c返回的内存

Swift语言本身是自带ARC的，用户很少直接关注内存问题。但是C函数如果返回内存到Swift空间，
Swift的ARC是无效的，需要手工释放C内存。

假设我们自己用C语言实现了一个字符串克隆的函数:

```c
char* MyStrDup(char* s) {
	return strdup(s);
}
```

在swift中可以这样使用：

```swift
@_silgen_name("MyStrDup")
func MyStrDup_swift(query:UnsafePointer<CChar>) -> UnsafeMutablePointer<CChar>

let p = MyStrDup_swift("hello swift-c!")
let s = String.fromCString(p)!
p.dealloc(1)
```

使用`String.fromCString(p)!`从C字符串构建一个swift字符串，然后手工调用`p.dealloc(1)`释放c字符串内存空间。

函数调用和内存管理是跨语言编程中最重要的两个基础问题，目前已久初步可以工作了。

## Go语言导出C静态库

## 交叉构建的参数

## 总结

TODO
