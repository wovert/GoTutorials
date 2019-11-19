package wovert

type Retriever struct {
	Contents string
}

func (r Retriever) Post(url string,
  form map[string]string) string {
	  r.Contents = form["contents"]
	  return "ok"
  }


// type Retriever interface {
// 	Get(url string) string
// }
func (r Retriever) Get(url string) string {
	return r.Contents
}