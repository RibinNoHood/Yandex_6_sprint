package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertMsg(msg string) string {
	r := []rune(msg)
	answer := ""
	if r[0] == '-' || r[0] == '.' {
		answer = morse.ToText(msg)
		//fmt.Println("Это морзе")
	} else {
		answer = morse.ToMorse(msg)
		//fmt.Println("Это текст")
	}
	return answer
}
