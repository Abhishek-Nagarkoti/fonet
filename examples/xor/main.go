package main

import (
	"fmt"
	"log"

	"github.com/Fontinalis/fonet"
)

var samples = [][][]float64{
	[][]float64{
		[]float64{
			0,
			0,
		},
		[]float64{
			0,
		},
	},
	[][]float64{
		[]float64{
			0,
			1,
		},
		[]float64{
			1,
		},
	},
	[][]float64{
		[]float64{
			1,
			0,
		},
		[]float64{
			1,
		},
	},
	[][]float64{
		[]float64{
			1,
			1,
		},
		[]float64{
			0,
		},
	},
}

func main() {
	n, err := fonet.NewNetwork([]int{2, 3, 1})
	if err != nil {
		log.Fatal(err)
	}
	n.Train(samples, 100000, 1.001)
	fmt.Println(n.Predict([]float64{0, 0}))
	fmt.Println(n.Predict([]float64{0, 1}))
	fmt.Println(n.Predict([]float64{1, 0}))
	fmt.Println(n.Predict([]float64{1, 1}))
}
