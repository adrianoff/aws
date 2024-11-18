package functions

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ConvertToImage(template, image_filename string) {
	var buffer1, buffer2 bytes.Buffer

	buffer1.WriteString(template)

	buffer2.WriteString("../server/images/")
	buffer2.WriteString(image_filename)
	buffer2.WriteString(".bmp")

	cmd, err := exec.Command("python3", "-V").Output()
	fmt.Println(string(cmd), err)
	cmd, err = exec.Command("python3", "../converter/convert.py", buffer1.String(), buffer2.String()).Output()
	fmt.Println(string(cmd), err)
}
