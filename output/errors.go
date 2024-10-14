package output

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(value any) {
	val, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", val)
	}
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвестный тип ошибки")
	}
}

func sum[T int | string](a, b T) T {
	switch d := any(a).(type) {
	case string:
		fmt.Println(d)
	}
	return a + b
}

