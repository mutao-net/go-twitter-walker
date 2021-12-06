package main
import (
    "fmt"
    "os"
)
func main()  {
	fp, err := os.Open("./target.txt")
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    buf := make([]byte, 64)
    for {
        n, err := fp.Read(buf)
        if n == 0 {
            break
        }
        if err != nil {
            panic(err)
        }
		// fmt.Println(string(buf))
		targets := []string{}
		targets = append(targets, string(buf))
		for _, target := range targets {
			fmt.Println(target)
		}
    }
}