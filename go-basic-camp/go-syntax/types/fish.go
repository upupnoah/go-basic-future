package types

type Fish struct {
	Name string
}

func (f *Fish) Add(index int, val any) {
	//TODO implement me
	panic("implement me")
}

func (f *Fish) Append(val any) {
	//TODO implement me
	panic("implement me")
}

func (f *Fish) Delete(index int) {
	//TODO implement me
	panic("implement me")
}

func (f *Fish) Swim() {
	println("I am swimming...")
}

// 扩展 Fish，但是不想让 Fish 的 Swim 方法被继承
// 可以访问 Fish 的导出字段

type FakeFish Fish
