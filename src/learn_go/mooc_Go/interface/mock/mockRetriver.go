package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf(
		"Retriever: {Contents=%s}",r.Contents)
}

func (r *Retriever) Get(url string) string { //并不需要显示声明实现哪个接口，实现其Get方法即可。注意接收者类型的Retriever是上面struct的，并非是接口的
	return r.Contents
}

func (p *Retriever) Post(url string, form map[string]string) string{	//实现的是post接口，因此此处的Retriever只是这个文件的
	p.Contents = form["contents"]
	return "Ok"
}

