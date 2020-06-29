package transactions

// PhoneInput Phone
type PhoneInput struct {
	AreaCode string `json:"areaCode"`
	Number   string `json:"number"`
}

// DocumentInput Document
type DocumentInput struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// SenderInput Sender
type SenderInput struct {
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Phone     PhoneInput      `json:"phone"`
	Documents []DocumentInput `json:"documents"`
	Hash      string          `json:"hash"`
}

// ItemInput Item
type ItemInput struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Amount      string `json:"amount"`
}

// AddressInput Address
type AddressInput struct {
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	District   string `json:"district"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
}

// ShippingInput Shipping
type ShippingInput struct {
	AddressRequired string       `json:"addressRequired"`
	Address         AddressInput `json:"address"`
	Type            string       `json:"type"`
	Cost            string       `json:"cost"`
}

// InstallmentInput Installment
type InstallmentInput struct {
	Quantity                      string `json:"quantity"`
	Value                         string `json:"value"`
	NoInterestInstallmentQuantity string `json:"noInterestInstallmentQuantity"`
}

// HolderInput Holder
type HolderInput struct {
	Name      string          `json:"name"`
	Documents []DocumentInput `json:"documents"`
	BirthDate string          `json:"birthDate"`
	Phone     PhoneInput      `json:"phone"`
}

// CreditCardInput CreditCard
type CreditCardInput struct {
	Token          string           `json:"token"`
	Installment    InstallmentInput `json:"installment"`
	Holder         HolderInput      `json:"holder"`
	BillingAddress AddressInput     `json:"billingAddress"`
}

// CheckoutInput CheckoutInput
type CheckoutInput struct {
	Mode            string          `json:"mode"`
	Method          string          `json:"method"`
	BankName        string          `json:"bank"`
	Sender          SenderInput     `json:"sender"`
	Currency        string          `json:"currency"`
	NotificationURL string          `json:"notificationURL"`
	Items           []ItemInput     `json:"items"`
	ExtraAmount     string          `json:"extraAmount"`
	Reference       string          `json:"reference"`
	Shipping        ShippingInput   `json:"shipping"`
	CreditCard      CreditCardInput `json:"creditCard"`
}

func (item *CheckoutInput) getPaymentXML() paymentXML {
	return paymentXML{}
}
