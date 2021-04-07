package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"calculator/calculator"
)

var (
	ErrInvalidArguments = errors.New("invalid number of arguments")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: ./calc [expression (2 + 2)]")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		panic(fmt.Errorf("%s: %d", ErrInvalidArguments, flag.NArg()))
	}
	calc := calculator.NewCalculator(flag.Arg(0))
	result, err := calc.Evaluate()
	if err != nil {
		panic(err)
	}
	log.Println(result)
}
