package basic
import (
	"net/http"
)
type Response struct {
	response *http.Response
	index int
}
func NewResponse(response *http.Response, index int) *Response{
	return &Response{response:response, index:index}
}
func (self *Response)GetRes() *http.Response{
	return self.response
}
func (self *Response)GetIndex() int{
	return self.index
}
