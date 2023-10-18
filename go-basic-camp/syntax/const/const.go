package _const

const internal = "包内可访问"
const External = "包外可访问"

func Const() {
	const a = "你好"
	println(a)
}

const (
	Status0 = iota
	Status1
	Status2
	Status3

	// 不管隔多少行，都是接着4

	Status4

	// 插入一个主动赋值的就中断了 iota

	Status6 = 6
	Status7
)

const (
	One = iota << 1
	Two
	Four
)
