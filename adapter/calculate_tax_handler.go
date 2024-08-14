package adapter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"codechallenge.test/domain"
	"codechallenge.test/service"
)

type CalculateTaxHandler struct {
	out     io.Writer
	service service.ICalculateTaxService
}

func NewCalculateTaxHandler(out io.Writer, cts service.ICalculateTaxService) *CalculateTaxHandler {
	return &CalculateTaxHandler{
		out:     out,
		service: cts,
	}
}

func (cth *CalculateTaxHandler) Execute() {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		var operations []domain.OperationStock
		err = json.Unmarshal([]byte(input), &operations)
		if err != nil {
			continue
		}

		ch := make(chan []domain.TaxPay)
		go func() {
			ch <- cth.service.Execute(operations)
		}()
		taxes := <-ch

		jsonResult, err := json.Marshal(taxes)
		if err != nil {
			return
		}
		fmt.Fprintln(cth.out, string(jsonResult))
	}
}
