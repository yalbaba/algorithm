package StockToQueue

import (
	"fmt"
	"github.com/yalbaba/DataStructures/Stack"
)

type myqueue struct {
	a *Stack.Stack
	b *Stack.Stack
}

func NewQu() *myqueue {
	return &myqueue{
		a: Stack.New(),
		b: Stack.New(),
	}
}

func (n *myqueue) Push(i int) {
	n.a.Push(i)
}

func (n *myqueue) Remove() int {
	for n.a.Empty() != true {
		i, _ := n.a.Pop()
		n.b.Push(i)
	}
	j, _ := n.b.Pop()
	return j.(int)
}

func main() {
	q := NewQu()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Remove())
}
