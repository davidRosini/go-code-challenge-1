# Code Challenge: Capital Gains

This challenge is written in Golang to process user input in `json` format to calculate taxes on financial transactions. The input format is:
```
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":20.00, "quantity": 5000}]
``` 
After processing, the output will be generated in the format:
```
[{"tax":0},{"tax":10000}]
``` 

## Project Structure

```
capital-gains/
  adapter/
    calculate_tax_handler.go
    calculate_tax_handler_test.go
  commons/
    helper.go
    helper_test.go
  domain/
    operation_state.go
    operation_stock.go
    tax_pay.go
  service/
    calculate_tax_service.go
    calculate_tax_service_test.go
  usecase/
    buy_operation_usecase.go
    buy_operation_usecase_test.go
    sell_operation_usecase.go
    sell_operation_usecase_test.go
  main.go
  go.mod
```

### Application Divided into 4 Layers:

- `adapter` — handles I/O (reads stdin, writes JSON to stdout) and interacts with the `service` layer
- `service` — orchestrates the use cases for execution
- `usecase` — contains the business rules for buy and sell operations
- `domain` — shared data models used across all layers
- `commons` — utility functions (weighted average, percentage, rounding)

## Installation

Requires Go 1.22.4+. Download and install from [go.dev](https://go.dev/doc/install).

## Execution

Using the terminal at the root of the project, execute the command:
```bash
go run .
```
or if you have an input file for the program, use:
```bash
cat input.txt | go run .
```

## Tests

To run the project's tests, use the terminal:
```bash
go test ./... -cover -v
```
