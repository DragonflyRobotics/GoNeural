package ops


struct Add {
	a: i32,
	b: i32,
	op: Operation,
}

func ConstructAdd(a: i32, b: i32) Add {
	return Add{a: a, b: b, op: op.Operation{name: "Add", unary: false}}
}
