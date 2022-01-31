package main

import (
	"fmt"

	"github.com/M1K8/Pazuzu/pkg/graph"
)

func main() {
	fmt.Println(graph.Get15MCryptoChart("c"))
	fmt.Println(graph.Get15MStocksChart("f"))
	//fmt.Println(graph.GetHCryptoChart("cro"))
	//fmt.Println(graph.GetHStocksChart("f"))

	//fmt.Println(graph.GetHCryptoChart("cro"))
	//fmt.Println(graph.GetHStocksChart("f"))

}
