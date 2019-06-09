
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"time"
	"ums/dao"
	"ums/ziface"
	"ums/znet"
)
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		username:=r.FormValue("username")
		password:=r.FormValue("password")

		fmt.Println("username:",username)
		fmt.Println("password:", password)

		user:=dao.Userinfo{
			Username:username,
			Password:password,
		}
		data,err:=json.Marshal(user)
		if err!=nil {
			fmt.Println("parse json err",err)
			return
		}

		res,err:=rpc(1,data)
		if err!=nil {
			fmt.Println("rpc failed err:",err)
		}
		w.Write(res.GetData())

	}
}

func rpc(msgid uint32,data []byte) (ziface.IMessage,error){

	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn,err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		return nil,errors.New("client start err, EXIT")
	}

	//发封包message消息
	dp := znet.NewDataPack()
	msg, _ := dp.Pack(znet.NewMsgPackage(msgid,[]byte(data)))
	_, err = conn.Write(msg)
	if err !=nil {
		return nil,errors.New("write error err")
	}

	//先读出流中的head部分
	headData := make([]byte, dp.GetHeadLen())
	_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
	if err != nil {
		return nil,errors.New("read head error")
	}
	//将headData字节流 拆包到msg中
	msgHead, err := dp.Unpack(headData)
	if err != nil {
		return nil,errors.New("server unpack err")
	}

	if msgHead.GetDataLen() > 0 {
		//msg 是有data数据的，需要再次读取data数据
		msg := msgHead.(*znet.Message)
		msg.Data = make([]byte, msg.GetDataLen())

		//根据dataLen从io中读取字节流
		_, err := io.ReadFull(conn, msg.Data)
		if err != nil {
			return nil,errors.New("server unpack data err")
		}

		fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
		return msg,nil
	}
	return nil,errors.New("empty bag from tcp server")

}
/*
	HttpServer
*/
func main() {

	fmt.Println("HttpServer ... start")
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}