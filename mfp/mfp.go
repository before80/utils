package mfp

import (
	"fmt"
	"reflect"
	"strconv"
)

func PrintHr() {
	fmt.Println("------------------")
}

func loopVerbs(verbs []string) (str string) {
	str += ": \t"
	for k, verb := range verbs {
		if k > 0 {
			str += " | %%" + verb + " -> %" + verb
		} else {
			str += "%%" + verb + " -> %" + verb
		}
	}
	return str
}

func PrintFmtVal(str string, v any, verbs []string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发现错误：\t", err)
		}
	}()

	str += loopVerbs(verbs) + "\n"
	var vs []any
	for len(vs) < len(verbs) {
		vs = append(vs, v)
	}
	fmt.Printf(str, vs...)
}

func PrintFmtValWithLC(str string, v any, verbs []string) {
	kind := reflect.TypeOf(v).Kind()
	if !(kind == reflect.Slice || kind == reflect.Array) {
		panic("非数组或切片不能调用该函数：PrintFmtValWithLC")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发现错误：\t", err)
		}
	}()

	str += loopVerbs(verbs)
	var vs []any
	for len(vs) < len(verbs) {
		vs = append(vs, v)
	}

	vv := reflect.ValueOf(v)
	str += " | len=" + strconv.Itoa(vv.Len())
	str += " | cap=" + strconv.Itoa(vv.Cap())
	str += "\n"
	fmt.Printf(str, vs...)
}

func PrintFmtValWithL(str string, v any, verbs []string) {
	kind := reflect.TypeOf(v).Kind()
	if !(kind == reflect.Slice || kind == reflect.Array || kind == reflect.Chan || kind == reflect.Map || kind == reflect.String) {
		panic("非数组、切片、通道、map、字符串不能调用该函数：PrintFmtValWithL")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发现错误：\t", err)
		}
	}()

	str += loopVerbs(verbs)
	var vs []any
	for len(vs) < len(verbs) {
		vs = append(vs, v)
	}

	vv := reflect.ValueOf(v)
	str += " | len=" + strconv.Itoa(vv.Len())
	str += "\n"
	fmt.Printf(str, vs...)
}
