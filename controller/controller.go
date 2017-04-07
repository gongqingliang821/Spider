package controller

import (
	//"fmt"
	"../config"
	"../download"
	"../analy"
	"../basic"
	"../process"
	"fmt"
	"sync"
	"net/http"
)
var wg sync.WaitGroup
type Controller struct {
	StartUrl string
	Depth  int
	Downloader download.GenDownloader  //下载组件
	Analyer analy.GenAnalyer   //页面分析组件  获取url以及文章title
	Process process.GenProcess
	Reqchan chan *basic.Request
	Reschan chan *basic.Response
	NextUrlchan chan basic.NextUrl
	Itemchan chan map[string]interface{}
}
func NewController(StartUrl string, Depth int) *Controller{
	return &Controller{StartUrl:StartUrl, Depth:Depth}
}
func (crt *Controller)Do() {
	config.InitConfig()
	crt.Downloader = download.NewDownload()
	crt.Analyer = analy.NewAnalyer()
	crt.Process = process.NewProcess()
	crt.Reqchan = make(chan *basic.Request, config.Config.Reqlen)
	crt.Reschan = make(chan *basic.Response, config.Config.Reqlen)
	crt.NextUrlchan = make(chan basic.NextUrl, config.Config.Urllen)
	crt.Itemchan = make(chan map[string]interface{}, config.Config.Itemlen)
	beginReq, _ := http.NewRequest(config.Config.ResquestMethod, crt.StartUrl, nil)
	baseReq := basic.NewRequest(beginReq, 0)
	crt.Reqchan <- baseReq
	wg.Add(3)
	go crt.DownloadManger()
	go crt.AnalyManager()
	go crt.ProcessManger()
	wg.Wait()
	crt.ShowTitle()
	fmt.Println("end")
}
func (crt *Controller)DownloadManger() {
	defer wg.Done()
	for req := range crt.Reqchan {
		res := crt.Downloader.Download(req)
		if(res != nil){
			crt.Reschan <- res
		}
	}
	fmt.Println("DownloadManger End")
	close(crt.Reschan)
}
func (crt *Controller)AnalyManager(){
	defer wg.Done()
	for res := range crt.Reschan{
		linklist,itemlist := crt.Analyer.Analy(res.GetRes())
		for _,link := range(linklist){
			NextUrl := basic.NewNextUrl(link, res.GetIndex() + 1)
			crt.NextUrlchan <- NextUrl
		}
		for _,item := range(itemlist){
			crt.Itemchan <- item
		}
	}
	close(crt.NextUrlchan)
	close(crt.Itemchan)
	fmt.Println("AnalyManager End")
}

func (crt *Controller)ProcessManger(){
	defer wg.Done()
	cnt := 0
	for NextUrl := range(crt.NextUrlchan){
		req,flag := crt.Process.Deallink(NextUrl)
		if !flag {
			continue
		}
		if(req.GetIndex() >= crt.Depth){
			close(crt.Reqchan)
			break
		}else{
			crt.Reqchan <- req
		}
		fmt.Println(NextUrl)
		cnt = cnt + 1
	}

	if(len(crt.NextUrlchan) !=0 ){
		for NextUrlchan := range(crt.NextUrlchan){
			fmt.Println(NextUrlchan)
			cnt = cnt + 1
		}
	}
	fmt.Println(cnt)
	fmt.Println("ProcessManger End")
}
func (crt *Controller)ShowTitle(){
	fmt.Println("显示的上一层的title")
	for value := range (crt.Itemchan){
		fmt.Println(value)
	}
}
