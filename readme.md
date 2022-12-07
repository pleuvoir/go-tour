

## 文档
https://go.dev/doc/effective_go#maps


## GO ROOT 和 GO PATH

GO ROOT是程序的安装包路径
GO PATH是工作目录，以后build什么的文件会在这里

go install的东西就会进入到 go path的bin里

如果环境变量已经配置了，那么go install的可执行文件就等于是可以直接执行的


vim ~/.bash_profile
source ~/.bash_profile

```shell
export GO_ROOT=/Users/pleuvoir/sdk/go1.19.2
export PATH=$PATH:$GO_ROOT/bin

export GO_PATH=/Users/pleuvoir/go
export PATH=$PATH:/$GO_PATH/bin


export GO_PROTO=/Users/pleuvoir/dev/support/protoc-3.20.2-osx-x86_64
export PATH=$PATH:/$GO_PROTO/bin
```

## main函数

**不要自己 go mod init main**

只能在 main package 下执行

一个目录下不能出现多个 package，即相同目录下所有文件的 package必须相同


## 包与模块

包建议和module的名称一致，一个文件夹就是一个目录。同一个文件夹下不同文件内的方法名称不能一样，也可以理解为文件夹下的所有文件是一个类。
import的永远是文件夹，因此保持文件夹个module名称一样很容易定位代码。

## 函数和方法的区别

1. 函数可以没有接受者
2. 方法必须有接受者，一个类型


可以理解为方法就是面向对象操作，函数是静态调用。

### 方法

```go
type User struct {
	name string `json:"name"`
	age  int    `json:"age"`
}

func (u *User) SetName(name string) {
	u.name = name
}

```

### 函数

```go
func main() {
    //...
}
```


## 三方包

在这里搜索包 类似Maven [https://pkg.go.dev/](https://pkg.go.dev/)


## 快捷键

option + enter 自动修改错误