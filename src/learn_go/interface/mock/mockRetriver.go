package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string { //并不需要显示声明实现哪个接口，实现其Get方法即可。注意接收者类型的Retriever是上面struct的，并非是接口的
	return r.Contents
}
