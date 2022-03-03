package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}

	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, item_exists := bill[item]
	unit_amount, unit_exists := units[unit]

	if !unit_exists {
		return false
	}

	if item_exists {
		bill[item] += unit_amount
	} else {
		bill[item] = unit_amount
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, item_exists := bill[item]
	_, unit_exists := units[unit]

	if !item_exists || !unit_exists {
		return false
	}

	bill[item] -= units[unit]

	if bill[item] < 0 {
		bill[item] += units[unit]
		return false
	}

	if bill[item] == 0 {
		delete(bill, item)
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	thing, exists := bill[item]

	return thing, exists
}
