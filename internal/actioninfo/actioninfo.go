package actioninfo

import "fmt"

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		if err := dp.Parse(v); err != nil {
			fmt.Println(err)
		}
		res, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}
