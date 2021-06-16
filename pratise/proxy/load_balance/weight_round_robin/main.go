package main

import (
	"errors"
	"fmt"
	"gateway/pratise/proxy/load_balance"
	"strconv"
)

func main() {
	rb := &WeightRoundRobinBalance{}
	rb.Add("127.0.0.1:2003", "4") //0
	// rb.Add("127.0.0.1:2004", "3") //1
	rb.Add("127.0.0.1:2005", "2") //2

	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
}

/**
代码实现一个加权负载均衡
Weight  初始化时对节点约定的权重
currentWeight  节点临时权重，每轮都会变化
effectiveWeight 节点有效权重，默认与Weight相同
totalWeight  所有节点有效权重之和：sum(effectiveWeight)
 */

/**
代码实现一个加权负载均衡
1、currentWeight=currentWeight+effecitveWeight
2、选中最大的currentWeight节点为选中节点
3、currentWeight=currentWeight-totalWeight(4+3+2=9)

请求次数  请求前currentWelght        选中的节点  请求后currentWelght
1  [serverA=4,serverB=3,serverC=2]  serverA  [serverA=-1,serverB=6,serverC=4]
2  [serverA=-1,serverB=6,serverC=4] serverB  [serverA=3,serverB=0,serverC=6]
3  [serverA=3,serverB=0,serverC=6]  serverc  [serverA=7,serverB=3,serverC=-1]
4  [serverA=7,serverB=3,serverC=-1]  serverA  [serverA=2,serverB=6,serverC=1]
5  [serverA=2,serverB=6,serverC=1]  serverB  [serverA=6,serverB=0,serverC=3]
6  [serverA=6,serverB=0,serverC=3]  serverA  [serverA=1,serverB=3,serverC=5]
7  [serverA=1,serverB=3,serverC=5]  serverc  [serverA=5,serverB=6,serverC=-2]
 */


type WeightRoundRobinBalance struct {
	curIndex int
	rss      []*WeightNode
	rsw      []int
	//观察主体
	conf load_balance.LoadBalanceConf
}

type WeightNode struct {
	addr            string // 服务器地址
	weight          int //权重值
	currentWeight   int //节点当前权重
	effectiveWeight int //有效权重
}

func (r *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("param len need 2")
	}

	// 权重值
	parInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &WeightNode{addr: params[0], weight: int(parInt)}
	node.effectiveWeight = node.weight
	r.rss = append(r.rss, node)
	return nil
}

func (r *WeightRoundRobinBalance) Next() string {
	total := 0
	var best *WeightNode
	for i := 0; i < len(r.rss); i++ {
		w := r.rss[i]
		//step 1 统计所有有效权重之和
		total += w.effectiveWeight

		//step 2 变更节点临时权重为的节点临时权重+节点有效权重
		w.currentWeight += w.effectiveWeight

		//step 3 有效权重默认与权重相同，通讯异常时-1, 通讯成功+1，直到恢复到weight大小
		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}

		//step 4 选择最大临时权重点节点
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}
	if best == nil {
		return ""
	}
	//step 5 变更临时权重为 临时权重-有效权重之和
	best.currentWeight -= total
	return best.addr
}

func (r *WeightRoundRobinBalance) Get(key string) (string, error) {
	return r.Next(), nil
}

func (r *WeightRoundRobinBalance) SetConf(conf load_balance.LoadBalanceConf) {
	r.conf = conf
}

//func (r *WeightRoundRobinBalance) Update() {
//	if conf, ok := r.conf.(*load_balance.LoadBalanceZkConf); ok {
//		fmt.Println("WeightRoundRobinBalance get conf:", conf.GetConf())
//		r.rss = nil
//		for _, ip := range conf.GetConf() {
//			r.Add(strings.Split(ip, ",")...)
//		}
//	}
//	if conf, ok := r.conf.(*LoadBalanceCheckConf); ok {
//		fmt.Println("WeightRoundRobinBalance get conf:", conf.GetConf())
//		r.rss = nil
//		for _, ip := range conf.GetConf() {
//			r.Add(strings.Split(ip, ",")...)
//		}
//	}
//}
