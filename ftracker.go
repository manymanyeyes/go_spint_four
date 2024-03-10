package ftracker

import (
	"fmt"
	"math"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага в метрах.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// distance float64 - дистанция пройденная в километрах.
// duration float64 — длительность тренировки в часах.
func meanSpeed(action int, distance, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return distance / duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// trainingType string — вид тренировки(Бег, Ходьба, Плавание).
// duration float64 — длительность тренировки в часах.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	distance := distance(action)
	var speed float64
	var calories float64
	switch {
	case trainingType == "Бег":
		speed = meanSpeed(action, distance, duration)
		calories = RunningSpentCalories(speed, duration, weight)
	case trainingType == "Ходьба":
		speed = meanSpeed(action, distance, duration)
		calories = WalkingSpentCalories(speed, duration, weight, height)
	case trainingType == "Плавание":
		speed = swimmingMeanSpeed(lengthPool, countPool, duration)
		calories = SwimmingSpentCalories(speed, duration, weight)
	default:
		return "неизвестный тип тренировки"
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18   // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 1.79 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// speed float64 — средняя скорость движения при беге в км/ч.
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
func RunningSpentCalories(speed, duration, weight float64) float64 {
	return ((runningCaloriesMeanSpeedMultiplier * speed * runningCaloriesMeanSpeedShift) * weight / mInKm * duration * minInH)
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// speed float64 — средняя скорость движения при ходьбе в км/ч.
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя в килограммах.
// height float64 — рост пользователя в сантиметрах.
func WalkingSpentCalories(speed, duration, weight, height float64) float64 {
	return ((walkingCaloriesWeightMultiplier*weight + (math.Pow(speed/kmhInMsec, 2)/height)*walkingSpeedHeightMultiplier) * duration * minInH)
}

// Константы для расчета калорий, расходуемых при плавании.
const (
	swimmingCaloriesMeanSpeedShift   = 1.1 // среднее количество сжигаемых колорий при плавании относительно скорости.
	swimmingCaloriesWeightMultiplier = 2   // множитель веса при плавании.
)

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool) * float64(countPool) / mInKm / duration
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
//
// Параметры:
//
// speed float64 — средняя скорость движения при плавании в км/ч.
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя в килограммах.
func SwimmingSpentCalories(speed, duration, weight float64) float64 {
	return (speed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}
