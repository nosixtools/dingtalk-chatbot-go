package main

import (
	"dingtalk-chatbot-go/message"
	"fmt"
	"dingtalk-chatbot-go/dingtalkclient"
	"container/list"
)

var HOST = "https://oapi.dingtalk.com/robot/send?access_token=c8e46870d0cd473e778763e4613b81644aa1b32f6ac59ae0ab398d404c600552"

func main() {
	//准备数据
	//text := "hello world"
	//mobiles := list.New()
	//mobiles.PushBack("13261656404")
	////发送数据
	//fmt.Println(dingtalkclient.Send(HOST, message.GetNewTextMessage(text, mobiles, false).ToJsonString()))

	//text := "#### 特色社会主义\n 建设社会主义新中国"
	//title := "中国"
	//messageUrl := "http://www.ruanyifeng.com/blog/2014/05/oauth_2_0.html"

	//msg := message.GetNewLinkMessage(title, text, "", messageUrl).ToJsonString();
	//fmt.Println(msg)
	//fmt.Println(dingtalkclient.Send(HOST, msg))

	//msg := message.GetNewMarkdownMessge(title, text, mobiles, false).ToJsonString()
	//fmt.Println(msg)
	//fmt.Println(dingtalkclient.Send(HOST, msg))
	//msg := message.GetNewActionCardMessage(title,text, "0", "0", "单独标题",  messageUrl).ToJsonString()
	//fmt.Println(msg)
	//fmt.Println(dingtalkclient.Send(HOST, msg))

	feeds := list.New()
	feeds.PushBack(message.Link{Title:"hello1", Text:"world4", PicURL:"http://www.ruanyifeng.com/blogimg/asset/2014/bg2014051203.png",MessageURL:"http://www.ruanyifeng.com/blogimg/asset/2014/bg2014051203.png"})
	feeds.PushBack(message.Link{Title:"hello2", Text:"world3", PicURL:"http://www.ruanyifeng.com/blogimg/asset/2014/bg2014051203.png",MessageURL:"http://www.ruanyifeng.com/blogimg/asset/2014/bg2014051203.png"})
	feeds.PushBack(message.Link{Title:"hello3", Text:"world2", PicURL:"http://www.ruanyifeng.com/blogimg/asset/2014/bg2014051203.png",MessageURL:"http://www.ruanyifeng.com/blogimg/asset/2014/bg2014051203.png"})
	feeds.PushBack(message.Link{Title:"hello4", Text:"world1", PicURL:"http://www.ruanyifeng.com/blogimg/asset/2014/bg2014051203.png",MessageURL:"http://www.ruanyifeng.com/blogimg/asset/2014/bg2014051203.png"})

	msg := message.GetNewFeedCardMessage(feeds).ToJsonString()
	fmt.Println(msg)
	fmt.Println(dingtalkclient.Send(HOST, msg))
}
