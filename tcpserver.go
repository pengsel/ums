package main
import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"ums/dao"
	"ums/ziface"
	"ums/znet"
)

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//Ping Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	//先读取客户端的数据
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	//再回写ping...ping...ping
	err := request.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type QueryRouter struct {
	znet.BaseRouter
}

func (this *QueryRouter)Handle(request ziface.IRequest){

	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	msg:=request.GetData()
	userinfo:=&dao.Userinfo{}
	if err:=json.Unmarshal(msg,userinfo);err!=nil {
		fmt.Println("msg 格式错误")
	}
	db, err := sql.Open("mysql", "duanpeng:123456@/ums?charset=utf8")
	if err!=nil {
		fmt.Println("===> can't login in database",err)
	}

	rows,err:=db.Query("SELECT count(*) FROM userinfo where username=? and password=?",userinfo.Username,userinfo.Password)
	if err!=nil {
		fmt.Println("sql error",err)
	}
	var count uint32
	for rows.Next(){
		if err:=rows.Scan(&count);err!=nil{
			fmt.Println("no result",err)
		}
	}
	if count>0 {
		err := request.GetConnection().SendBuffMsg(1, []byte("{\"success\":true,\"msg\":\"found user\"}"))
		if err!=nil {
			fmt.Println("send msg err",err)
		}
	}else {
		err := request.GetConnection().SendBuffMsg(1, []byte("{\"success\":false,\"msg\":\"user NOT FOUND\"}"))
		if err!=nil {
			fmt.Println("send msg err",err)
		}
	}





}

func main() {
	//1 创建一个server句柄
	s := znet.NewServer()

	//2 配置路由
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1,&QueryRouter{})

	//3 开启服务
	s.Serve()
}