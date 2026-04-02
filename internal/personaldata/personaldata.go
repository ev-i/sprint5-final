package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	fmt.Printf(`Имя: %s
Вес: %.2f кг.
Рост: %.2f м.
`, p.Name, p.Weight, p.Height)
}
