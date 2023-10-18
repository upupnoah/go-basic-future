package types

type Fish struct {
}

func (f Fish) Swim() {
	println("鱼在游")
}

// Yu 鱼
// 鱼是 Fish 的别名
type Yu = Fish

type FakeFish struct {
}

func (f FakeFish) FakeSwim() {
	println("假的鱼在游")
}

func UseFish() {
	f1 := Fish{}
	f1.Swim()
	f2 := FakeFish{}
	// f2 将不能调用 Fish 上的方法，
	// 因为 f2 是一个全新的类型
	f2.FakeSwim()

	// 类型转换
	f3 := Fish(f2)
	f3.Swim()

	y := Yu{}
	y.Swim()
}
