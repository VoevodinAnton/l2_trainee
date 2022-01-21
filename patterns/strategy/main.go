package main

import "strategy/pkg"

func main() {
	strategy := &pkg.FilterFoFirstAlgorithm{}
	pv := pkg.InitStrategy(strategy)
	pv.UserFilters["role"] = 1
	pkg.GetData(pv)

	strategySecond := &pkg.FilterFoSecondAlgorithm{}
	pv.UserFilters["role"] = 2
	pv.SetStrategy(strategySecond)
	pkg.GetData(pv)
}
