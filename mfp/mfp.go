package mfp

import (
	"fmt"
)

func PrintHr() {
	fmt.Println("------------------")
}

func PrintFmtVal(str string, v any, verbs []string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发现错误：\t", err)
		}
	}()

	str += ": \t"
	var vs []any
	for k, verb := range verbs {
		vs = append(vs, v)
		if k > 0 {
			str += " | %%" + verb + " -> %" + verb
		} else {
			str += "%%" + verb + " -> %" + verb
		}
	}
	str += "\n"
	fmt.Printf(str, vs...)
}

//func PrintSl(str string, sl any) {
//	t := reflect.TypeOf(sl)
//	fmt.Println("str=", sl, ",长度=", len(sl), ",容量=", cap(sl))
//}
