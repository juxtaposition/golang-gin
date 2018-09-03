package routerEze

type Person struct{
	name string
	age int
}
 
func (e *Person) SayHello(greet string) string{
	return greet
}
