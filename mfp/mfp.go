package mfp

import "fmt"

func PrintFmtVal(str string, v any, verbs []string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发现错误：\t", err)
		}
	}()

	str += ": \t"
	var vs []any
	for _, verb := range verbs {
		vs = append(vs, v)
		str += "%%" + verb + " -> %" + verb + " | "
	}
	str += "\n"
	fmt.Printf(str, vs...)
}
