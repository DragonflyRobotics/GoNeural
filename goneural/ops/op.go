package ops

struct Operation {
	string name;
	bool unary;
}

func (op *Operation) String() string {
	if op.unary {
		return fmt.Sprintf("Unary operation: %s", op.name)
	}
	return fmt.Sprintf("Binary operation: %s", op.name)
}
