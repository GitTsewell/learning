* [Go 语言编译器的 "//go:" 详解](https://segmentfault.com/a/1190000016743220)
* [go语言高级编程——–汇编语言部分学习笔记](https://www.codenong.com/cs106480140/)
* [阮一峰-汇编语言入门教程](http://www.ruanyifeng.com/blog/2018/01/assembly-language-primer.html)


## 编译原理
Go 的编译器在逻辑上可以被分成四个阶段：词法与语法分析、类型检查和 AST 转换、通用 SSA 生成和最后的机器代码生成

go tool compile -S main.go 查看调用
