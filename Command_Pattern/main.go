package main

import "github.com/user/sales"

type Salesperson struct {
	CurrentCommand sales.Command
}

func (s *Salesperson) SellProduct() {
	details := sales.SaleDetails{
		// populated
	}

	command := &sales.SaleCommand{
		Device:  &sales.SamsungTV{},
		Details: details,
	}

	s.CurrentCommand = command
	s.CurrentCommand.Execute(nil)
}

func (s *Salesperson) ProcessReturn() {
	// similar to sale
}
