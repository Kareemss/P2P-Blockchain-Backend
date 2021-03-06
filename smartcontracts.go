package main

func TransactionSmartContract(Transaction Order) {
	SellerProfile, _ := GetUser(2, Transaction.Seller)
	BuyerProfile, _ := GetUser(2, Transaction.Buyer)
	value := Transaction.Amount * Transaction.Price
	if Transaction.Issuer == Transaction.Seller {
		AddBalance(SellerProfile.Email, "currency-balance", value)
		AddBalance(BuyerProfile.Email, "energy-balance", Transaction.Amount)
		AddBalance(BuyerProfile.Email, "currency-balance", -value)
	} else {
		AddBalance(SellerProfile.Email, "energy-balance", -Transaction.Amount)
		AddBalance(SellerProfile.Email, "currency-balance", value)
		AddBalance(BuyerProfile.Email, "energy-balance", Transaction.Amount)
	}
	UpdateOrder(Transaction.OrderID, Transaction.Amount)
	SellerProfile.CompletedTransaction += 1
	BuyerProfile.CompletedTransaction += 1
	UpdateTransactionCount(SellerProfile.Email,
		SellerProfile.CompletedTransaction)
	UpdateTransactionCount(BuyerProfile.Email,
		BuyerProfile.CompletedTransaction)
}

func UpdateOrder(OrderID int, TAmount float32) {
	Order, _ := GetOrder(OrderID)
	UpdateAmount := Order.Amount - TAmount
	if UpdateAmount <= 0 {
		DeleteDocFromDB("Market", "Orders", "_id", OrderID)

	} else {
		UpdateFromDB("Market", "Orders", "_id", OrderID, "amount", UpdateAmount)
	}
}

func UpdateTransactionCount(Email string, NewCount int) {
	UpdateFromDB("Users", "Users", "email", Email,
		"completed-transactions", NewCount)
}

func DeleteOrder(Order Order) {
	// Order, _ := GetOrder(OrderID.(int))
	Issuer, _ := GetUser(2, Order.Issuer)
	if Order.Issuer == Order.Buyer {
		Balance := Order.Amount * Order.Price
		AddBalance(Issuer.Email, "currency-balance", Balance)
	} else {
		AddBalance(Issuer.Email, "energy-balance", Order.Amount)
	}

}
