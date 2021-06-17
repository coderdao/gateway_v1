package main

import (
	"fmt"
	"math/rand"
)

// 接口定义
type LoadBalance interface {
	//选择一个后端Server
	//参数remove是需要排除选择的后端Server
	Next(remove []string) *Server
	//更新可用Server列表
	UpdateServers(servers []*Server)
}

// 后端Server定义
type Server struct {
	//主机地址
	Host string
	//主机名
	Name string
	Id int
	//主机是否在线
	Online bool
}

type LoadBalanceRandom struct{
	servers []*Server
}

// 实例化 随机均衡负载
func NewLoadBalanceRandom(servers []*Server) *LoadBalanceRandom{
	newBalance := &LoadBalanceRandom{}
	newBalance.UpdateServers(servers)
	return newBalance
}

//选择一个后端Server
func (r *LoadBalanceRandom) Next() *Server {
	if len(r.servers) == 0 {
		return nil
	}

	curIndex := rand.Intn(len(r.servers))
	return r.servers[curIndex]
}

func (r *LoadBalanceRandom) Get(key string) (*Server, error) {
	return r.Next(), nil
}

//系统运行过程中，后端可用Server会更新
func (this *LoadBalanceRandom) UpdateServers(servers []*Server) {
	newServers:=make([]*Server,0)
	for _,e:=range servers {
		if e.Online==true {
			newServers=append(newServers,e)
		}
	}
	this.servers=newServers
}

func main() {
	count:=make([]int,4)
	servers:=make([]*Server,0)
	servers=append(servers,&Server{Host:"1",Id:0,Online:true})
	servers=append(servers,&Server{Host:"2",Id:1,Online:true})
	servers=append(servers,&Server{Host:"3",Id:2,Online:true})
	servers=append(servers,&Server{Host:"4",Id:3,Online:true})
	lb:=NewLoadBalanceRandom(servers)

	// 创建4个Server，随机选择100000次。查看4台机器 被选中次数
	for i:=0;i<100000;i++{
		c:=lb.Next()
		count[c.Id]++
	}
	fmt.Println(count)
}