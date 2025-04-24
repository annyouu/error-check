package target

// 値レシーバでError()を定義
type MyError struct{}

func (MyError) Error() string {
	return "my error"
}

// ポインタレシーバでError()を定義
type PtrError struct{}

func (*PtrError) Error() string {
	return "ptr error"
}

// Error()メソッドを持たない型
type NoError struct{}

// その他メソッドを持つ型
type Ohter struct{}

func (Ohter) Foo() {}

