package model

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetName() string {
	return u.Name
}

func Print() {
	println("SDK中的代码被调用了。")
}
