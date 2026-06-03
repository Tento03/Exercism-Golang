package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var unit= map[string]int{"quarter_of_a_dozen":3,"half_of_a_dozen":6,"dozen":12,"small_gross":120,"gross":144,"great_gross":1728}
    return unit
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	var bill=make(map[string]int)
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	k,v := units[unit]
    if (!v){
        return false
    }

    bill[item]+=k
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    currentQty, itemExists := bill[item]
    if !itemExists {
        return false
    }

    qtyToRemove, unitExists := units[unit]
    if !unitExists {
        return false
    }

    newQty := currentQty - qtyToRemove

    if newQty < 0 {
        return false
    }

    if newQty == 0 {
        delete(bill, item)
        return true
    }

    bill[item] = newQty
    return true
}
// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	k,v := bill[item]
    if !v{
        return 0,false
    }

    return k,true
}
