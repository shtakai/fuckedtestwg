package compare

func Square(i int) int {
	return i * i
}

// Dog is used a demo type
type Dog struct {
	Name string
	Age  int
	//fn   func()
	//thing struct {
	//	Name string
	//	Age  string
	//}
}

// DogWithFn is used a fucked up demo type
type DogWithFn struct {
	Name string
	Age  int
	Fn   func()
}
