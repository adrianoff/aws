package functions

import (
	"fmt"
	"os"
)

func ReadTemplate() []byte {
	template, err := os.ReadFile("../templates/index.html")
	if err != nil {
		fmt.Println(err)
	}

	return template
}
