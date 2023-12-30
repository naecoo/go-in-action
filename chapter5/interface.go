package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func curl(url string) {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./interface <url>")
		os.Exit(-1)
	}

	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

func bufferTest() {
	var b bytes.Buffer
	b.Write([]byte("Hello"))
	fmt.Fprintf(&b, " World!\n")
	io.Copy(os.Stdout, &b)
}

type notifier interface {
	notify()
}

type user1 struct {
	name  string
	email string
}

func (u *user1) notify() {
	fmt.Printf("name: %s, email: %s\n", u.name, u.email)
}

type user2 struct {
	name  string
	email string
}

func (u user2) notify() {
	fmt.Printf("name: %s, email: %s\n", u.name, u.email)
}

type admin struct {
	name  string
	email string
}

func (u *admin) notify() {
	fmt.Printf("name: %s, email: %s\n", u.name, u.email)
}

func sendNotify(n notifier) {
	n.notify()
}

func userTest() {
	var u1 notifier
	// 如果用了指针接受者，这里也必须用指针，方法才能兼容接口
	u1 = &user1{"user1", "hello@mail.com"}
	sendNotify(u1)
	// 使用值接受者，指针和值都可以
	var u2, u3 notifier
	u2 = user2{"user2", "hello@mail.com"}
	sendNotify(u2)
	u3 = &user2{"user3", "hello@mail.com"}
	sendNotify(u3)

	// 多态
	a1 := &admin{"dd", "admin@mail.com"}
	sendNotify(u1)
	sendNotify(a1)
}

type tt struct {
	u user1
	admin
	level uint8
}

func main() {
	// curl(os.Args[1])
	bufferTest()
	userTest()

	t1 := tt{
		u:     user1{"t1", "t1@mail.com"},
		admin: admin{"t1", "a1@mail.com"},
		level: 1,
	}
	t1.notify()
}
