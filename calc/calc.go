package calc

type Calc struct {
	name string
}

func (c Calc) Add(a int, b int) int{
	return a + b
}

func (c Calc) Sub(a int, b int) int{
	return a - b
}
