package filelist

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//定义userError类型，两边的userError自己管自己的
type userError string

func (e userError) Error() string{
	return e.Message()
}

func (e userError) Message() string{
	return string(e)
}

const prefix = "/list/"
func HandleFileList (writer http.ResponseWriter, request *http.Request) error{
	if strings.Index(request.URL.Path,prefix) != 0{
		return userError("Path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]	//将list前缀给去掉，拿到相对路径
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer,err.Error(),http.StatusInternalServerError)
		//return
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	writer.Write(all)
	return nil
}
