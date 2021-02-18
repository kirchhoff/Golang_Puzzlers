package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	s := list.New()
	for i := 0; i < 10; i++ {
		s.PushBack(i)
	}

	for i := s.Front(); i != nil; i = i.Next() {
		fmt.Printf("%T %v\n", i.Value, i.Value)
	}

	var r ring.Ring
	fmt.Println(r.Len())

	var badmap3 = map[interface{}]int{
		[1]int{1}: 2,
	}

	fmt.Println(badmap3[[1]int{1}])

}
