package checkout

type TaxFn func(Order) Money

func NoTax(Order) Money {
	return 0
}

func IVA16(order Order) Money {
	sub := order.CalcSubtotal()
	return sub * 16 / 100
}

func NewTaxByState(state string) TaxFn {
	switch state {
	case "CDMX":
		return func(order Order) Money {
			return order.CalcSubtotal() * 16 / 100
		}
	case "NL":
		return func(order Order) Money {
			return order.CalcSubtotal() * 15 / 100
		}
	case "QRO":
		return func(order Order) Money {
			return order.CalcSubtotal() * 20 / 100
		}
	case "JAL":
		return func(order Order) Money {
			return order.CalcSubtotal() * 14 / 100
		}
	case "PUE":
		return func(order Order) Money {
			return order.CalcSubtotal() * 13 / 100
		}
	default:
		return NoTax

	}
}
