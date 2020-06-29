package transactions

const (
	path = "/sessions"
)

// Transactions Transactions
type Transactions interface {
	Checkout(*CheckoutInput) error
	Cancel()
	Refunds()
	ConsultNotification()
	QueryTransactions()
	TransactionDetails()
}
