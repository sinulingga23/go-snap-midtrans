package main

import (
	"log"

	"github.com/sinulingga23/go-snap-midtrans/constants"
	"github.com/sinulingga23/go-snap-midtrans/payload"
	"github.com/sinulingga23/go-snap-midtrans/service"
	"github.com/sinulingga23/go-snap-midtrans/utils"
)

func main() {
	transactionId, errGenerateRandomNumberString := utils.GenerateRandomNumberString(constants.LENGTH_TRANSACTION_ID)
	if errGenerateRandomNumberString != nil {
		log.Printf("errGenerateRandomNumberString: %v", errGenerateRandomNumberString)
	}

	snapService := service.NewSnapService()
	response, errAcquireToken := snapService.AcquireToken(payload.AcquireTokenSnapRequest{
		TransactionDetailsSnap: payload.TransactionDetailsSnap{
			OrderId:     transactionId,
			GrossAmount: 10000000,
		},
		CustomerDetailsSnap: payload.CustomerDetailsSnap{
			FirstName: "denny",
			LastName:  "rezky sinulingga",
			Email:     "sinulinggatwo@gmail.com",
			Phone:     "0191919",
		},
	})
	if errAcquireToken != nil {
		log.Printf("errAcquireToken: %v", errAcquireToken)
	}
	log.Printf("%v", response)
}
