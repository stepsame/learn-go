package main

// 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。

func getValue() (ret string) {
	defer func() {
		r := recover()
		if r != nil {
			println("i am recover from panic ", r)
			ret = "wtf"
		}
	}()
	panic("haha")
}

func main() {
	println(getValue())
}
