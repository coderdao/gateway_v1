package load_balance

import (
	"gateway/pratise/proxy/load_balance/hash"
	"gateway/pratise/proxy/load_balance/random_load"
	"gateway/pratise/proxy/load_balance/round_robin"
	"gateway/pratise/proxy/load_balance/weight_round_robin"
)

type LbType int

const (
	LbRandom LbType = iota
	LbRoundRobin
	LbWeightRoundRobin
	LbConsistentHash
)

func LoadBanlanceFactory(lbType LbType) LoadBalanc {
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
		lb := &RandomBalance{}
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	case LbConsistentHash:
		lb := NewConsistentHashBanlance(10, nil)
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	case LbRoundRobin:
		lb := &RoundRobinBalance{}
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	case LbWeightRoundRobin:
		lb := &WeightRoundRobinBalance{}
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	default:
		lb := &RandomBalance{}
		lb.SetConf(mConf)
		mConf.Attach(lb)
		lb.Update()
		return lb
	}
}
