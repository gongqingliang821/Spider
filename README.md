# Spider
参考网络爬虫，用go语言简单写了个爬虫，用来练手，也可以用来爬取一个主题网页下面的连接.


运行方法直接 go run main.go


相关配置文件在config/config.go下面



想爬取那个网页直接修改main.go文件 

大概思路：

1、download组件负责发起网络请求

2、analy组件负责分析附加的连接，及网页的title

3、process组件负责构造一个新的网页请求

4、basic下面定义的基本请求、响应等数据结构

5、每个组件功能分别起一个协程来处理，通过chan来传递消息
       Reqchan chan *basic.Request  
       Reschan chan *basic.Response  	
       NextUrlchan chan basic.NextUrl
	     Itemchan chan map[string]interface{}
       
6.引用第三方库github.com/PuerkitoBio/goquery
       
