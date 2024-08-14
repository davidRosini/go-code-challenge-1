package adapter

import (
	"os"

	"codechallenge.test/service"
	"codechallenge.test/usecase"
)

func InitReader() {
	buyOperationUsecase := usecase.NewBuyOperationUsecase()
	sellOperationUsecase := usecase.NewSellOperationUsecase(0.2)
	calculateTaxService := service.NewCalculateTaxService(buyOperationUsecase, sellOperationUsecase)
	calculateTaxHandler := NewCalculateTaxHandler(os.Stdout, calculateTaxService)
	calculateTaxHandler.Execute()
}
