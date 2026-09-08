package order

import "strconv"

// Блок order — объединяющая задача.
// Здесь нужно закрепить:
// constants/iota, switch, int conversion, bool,
// вычисление цены в копейках и сборку итоговой строки.

// Статусы заказа.
//
// TODO: задайте StatusNew=0, StatusPaid=1 и StatusCanceled=2 как последовательные значения в указанном порядке.
const (
	StatusNew = iota
	StatusPaid
	StatusCanceled
)

// OrderSummary собирает краткое описание заказа.
//
// TODO: верните строку "status=<status_text> payment=<payment_text> price_kop=<price>". Статусы: new, paid, canceled, неизвестный — unknown; paid=true даёт paid, иначе not_paid; отрицательную цену считайте 0, остальные рубли переведите в копейки.
func OrderSummary(status int, priceRub int, paid bool) string {
	//status
	var statusstr string
	switch status {
	case StatusPaid:
		statusstr = "paid"
	case StatusNew:
		statusstr = "new"
	case StatusCanceled:
		statusstr = "canceled"
	default:
		statusstr = "unknown"
	}
	//price
	var priceKorp string
	if priceRub < 0 {
		priceKorp = "0"
	} else {
		priceKorp = strconv.Itoa(priceRub * 100)
	}
	//paid
	var paidstr string
	if paid == true {
		paidstr = "paid"
	} else {
		paidstr = "not_paid"
	}
	return "status=" + statusstr + " payment=" + paidstr + " price_kop=" + priceKorp
}
