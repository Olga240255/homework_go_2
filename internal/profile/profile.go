package profile

import (
	"strconv"
	"strings"
)

// Блок profile — объединяющая задача.
// Здесь нужно закрепить:
// string normalization, int, bool, if,
// форматирование строки и простую бизнес-логику.

// BuildUserCard собирает карточку пользователя.
//
// TODO: верните строку "name=<name> age=<age> group=<group> status=<status>". Имя очистите по краям, приведите к нижнему регистру и сделайте первый символ заглавным; пустое имя замените на "Unknown". Возраст от 18 — adult, иначе minor; active=true — active, иначе inactive.
func BuildUserCard(name string, age int, active bool) string {
	//name
	name = strings.Title(strings.ToLower(strings.TrimSpace(name)))
	if name == "" {
		name = "Unknown"
	}
	//age
	var gr string
	if age >= 18 {
		gr = "adult"
	} else {
		gr = "minor"
	}
	//group&status
	var st string
	if active == true {
		st = "active"
	} else {
		st = "inactive"
	}
	return "name=" + name + " age=" + strconv.Itoa(age) + " group=" + gr + " status=" + st
}
