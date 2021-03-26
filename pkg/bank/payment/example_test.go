package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExamplePaymentSources(){
	cards := []types.Card{
		{
			Active: true,
			Balance: 10_000_00,
			PAN: "5058 xxxx xxxx 9999",
		},{
			Active: false,
			Balance: 10_00,
			PAN: "5058 xxxx xxxx 8888",
		},{
			Active: true,
			Balance: -10_00,
			PAN: "5058 xxxx xxxx 8989",
		},
	}

	payments := PaymentSources(cards)
	for _, payment := range payments {
		fmt.Println(payment.Number)
	}

}