package checkout

import (
	"errors"
	"fmt"
	"strings"
)

func TryChangeCustomerByValue(o Order, name string) {
	o.Customer = name
}

func ChangeCustomerByPointer(o *Order, name string) {
	o.Customer = name
}

func setCity(o *Order, city string) {
	if o.Meta == nil {
		o.Meta = map[string]string{}
	}
	o.Meta["city"] = city
}

func setZone(o *Order, zone string) {
	if o.Meta == nil {
		o.Meta = map[string]string{}
	}
	o.Meta["zone"] = zone
}

func ValidateOrder(o Order) error {
	if o.ID == "" {
		return errors.New("El ID de la orden es obligatorio")
	}

	if o.Customer == "" {
		return errors.New("El cliente (customer) de la orden es obligatorio")
	}

	if len(o.Items) == 0 {
		return errors.New("La orden debe tener al menos un 1 elemento")
	}

	for index, item := range o.Items {
		if item.SKU == "" {
			return fmt.Errorf("Elemento [%d]: El SKU es obligatorio", index)
		}

		if item.Qty <= 0 {
			return fmt.Errorf("Elemento [%s]: Su cantidad debe ser mayor a 0", item.SKU)
		}

		if item.Price < 0 {
			return fmt.Errorf("Elemento [%s]: El precio debe ser mayor a cero", item.SKU)
		}
	}
	return nil
}

func ParseCoupon(code string) (Coupon, error) {
	coupon := strings.TrimSpace(strings.ToUpper(code))

	if coupon == "" {
		return Coupon{}, errors.New("Cupón vacío")
	}

	switch coupon {
	case "SAVE10":
		return Coupon{
			Code: coupon,
			Kind: "PERCENT",
			Val:  10,
		}, nil
	case "LESS500":
		return Coupon{
			Code: coupon,
			Kind: "FLAT",
			Val:  500,
		}, nil
	case "FREESHIP":
		return Coupon{
			Code: coupon,
			Kind: "FREESHIP",
			Val:  0,
		}, nil
	default:
		return Coupon{}, fmt.Errorf("Cupón %q: Es inválido", code)
	}

}
