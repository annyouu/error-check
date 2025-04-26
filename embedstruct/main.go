package main

import "fmt"

// 埋め込まれる型
type Profile struct {
	Name string
	Age int
}

// メソッドも定義する
func (p Profile) Greet() {
	fmt.Println("こんにちは! 私の名前は", p.Name)
}

// 埋め込む側の型(親)
type User struct {
	Profile // 埋め込みフィールド
	Email string
	IsActive bool
}

// 実行
func main() {
	// User型を作成
	u := User{
		Profile: Profile{
			Name: "Alice",
			Age: 28,
		},
		Email: "alice@example.com",
		IsActive: true,
	}

	fmt.Println("名前:", u.Name)
	fmt.Println("年齢:", u.Age)
	fmt.Println("メール:", u.Email)
	fmt.Println("アクティブ:", u.IsActive)

	// 埋め込まれた型のメソッドも呼ぶ。
	u.Greet()
}