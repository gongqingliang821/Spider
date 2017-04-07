package analy

import (
	"../basic"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type GenAnalyer interface {
	Analy(httpres *http.Response)([]string, []map[string]interface{})
}

//解析请求函数
func parse(httpres *http.Response)([]string, []map[string]interface{}){
	linklist := make([]string, 0);
	itemlist := make([]map[string]interface{}, 0)
	doc, _ := goquery.NewDocumentFromResponse(httpres)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exits := s.Attr("href")
		if exits {
			link = basic.CheckLink(link)
			if link != "" {
				linklist = append(linklist, link)
			}
		}
	})
	//保存每个页面的标题
	title := strings.TrimSpace(doc.Find("head title").Text())
	if title != "" {
		item := make(map[string]interface{})
		item["标题"] = title
		itemlist = append(itemlist, item)
	}
	return linklist, itemlist

}
type Analyer struct {
	linklist []string
	itemlist []map[string]interface{}
}
func NewAnalyer() *Analyer{
	return &Analyer{make([]string,0),make([]map[string]interface{},0)}
}
func (self *Analyer) Analy(httpres *http.Response)([]string, []map[string]interface{}){
	defer httpres.Body.Close()
	return parse(httpres)
}