package person

var Name string = "wovert"
var Age int = 99

// func Test() {
// 	Name = "hello world"
// 	Age = 10
// }

func init() {
	Name = "person init function"
	Age = 100
}