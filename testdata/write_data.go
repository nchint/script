package main
import (
	"fmt"
	"os"
)

// usage: write_data.go <data-to-stdout> <data-to-sdterr>
func main() {
    fmt.Fprintln(os.Stdout, os.Args[1])
    fmt.Fprintln(os.Stderr, os.Args[2])
}
