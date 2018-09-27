package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gosuri/uilive"
)

func main() {
	writer := uilive.New()
	writer.Start()

	for i := 0; i <= 100; i++ {
		var table []string
		switch i % 8 {
		case 0:
			table = []string{"X  ", "   ", "   "}
		case 1:
			table = []string{" X ", "   ", "   "}
		case 2:
			table = []string{"  X", "   ", "   "}
		case 3:
			table = []string{"   ", "  X", "   "}
		case 4:
			table = []string{"   ", "   ", "  X"}
		case 5:
			table = []string{"   ", "   ", " X " }
		case 6:
			table = []string{"   ", "   ", "X  "}
		case 7:
			table = []string{"   ", "X  ", "   "}
		}

		fmt.Fprintln(writer, strings.Join(table, "\n"))
		writer.Flush()
		time.Sleep(time.Millisecond * 60)
	}

	fmt.Fprintln(writer, "Finished.")
	writer.Stop()
}
