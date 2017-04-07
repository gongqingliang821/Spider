package basic

import (
	"net/url"
	"strings"
	"../config"
)
func CheckLink(link string) string{
	u, err := url.Parse(link)
	if(err != nil){
		return ""
	}
	if u.Scheme == "" {
		return ""
	}
	if strings.EqualFold(u.Scheme, "http") ||  strings.EqualFold(u.Scheme, "https"){
		return link
	}
	if flag := strings.HasPrefix(link, config.Config.StartUrl); flag != true {
		link = strings.Join([]string{config.Config.StartUrl, link}, "")
		return link
	}
	return ""
}
