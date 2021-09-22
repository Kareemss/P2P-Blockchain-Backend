package main

// kjhjkhkjhkjhkj
func TransactionSmartContract(Transaction Order) {
	SellerProfile, _ := GetUser(2, Transaction.Seller)
	BuyerProfile, _ := GetUser(2, Transaction.Buyer)
	value := Transaction.Amount * Transaction.Price
	AddBalance(SellerProfile.Email, "energy-balance", -Transaction.Amount)
	AddBalance(SellerProfile.Email, "currency-balance", value)
	AddBalance(BuyerProfile.Email, "energy-balance", Transaction.Amount)
	AddBalance(BuyerProfile.Email, "currency-balance", -value)
	UpdateOrder(Transaction.Issuer, Transaction.Amount)
}

func UpdateOrder(Issuer string, TAmount float32) {
	Order, _ := GetOrder(Issuer)
	UpdateAmount := Order.Amount - TAmount
	if UpdateAmount == 0 {
		DeleteDocFromDB("Market", "Orders", "issuer", Issuer)
	} else {
		UpdateFromDB("Market", "Orders", "issuer", Issuer, "amount", UpdateAmount)
	}
}
