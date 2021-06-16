package load_balance

import (
	"gateway/pratise/proxy/load_balance/hash"
	"gateway/pratise/proxy/load_balance/random_load"
	"gateway/pratise/proxy/load_balance/round_robin"
	"gateway/pratise/proxy/load_balance/weight_round_robin"
)

type LbType int

type LoadBalance interface {
	Add(...string) error
	Get(string) (string, error)

	// 后期服务发现补充
	Update()
}

const (
	LbRandom LbType = iota
	LbRoundRobin
	LbWeightRoundRobin
	LbConsistentHash
)

func LoadBanlanceFactory(lbType LbType) LoadBalance {
	switch lbType {
	case LbRandom:
		return &random_load.RandomBalance{}
	case LbConsistentHash:
		return hash.NewConsistentHashBanlance(10, nil)
	case LbRoundRobin:
		return &round_robin.RoundRobinBalance{}
	case LbWeightRoundRobin:
		return &weight_round_robin.WeightRoundRobinBalance{}
	default:
		return &random_load.RandomBalance{}
	}
}

func LoadBanlanceFactorWithConf(lbType LbType, mConf LoadBalanceConf) LoadBalance {
	//观察者模式
	switch lbType {
	case LbRandom:
		lb := &random_load.RandomBalance{}
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	case LbConsistentHash:
		lb := hash.NewConsistentHashBanlance(10, nil)
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	case LbRoundRobin:
		lb := &round_robin.RoundRobinBalance{}
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	case LbWeightRoundRobin:
		lb := &weight_round_robin.WeightRoundRobinBalance{}
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	default:
		lb := &random_load.RandomBalance{}
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	}
}
