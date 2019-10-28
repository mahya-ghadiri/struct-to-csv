package csv

import (
	"fmt"
	"github.com/oleiade/reflections"
	"log"
	"reflect"
	"strings"
)

func GetOutputForCsv(dataSlice []interface{}) (Output string) {
	if len(dataSlice) == 0 {
		return
	}
	var OutputLine []string
	var fields []string
	fields, _ = reflections.Fields(dataSlice[0])
	for i := 0; i <= len(fields)-1; i++ {
		OutputLine = append(OutputLine, fields[i])
	}
	Output += strings.Join(OutputLine, ",") + "\n"

	for _, structData := range dataSlice {
		OutputLine = []string{}
		var fields []string
		fields, _ = reflections.Fields(structData)
		for i := 0; i <= len(fields)-1; i++ {
			rowField, err := reflections.GetField(structData, fmt.Sprintf(fields[i]))
			if err != nil {
				log.Fatal(err)
			}
			switch fieldKind, _ := reflections.GetFieldKind(structData, fields[i]); fieldKind {
			case reflect.String:
				rowField = strings.Replace(rowField.(string), ",", "", -1)
			case reflect.Bool:
				if rowField == true {
					rowField = 1
				} else {
					rowField = 0
				}
			case reflect.Struct:
				rowField = ""
			case reflect.Ptr:
				continue
			}
			OutputLine = append(OutputLine, fmt.Sprintf("%v", rowField))
		}
		Output += strings.Join(OutputLine, ",") + "\n"
	}
	return Output
}
