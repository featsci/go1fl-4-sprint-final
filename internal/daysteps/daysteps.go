package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	var s []string = strings.Split(data, ",")

	if len(s) != 2 {
		return 0, time.Duration(0), errors.New("data != 2")
	}

	counStep, err := strconv.Atoi(s[0])
	if err != nil {
		fmt.Println(s[0])
		return 0, time.Duration(0), err
	}

	tDuration, err := time.ParseDuration(s[1])
	if err != nil {
		fmt.Println(s[1])
		return 0, time.Duration(0), err
	}

	return counStep, tDuration, nil

}

// DayActionInfo обрабатывает входящий пакет, который передаётся в виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает функция. Если пакет невалидный, storage возвращается без изменений.

func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	// fmt.Println(parsePackage(data))
	return fmt.Sprint(parsePackage(data))
}
