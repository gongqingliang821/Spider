package basic
type NextUrl struct {
	link string
	index int
}
func NewNextUrl(link string, index int) NextUrl{
	return NextUrl{link, index}
}
func (self NextUrl)GetUrl()string{
	return self.link
}
func (self NextUrl)GetIndex()int{
	return self.index
}
