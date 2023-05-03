package customtypes

type PersonOrEmployee interface {
	Person | Employee
}
type SlicePersonOrEmployee interface {
	[]Person | []Employee
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Name     string
	Position string
}