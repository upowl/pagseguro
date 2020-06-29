package transactions

import "encoding/xml"

// phoneXML Phone
type phoneXML struct {
	XMLName  xml.Name `xml:"phone"`
	Text     string   `xml:",chardata"`
	AreaCode string   `xml:"areaCode"`
	Number   string   `xml:"number"`
}

// documentXML Document
type documentXML struct {
	XMLName xml.Name `xml:"document"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type"`
	Value   string   `xml:"value"`
}

// documentsXML Documents
type documentsXML struct {
	XMLName  xml.Name      `xml:"documents"`
	Text     string        `xml:",chardata"`
	Document []documentXML `xml:"document"`
}

// senderXML Sender
type senderXML struct {
	XMLName   xml.Name     `xml:"sender"`
	Text      string       `xml:",chardata"`
	Name      string       `xml:"name"`
	Email     string       `xml:"email"`
	Phone     phoneXML     `xml:"phone"`
	Documents documentsXML `xml:"documents"`
	Hash      string       `xml:"hash"`
}

// itemXML Item
type itemXML struct {
	XMLName     xml.Name `xml:"item"`
	Text        string   `xml:",chardata"`
	ID          string   `xml:"id"`
	Description string   `xml:"description"`
	Quantity    string   `xml:"quantity"`
	Amount      string   `xml:"amount"`
}

// itemsXML Items
type itemsXML struct {
	XMLName xml.Name  `xml:"items"`
	Text    string    `xml:",chardata"`
	Item    []itemXML `xml:"item"`
}

// addressXML Address
type addressXML struct {
	XMLName    xml.Name `xml:"address"`
	Text       string   `xml:",chardata"`
	Street     string   `xml:"street"`
	Number     string   `xml:"number"`
	Complement string   `xml:"complement"`
	District   string   `xml:"district"`
	City       string   `xml:"city"`
	State      string   `xml:"state"`
	Country    string   `xml:"country"`
	PostalCode string   `xml:"postalCode"`
}

// shippingXML Shipping
type shippingXML struct {
	XMLName         xml.Name   `xml:"shipping"`
	Text            string     `xml:",chardata"`
	AddressRequired string     `xml:"addressRequired"`
	Address         addressXML `xml:"address"`
	Type            string     `xml:"type"`
	Cost            string     `xml:"cost"`
}

// bankXML Bank
type bankXML struct {
	XMLName xml.Name `xml:"bank"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"name"`
}

// installmentXML Installment
type installmentXML struct {
	XMLName                       xml.Name `xml:"installment"`
	Text                          string   `xml:",chardata"`
	Quantity                      string   `xml:"quantity"`
	Value                         string   `xml:"value"`
	NoInterestInstallmentQuantity string   `xml:"noInterestInstallmentQuantity"`
}

// holderXML Holder
type holderXML struct {
	XMLName   xml.Name     `xml:"holder"`
	Text      string       `xml:",chardata"`
	Name      string       `xml:"name"`
	Documents documentsXML `xml:"documents"`
	BirthDate string       `xml:"birthDate"`
	Phone     phoneXML     `xml:"phone"`
}

// creditCardXML CreditCard
type creditCardXML struct {
	XMLName        xml.Name       `xml:"creditCard"`
	Text           string         `xml:",chardata"`
	Token          string         `xml:"token"`
	Installment    installmentXML `xml:"installment"`
	Holder         holderXML      `xml:"holder"`
	BillingAddress addressXML     `xml:"billingAddress"`
}

// paymentXML Payment
type paymentXML struct {
	XMLName         xml.Name      `xml:"payment"`
	Text            string        `xml:",chardata"`
	Mode            string        `xml:"mode"`
	Method          string        `xml:"method"`
	Bank            bankXML       `xml:"bank"`
	Sender          senderXML     `xml:"sender"`
	Currency        string        `xml:"currency"`
	NotificationURL string        `xml:"notificationURL"`
	Items           itemsXML      `xml:"items"`
	ExtraAmount     string        `xml:"extraAmount"`
	Reference       string        `xml:"reference"`
	Shipping        shippingXML   `xml:"shipping"`
	CreditCard      creditCardXML `xml:"creditCard"`
}
