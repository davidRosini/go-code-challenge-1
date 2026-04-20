package main

import (
	"os"

	"codechallenge.test/adapter"
	"codechallenge.test/service"
	"codechallenge.test/usecase"
)

func main() {
	buyOperationUsecase := usecase.NewBuyOperationUsecase()
	sellOperationUsecase := usecase.NewSellOperationUsecase(0.2)
	calculateTaxService := service.NewCalculateTaxService(buyOperationUsecase, sellOperationUsecase)
	calculateTaxHandler := adapter.NewCalculateTaxHandler(os.Stdin, os.Stdout, calculateTaxService)
	calculateTaxHandler.Execute()
}
