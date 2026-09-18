package text

import (
	"slices"
	"strings"
)

// Блок text закрепляет работу со строками:
// string, byte, rune, immutable string, Unicode,
// пакет strings и простые операции над текстом.

// ByteLen возвращает длину строки в байтах.
//
// TODO: верните размер s в байтах. Для Unicode-строк количество байт может быть больше количества символов.
func ByteLen(s string) int {
	return len(s)
}

// RuneLen возвращает количество Unicode-символов в строке.
//
// TODO: верните количество Unicode-символов в s. Кириллица и emoji должны считаться как отдельные символы.
func RuneLen(s string) int {
	runelist := []rune(s)
	return len(runelist)
}

// FirstRune возвращает первый Unicode-символ строки.
//
// TODO: верните первый Unicode-символ s как строку. Для пустой строки верните "".
func FirstRune(s string) string {
	if s == "" {
		return ""
	}
	runelist := []rune(s)
	return string(runelist[0])
}

// LastRune возвращает последний Unicode-символ строки.
//
// TODO: верните последний Unicode-символ s как строку. Для пустой строки верните "".
func LastRune(s string) string {
	runelist := []rune(s)
	if len(runelist) == 0 {
		return ""
	}
	return string(runelist[len(runelist)-1])
}

// Trim убирает пробелы по краям строки.
//
// TODO: верните s без пробельных символов по краям. Внутреннее содержимое строки не изменяйте.
func Trim(s string) string {
	s = strings.TrimSpace(s)
	return s
}

// ToLower переводит строку в нижний регистр.
//
// TODO: верните s в нижнем регистре с корректной обработкой латиницы и кириллицы.
func ToLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// ToUpper переводит строку в верхний регистр.
//
// TODO: верните s в верхнем регистре с корректной обработкой латиницы и кириллицы.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// NormalizeEmail нормализует email.
//
// TODO: удалите пробельные символы по краям email и приведите всю строку к нижнему регистру.
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// ContainsWord проверяет, содержит ли text подстроку word.
//
// TODO: верните true, если word встречается внутри text. Поиск чувствителен к регистру; пустая строка word считается найденной.
func ContainsWord(text, word string) bool {
	return strings.Contains(text, word)
}

// ReplaceFirstRune заменяет первый Unicode-символ строки.
//
// TODO: верните новую строку, заменив первый Unicode-символ s на r. Для пустой строки верните "".
func ReplaceFirstRune(s string, r rune) string {
	runelist := []rune(s)
	if len(runelist) == 0 {
		return ""
	}
	runelist[0] = r
	return string(runelist)
}

// ReverseRunes разворачивает строку по Unicode-символам.
//
// TODO: верните s в обратном порядке по Unicode-символам. Кириллица и emoji не должны повреждаться.
func ReverseRunes(s string) string {
	runelist := []rune(s)
	slices.Reverse(runelist)
	return string(runelist)
}

// Initials возвращает инициалы имени и фамилии.
//
// TODO: очистите имя и фамилию по краям, возьмите первые Unicode-символы непустых частей, переведите их в верхний регистр и соедините без разделителя.
func Initials(firstName, lastName string) string {
	r1 := []rune(strings.ToUpper(strings.TrimSpace(firstName)))
	r2 := []rune(strings.ToUpper(strings.TrimSpace(lastName)))
	var res []rune
	if len(r1) > 0 {
		res = append(res, r1[0])
	}
	if len(r2) > 0 {
		res = append(res, r2[0])
	}
	return string(res)
}

// RepeatWord повторяет слово count раз.
//
// TODO: верните word, повторённое count раз без разделителя. При count <= 0 верните "".
func RepeatWord(word string, count int) string {
	if count <= 0 {
		return ""
	}
	var str1 string
	for i := 0; i < count; i++ {
		str1 = str1 + word
	}
	return str1
}

// JoinWithComma объединяет строки через запятую.
//
// TODO: соедините элементы values через запятую без дополнительных пробелов. Для пустого или nil-среза верните "".
func JoinWithComma(values []string) string {
	if len(values) == 0 {
		return ""
	}
	str1 := values[0]
	for i := 1; i < len(values); i++ {
		str1 = str1 + "," + values[i]
	}
	return str1

}

// IsPalindrome проверяет, является ли строка палиндромом.
//
// TODO: верните true, если s читается одинаково в обоих направлениях после удаления пробелов и приведения к одному регистру. Поддержите Unicode.
func IsPalindrome(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	runebase := []rune(s)
	var rune1 []rune
	for i, val := range runebase {
		if val != ' ' {
			rune1 = append(rune1, runebase[i])
		}
	}
	rune2 := make([]rune, len(rune1))
	for i := range rune1 {
		rune2[len(rune1)-1-i] = rune1[i]
	}

	if string(rune2) == string(rune1) {
		return true
	}
	return false
}
