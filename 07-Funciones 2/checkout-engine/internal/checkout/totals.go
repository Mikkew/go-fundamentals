package checkout

func NewOrder(id, customer string) Order {
	return Order{
		ID:       id,
		Customer: customer,
		Items:    []Item{},
		Meta:     map[string]string{},
	}
}

func (o *Order) AddItem(item Item) {
	o.Items = append(o.Items, item)
}

func (o *Order) RemoveItem(sku string) bool {
	for i := range o.Items {
		if o.Items[i].SKU == sku {
			o.Items = append(o.Items[:i], o.Items[i+1:]...)
			return true
		}
	}
	return false
}

func CalcLiteTotal(item Item) Money {
	return item.Price * Money(item.Qty)
}

func (order *Order) CalcSubtotal() Money {
	var sum Money
	for _, item := range order.Items {
		sum += CalcLiteTotal(item)
	}
	return sum
}

func (order *Order) CalcTotalQty() int {
	total := 0
	for _, item := range order.Items {
		total += item.Qty
	}
	return total
}

func (order *Order) AddItems(items ...Item) {
	order.Items = append(order.Items, items...)
}

func (order *Order) FindItem(sku string) (Item, bool) {
	for _, item := range order.Items {
		if item.SKU == sku {
			return item, true
		}
	}

	return Item{}, false
}

func GetMeta(order Order, key string) (string, bool) {
	if order.Meta == nil {
		return "", false
	}

	val, ok := order.Meta[key]
	return val, ok
}

func IndexOfItem(order Order, sku string) (int, bool) {
	for index, item := range order.Items {
		if item.SKU == sku {
			return index, true
		}
	}

	return -1, false
}

func ApplyDiscounts(order Order, fns ...DiscountFn) Money {
	var discount Money
	for _, fn := range fns {
		discount += fn(order)
	}
	sub := order.CalcSubtotal()
	if discount > sub {
		return sub
	}
	return discount
}

func Compute(order Order, bundle Money, tax TaxFn, ship ShippingFn, discounts ...DiscountFn) (t Totals, err error) {
	defer Track("Compute")()
	if err = ValidateOrder(order); err != nil {
		return Totals{}, err
	}

	t.Subtotal = order.CalcSubtotal()

	if bundle > 0 {
		t.Discount = bundle
	} else {
		t.Discount = ApplyDiscounts(order, discounts...)
	}

	t.Tax = tax(order)
	t.Shipping = ship(order)
	t.Total = t.Subtotal - t.Discount + t.Tax + t.Shipping

	return t, nil // return
}
