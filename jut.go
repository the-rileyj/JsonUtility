package jut

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	var obj interface{}
	var spacing uint

	flag.UintVar(&spacing, "s", 2, "Effects pretty print indentation")
	flag.Parse()

	if err := json.NewDecoder(os.Stdin).Decode(&obj); err != nil {
		fmt.Println(errors.Wrap(err, "Invalid JSON"))
		return
	}

	bytes, err := json.MarshalIndent(obj, "", strings.Repeat(" ", int(spacing)))

	if err != nil {
		fmt.Println(errors.Wrap(err, "Could not marshal JSON for output"))
		return
	}

	fmt.Println(string(bytes))
}
