
https://go.dev/doc/effective_go#maps


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