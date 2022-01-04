package main
import (
	"fmt"
	"strconv"
	"errors"
)

func div(n1, n2 int) (int, error) {
	if n2 == 0 {
		return -1, errors.New("除数为0")
	}
	return n1 / n2, nil
}

func main() {
	value, err := strconv.Atoi("xxx")
	fmt.Printf("%T\n", err)
	fmt.Println(err)
	fmt.Println(value)

	e := fmt.Errorf("自定义错误")
	fmt.Printf("%T, %#v\n", e, e)

	e2 := errors.New("自定义错误")
	fmt.Printf("%T, %#v\n", e2, e2)

	if result, err := div(1, 0); err == nil {
		fmt.Println("result=", result)
	} else {
		fmt.Println("error=", err)
	}



}