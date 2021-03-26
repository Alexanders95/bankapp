package main

import (
	"fmt"
	"bank/pkg/bank/types"
	_ "bank/pkg/bank/payment"

)

func main () {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 10,
		},{
			ID: 2,
			Amount: 11,
		},{
			ID: 3,
			Amount: 50,
		},{
			ID: 4,
			Amount: 99,
		},{
			ID: 5,
			Amount: 55,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)
	//Output: 3

}

// Max ...
func Max( payments []types.Payment) (types.Payment) {
	
	max := payments[0]
	for _, payment := range payments {
		if max.Amount < payment.Amount {
			max = payment
		}
	}
	return max
}
