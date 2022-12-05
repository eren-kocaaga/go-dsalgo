package bridge

import "fmt"

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
}

func (w *Windows) SetPrinter(printer Printer) {
	w.printer = printer
}
