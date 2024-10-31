# Important
> [!IMPORTANT]  
> 研究了两天，发现 golang 在桌面这边就真的一般。
> 我开始使用的是 golang.design/x/clipboard, 发现用起来并不舒服[Issue](https://github.com/golang-design/clipboard/issues/66)
> 
> 后面我就去查找相关其他的库，在reddit上找到了这篇[帖子](https://www.reddit.com/r/commandline/comments/1dfhu2o/introducing_clipper_a_handy_commandline_clipboard/)
> 看了下它使用的库，找到了这个[github仓库](https://github.com/atotto/clipboard)
> 发现就是调用的 xclip 或者 wl-clip .....
> 所以，这个项目放弃了，可以说是了解下吧，了解了下剪切板之类的内容。

> [!IMPORTANT]  
> After two days of research, I found that Golang’s capabilities for desktop applications are quite limited. Initially, I used golang.design/x/clipboard, but it didn’t feel very user-friendly [Issue](https://github.com/golang-design/clipboard/issues/66).
> Later, I looked into other libraries and found this [post](https://www.reddit.com/r/commandline/comments/1dfhu2o/introducing_clipper_a_handy_commandline_clipboard/) on Reddit. From there, I discovered the library it used and found this [GitHub repository](https://github.com/atotto/clipboard), which essentially relies on calling xclip or wl-clip.
> So, I’ve decided to drop this project. Let’s just say I’ve learned more about clipboard functionality through this exploration.
# cl
clipboard cli utils

