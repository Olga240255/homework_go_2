package integer

import "math"

// Блок integer закрепляет работу с целыми числами:
// int, int64, uint, арифметика, целочисленное деление,
// остаток от деления, сравнения и простые вычисления.
//

// Add складывает два целых числа.
//
// TODO: верните сумму a и b. Поддерживаются положительные, отрицательные и нулевые значения.
func Add(a, b int) int {
	return a + b
}

// Subtract вычитает второе число из первого.
//
// TODO: верните результат вычитания b из a. Поддерживаются положительные, отрицательные и нулевые значения.
func Subtract(a, b int) int {
	return b - a
}

// Multiply умножает два целых числа.
//
// TODO: верните произведение a и b. Поддерживаются положительные, отрицательные и нулевые значения.
func Multiply(a, b int) int {
	return a * b
}

// Divide делит первое число на второе через целочисленное деление.
//
// TODO: верните результат целочисленного деления a на b. Если b равен 0, верните 0.
func Divide(a, b int) int {
	return a / b
}

// Remainder возвращает остаток от деления первого числа на второе.
//
// TODO: верните остаток от целочисленного деления a на b. Если b равен 0, верните 0; для отрицательных чисел сохраните правила Go.
func Remainder(a, b int) int {
	return a % b
}

// IsEven проверяет, является ли число чётным.
//
// TODO: верните true для чётного числа и false для нечётного. Ноль и отрицательные числа также должны обрабатываться корректно.
func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

// LastDigit возвращает последнюю цифру числа.
//
// TODO: верните последнюю десятичную цифру числа как значение от 0 до 9. Для отрицательного числа результат тоже должен быть положительным.
func LastDigit(n int) int {
	return int(math.Abs(float64(n % 10)))
}

// Max возвращает большее из двух чисел.
//
// TODO: верните большее из двух значений. При равенстве допустимо вернуть любое из них, так как значения совпадают.
func Max(a, b int) int {
	switch {
	case a >= b:
		return a
	default:
		return b
	}
}

// Min возвращает меньшее из двух чисел.
//
// TODO: верните меньшее из двух значений. При равенстве допустимо вернуть любое из них, так как значения совпадают.
func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

// Clamp ограничивает значение диапазоном [min, max].
//
// TODO: ограничьте value диапазоном [min, max]: ниже диапазона верните min, выше — max, внутри диапазона — исходное value. Считайте, что min <= max.
func Clamp(value, min, max int) int {
	switch {
	case value < min:
		return min
	case value > max:
		return max
	default:
		return value
	}
}

// SumThree складывает три целых числа.
//
// TODO: верните сумму a, b и c. Поддерживаются положительные, отрицательные и нулевые значения.
func SumThree(a, b, c int) int {
	return a + b + c
}

// Average возвращает среднее арифметическое двух целых чисел.
//
// TODO: верните среднее арифметическое a и b как int. Дробная часть должна отбрасываться по правилам целочисленной арифметики Go.
func Average(a, b int) int {
	return (a + b) / 2
}

// IntToInt64 явно преобразует int в int64.
//
// TODO: верните то же числовое значение с типом int64.
func IntToInt64(n int) int64 {
	return int64(n)
}

// NonNegativeToUint преобразует неотрицательное int-число в uint.
//
// TODO: для n >= 0 верните то же значение с типом uint; для отрицательного n верните 0.
func NonNegativeToUint(n int) uint {
	if n < 0 {
		return 0
	} else {
		return uint(n)
	}
}

// CountPages считает количество страниц для списка элементов.
//
// TODO: верните количество страниц для totalItems элементов по pageSize элементов на странице, округляя вверх. При totalItems <= 0 или pageSize <= 0 верните 0.
func CountPages(totalItems, pageSize int) int {
	if totalItems%pageSize == 0 {
		return totalItems / pageSize
	} else {
		return (totalItems/pageSize + 1)
	}
}
