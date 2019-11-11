package queue

// type Queue []int
type Queue []interface{}

// func (q *Queue) Push(v interface{}) {
// 	*q = append(*q, v)
// }

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int)) // 限定类型
}

// func (q *Queue) Pop() interface{} {
// 	head := (*q)[0]
// 	*q = (*q)[1:]
// 	return head
// }
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int) // 限定类型
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}