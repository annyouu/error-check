package mypkg

import "context"


// 検出されるやつ
type S1 struct {
	ctx context.Context
}

type S2 struct {
	Ctx context.Context
}

// されない例
type S3 struct {
	name string
	age int
}


type S4 struct {
	context.Context
}