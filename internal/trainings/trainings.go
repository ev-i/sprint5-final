package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	datastringSlice := strings.Split(datastring, ",") // строка с данными формата "3456,Ходьба,3h00m"
	if len(datastringSlice) != 3 {
		return errors.New("datastring parse error")
	}
	steps, err := strconv.Atoi(datastringSlice[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps <= 0")
	}
	t.Steps = steps
	t.TrainingType = datastringSlice[1]
	duration, err := time.ParseDuration(datastringSlice[2])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	if duration <= 0.0 {
		return errors.New("duration <= 0.0")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Personal.Height)
	mean := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	cal := 0.0
	switch t.TrainingType {
	case "Бег":
		cal, _ = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Ходьба":
		cal, _ = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	ret := fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, t.TrainingType, t.Duration.Hours(), dist, mean, cal)
	return ret, nil
}
