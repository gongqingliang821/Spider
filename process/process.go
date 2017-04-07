package process

import (
	"../basic"
	"../config"
	//"fmt"
	"net/http"
)

type GenProcess interface {
	Deallink(nextUrl basic.NextUrl)(*basic.Request,bool)
}
type Process struct {
	Vurl map[string]bool
}
func NewProcess() *Process{
	return &Process{make(map[string]bool)}
}
func(self *Process)Deallink(nextUrl basic.NextUrl)(*basic.Request,bool){
	url := nextUrl.GetUrl()  //获取要访问的url
	if _, visited := self.Vurl[url]; visited{
		return nil, false   //访问过直接返回
	}
	httpReq, err := http.NewRequest(config.Config.ResquestMethod, url, nil)
	if(err != nil){
		return nil,false
	}
	request := basic.NewRequest(httpReq, nextUrl.GetIndex()) //转化为构造请求
	self.Vurl[url] = true
	return request, true
}