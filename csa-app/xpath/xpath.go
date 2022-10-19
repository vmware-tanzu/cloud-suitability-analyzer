package xpath

import (
	"errors"
	"fmt"
	"gopkg.in/xmlpath.v2"
	"os"
)

func xpath(query string, filePath string) (string, error) {
	path := xmlpath.MustCompile(query)
	file, _ := os.Open(filePath)
	root, err := xmlpath.Parse(file)
	if err != nil {
		fmt.Println(err)
	}
	if value, ok := path.String(root); ok {
		return value, nil
	} else {
		return "" ,errors.New("Not found!")
	}

}
