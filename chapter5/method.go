package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// 必须使用指针接受者，不然只是修改副本，调用方感知不到
func (u *user) changeEmail(email string) {
	u.email = email
}

// 嵌套struct测试
type du struct {
	user
}

// 使用值接受者只能修改内部的副本
func (d du) changeEmail(email string) {
	d.user.email = email
	d.user.notify()
}

func main() {
	// bill := user{"Bill", "bill@gmail.com"}
	// bill.notify()

	// lisa := &user{"Lisa", "lisa@gmail.com"}
	// lisa.notify()

	// bill.changeEmail("bill-new-email@gmail.com")
	// bill.notify()

	// lisa.changeEmail("lisa-new-email@gmail.com")
	// lisa.notify()

	era := du{
		user: user{"dd", "dd@gmail.com"},
	}
	era.user.notify()
	era.changeEmail("d123@gmail.com")
	era.user.notify()
}
