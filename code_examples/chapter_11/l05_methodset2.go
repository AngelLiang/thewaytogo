// methodset2.go
/*
Topic: Using method sets with interfaces
*/
package main

import (
	"fmt"
)

// 创建一个包含 int 元素的 List 类型
type List []int

// 添加 Len 方法
// defined on a value
func (l List) Len() int { return len(l) }

// 添加 Append 方法
// defined for a pointer
func (l *List) Append(val int) { *l = append(*l, val) }

// Appender 接口
type Appender interface {
    Append(int)
}

func CountInto(a Appender, start, end int) {
    for i := start; i <= end; i++ {
        a.Append(i)
    }
}

type Lener interface {
    Len() int
}

func LongEnough(l Lener) bool {
    return l.Len()*10 > 42
}

func main() {
    // A bare value
    var lst List
    // compiler error:
    // cannot use lst (type List) as type Appender in function argument:
    // List does not implement Appender (Append method requires pointer receiver)
    // CountInto(lst, 1, 10) // INVALID: Append has a pointer receiver
    
    if LongEnough(lst) {  // VALID: Identical receiver type
        fmt.Printf(" - lst is long enough")
    }

    // A pointer value
    plst := new(List)
    CountInto(plst, 1, 10) // VALID: Identical receiver type
    if LongEnough(plst) {  // VALID: a *List can be dereferenced for the receiver
        fmt.Printf(" - plst is long enough")  //  - plst is long enoug
    }      
}
