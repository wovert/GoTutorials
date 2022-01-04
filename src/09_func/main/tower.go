package main


import (
	"fmt"
)

// 汉诺塔游戏
// n个盘子 start() end() temp()
// n-1 start -> end -> temp
// start -> end
// n-1 temp -> start -> end
// 终止条件 1 -> start -> end

func tower(start string, end string, temp string, layer int) {
	if (layer == 1) {
		fmt.Println(start, "->", end)
		return
	}
	tower(start, temp, end, layer-1)
	fmt.Println(start, "->", end)
	tower(temp, end, start, layer -1)
}

func main() {
	tower("塔1","塔3","塔2",3)
}