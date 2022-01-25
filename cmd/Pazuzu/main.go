package main

import (
	"fmt"

	"github.com/M1K8/Pazuzu/pkg/graph"
)

func main() {
	fmt.Println(graph.GetHStocksChart("aapl"))
	fmt.Println(graph.Get15MStocksChart("aapl"))
	fmt.Println(graph.GetDStocksChart("aapl"))

}
