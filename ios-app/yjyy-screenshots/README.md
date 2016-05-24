# 野鸡医院 图片资源

截屏中因为涉及到微信分享的截图, 无法用模拟器截屏.
因此, 选用了iPod和iPadPro真机截屏, 分别用于各个尺寸的iPhone和iPad预览图.
其中, iPhone4s和其它新的iPhone型号的长宽比例稍有不同.
而iPad和iPadPro长宽比相同, 但是屏幕大小相差1倍, 因此显示的元素布局会稍有不同.

不过, 预览图只是一个示意图, 并不需要和每个型号的真机完全相同,
所以只选用了iPod和iPadPro两个型号的真机截屏来代替其它的设备.

目录结构:

- yjyy-appstore.png: AppStore安装地址的二维码
- yjyy-logo@1024x1024.png: 原始的logo图标
- raw目录: 原始的截屏图片, 含ps之后的预览图
- output: 各种尺寸的输出图片, 含logo和预览图, 不要保存

外部的 yjyy-swift 和 yjyy-ruby 工程不要直接引用当前目录的资源,
如果需要请复制1份到自己的过程目录.
