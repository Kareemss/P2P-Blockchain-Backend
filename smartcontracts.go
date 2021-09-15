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
}
