package basic
import (
	"net/http"
)
type Request struct {
	httpReq *http.Request
	index int
}
func NewRequest(httpReq *http.Request, index int) *Request{
	return &Request{httpReq:httpReq, index:index}
}
func (req *Request) GetReq() *http.Request{
	return req.httpReq
}
func (req *Request) GetIndex() int{
	return req.index
}