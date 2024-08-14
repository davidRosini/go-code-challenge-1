# Code Challenge: Capital Gains

This challenge is written in Golang to process user input in `json` format to calculate taxes on financial transactions. The input format is:
```
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":20.00, "quantity": 5000}]
``` 
After processing, the output will be generated in the format:
```
[{"tax":0}, {"tax":10000}]
``` 

## Project Structure

```capital-gains/
  adapter/
    ...
  domain/
    ...
  service/
    ...
  usecase/
    ...
  main.go
  go.mod
```

### Application Divided into 3 Layers:

- `adapter` layer for reading and generating the application output while interacting with the `service` layer
- `service` layer that orchestrates the use cases for execution
- `usecase` layer that focuses on the execution rules of the application
- `domain` layer shared to define the data model used by the application

## Installation

Download and install Golang from the link [download and install](https://go.dev/doc/install) 

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
