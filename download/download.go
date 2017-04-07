package download

import (
	"net/http"
	"../basic"
	//"fmt"
)
type GenDownloader interface {
	Download(req *basic.Request) *basic.Response
}

type Download struct {
	Client *http.Client
}
func NewDownload() *Download{
	return &Download{&http.Client{}}
}
func (self *Download)Download(req *basic.Request) *basic.Response{
	httpRes, err := self.Client.Do(req.GetReq())
	//fmt.Println(err)
	if(err != nil){
		return nil
	}
	response := basic.NewResponse(httpRes, req.GetIndex())
	//fmt.Println(response)
	return response
}