package main

import (
	"fmt"
)

type A struct {
	Text string
	Name string
}

func (a *A) Say() {
	fmt.Printf("A::Say():%s\n", a.Text)
}

type B struct {
	A
	Name string
}

func (b *B) Say() {
	b.A.Say()
	fmt.Printf("B::Sayaaa():%s\n", b.Text)
}

func (b *B) Sayaaa() {
	b.A.Say()
	fmt.Printf("B::Say():%s\n", b.Text)
}

func main() {
	var aaa int32 = -1
	if aaa < 0 {
		fmt.Println(aaa)
	}
	return
	b := B{A{"hello, world", "张三"}, "李四"}

	b.Say()
	fmt.Println("子类的名字为：", b.Name)

	fmt.Println("父类的名字为：", b.A.Name)
}
