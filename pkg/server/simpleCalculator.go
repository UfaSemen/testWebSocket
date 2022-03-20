package server

//basic calculator for sum and product
type SimpleCalculator struct{}

//sum function
func (sc SimpleCalculator) Sum(n1, n2 int) int {
	return n1 + n2
}

//product function
func (sc SimpleCalculator) Product(n1, n2 int) int {
	return n1 * n2
}
