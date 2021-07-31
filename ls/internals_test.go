package ftls

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDotDot(t *testing.T) {
	objects := getDotsInfo()
	data, err := json.Marshal(objects)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
