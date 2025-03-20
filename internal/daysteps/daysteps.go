package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/featsci/go1fl-4-sprint-final/internal/spentcalories"
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

	countSteps, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, time.Duration(0), err
	}

	timeDuration, err := time.ParseDuration(s[1])
	if err != nil {
		return 0, time.Duration(0), err
	}

	return countSteps, timeDuration, nil

}

// DayActionInfo обрабатывает входящий пакет, который передаётся в виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает функция. Если пакет невалидный, storage возвращается без изменений.

func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	countSteps, timeDuration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if countSteps <= 0 {
		return ""
	}
	calories := spentcalories.WalkingSpentCalories(countSteps, weight, height, timeDuration)

	distanceSteps := (float64(countSteps) / StepLength) / float64(1000)

	s := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", countSteps, distanceSteps, calories)

	return s
}
