package wovert

type Retriever struct {
	Contents string
}


// type Retriever interface {
// 	Get(url string) string
// }
func (r Retriever) Get(url string) string {
	return r.Contents
}