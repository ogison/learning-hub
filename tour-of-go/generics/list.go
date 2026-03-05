package main

import "fmt"

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Push adds a new value at the end of the list
func (l *List[T]) Push(val T) {
	node := l
	// 末尾まで移動
	for node.next != nil {
		node = node.next
	}
	node.next = &List[T]{val: val}
}

// Print outputs all values in the list
func (l *List[T]) Print() {
	for node := l.next; node != nil; node = node.next { // l はダミーヘッド
		fmt.Println(node.val)
	}
}

func main() {
	// ダミーヘッドで初期化
	list := &List[int]{}

	// 値を追加
	list.Push(1)
	list.Push(2)
	list.Push(3)

	// 出力
	list.Print()
}
