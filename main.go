package linter

type Receiver struct{}

func (receiver Receiver) Hello() { _ = receiver }
