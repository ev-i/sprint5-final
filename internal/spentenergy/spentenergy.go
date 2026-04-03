package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0.0, errors.New("steps <= 0")
	}
	if weight <= 0 {
		return 0.0, errors.New("weight <= 0")
	}
	if height <= 0 {
		return 0.0, errors.New("height <= 0")
	}
	if duration <= 0 {
		return 0.0, errors.New("duration <= 0")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	wsd := weight * meanSpeed * durationMin
	cal := wsd / minInH
	calWithCoef := cal * walkingCaloriesCoefficient
	return calWithCoef, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0.0, errors.New("steps <= 0")
	}
	if weight <= 0 {
		return 0.0, errors.New("weight <= 0")
	}
	if height <= 0 {
		return 0.0, errors.New("height <= 0")
	}
	if duration <= 0 {
		return 0.0, errors.New("duration <= 0")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	wsd := weight * meanSpeed * durationMin
	cal := wsd / minInH
	return cal, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0.0
	}
	distance := Distance(steps, height)
	durationH := duration.Hours()
	meanSpeed := distance / durationH
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	stepsM := stepLength * float64(steps)
	stepsKm := stepsM / mInKm
	return stepsKm
}
