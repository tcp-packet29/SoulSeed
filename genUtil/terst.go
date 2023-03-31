package genUtil

import (
	"fmt"
)

func checkErr(er error) {
	if er != nil {
		fmt.Println(er)
	}
}
func main() {
	fmt.Println("MONOKAI")
}
