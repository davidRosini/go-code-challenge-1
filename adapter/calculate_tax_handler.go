package adapter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"codechallenge.test/domain"
	"codechallenge.test/service"
)

type CalculateTaxHandler struct {
	in      io.Reader
	out     io.Writer
	service service.ICalculateTaxService
}

func NewCalculateTaxHandler(in io.Reader, out io.Writer, cts service.ICalculateTaxService) *CalculateTaxHandler {
	return &CalculateTaxHandler{
		in:      in,
		out:     out,
		service: cts,
	}
}

func (cth *CalculateTaxHandler) Execute() {
	reader := bufio.NewReader(cth.in)

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

		taxes := cth.service.Execute(operations)

		jsonResult, err := json.Marshal(taxes)
		if err != nil {
			return
		}
		fmt.Fprintln(cth.out, string(jsonResult))
	}
}
