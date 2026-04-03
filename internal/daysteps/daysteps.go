package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	datastringSlice := strings.Split(datastring, ",") // строка с данными формата "678,0h50m"
	if len(datastringSlice) != 2 {
		return errors.New("datastring parse error")
	}
	steps, err := strconv.Atoi(datastringSlice[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps <= 0")
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(datastringSlice[1])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	if duration <= 0.0 {
		return errors.New("duration <= 0.0")
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	ret := fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`, ds.Steps, dist, cal)
	return ret, nil
}
