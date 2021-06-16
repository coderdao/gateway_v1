package random_load

import (
	"errors"
	"fmt"
	"gateway/pratise/proxy/load_balance"
	"math/rand"
	"strings"
)

/**
	随机负载均衡
 */
type RandomBalance struct {
	curIndex int
	rss []string

	// 观察主体
	conf load_balance.LoadBalanceConf
}

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at last")
	}

	addr := params[0]
	r.rss = append(r.rss, addr)
	return nil
}

func (r *RandomBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}

	r.curIndex = rand.Intn(len(r.rss))
	return r.rss[r.curIndex]
}

func (r *RandomBalance) Get(key string) (string, error) {
	return r.Next(), nil
}

func (r *RandomBalance) Update() {
	if conf, ok := r.conf.(*load_balance.LoadBalanceZkConf); ok {
		fmt.Println("Update get conf:", conf.GetConf())
		r.rss = []string{}
		for _, ip := range conf.GetConf() {
			r.Add(strings.Split(ip, ",")...)
		}
	}
	if conf, ok := r.conf.(*load_balance.LoadBalanceCheckConf); ok {
		fmt.Println("Update get conf:", conf.GetConf())
		r.rss = nil
		for _, ip := range conf.GetConf() {
			r.Add(strings.Split(ip, ",")...)
		}
	}
}