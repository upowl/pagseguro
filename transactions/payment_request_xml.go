package transactions

import "encoding/xml"

// PhoneXML Phone
type PhoneXML struct {
	XMLName  xml.Name `xml:"phone"`
	Text     string   `xml:",chardata"`
	AreaCode string   `xml:"areaCode"`
	Number   string   `xml:"number"`
}

// DocumentXML Document
type DocumentXML struct {
	XMLName xml.Name `xml:"document"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type"`
	Value   string   `xml:"value"`
}

// DocumentsXML Documents
type DocumentsXML struct {
	XMLName  xml.Name      `xml:"documents"`
	Text     string        `xml:",chardata"`
	Document []DocumentXML `xml:"document"`
}

// SenderXML Sender
type SenderXML struct {
	XMLName   xml.Name     `xml:"sender"`
	Text      string       `xml:",chardata"`
	Name      string       `xml:"name"`
	Email     string       `xml:"email"`
	Phone     PhoneXML     `xml:"phone"`
	Documents DocumentsXML `xml:"documents"`
	Hash      string       `xml:"hash"`
}

// ItemXML Item
type ItemXML struct {
	XMLName     xml.Name `xml:"item"`
	Text        string   `xml:",chardata"`
	ID          string   `xml:"id"`
	Description string   `xml:"description"`
	Quantity    string   `xml:"quantity"`
	Amount      string   `xml:"amount"`
}

// ItemsXML Items
type ItemsXML struct {
	XMLName xml.Name  `xml:"items"`
	Text    string    `xml:",chardata"`
	Item    []ItemXML `xml:"item"`
}

// AddressXML Address
type AddressXML struct {
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

// ShippingXML Shipping
type ShippingXML struct {
	XMLName         xml.Name   `xml:"shipping"`
	Text            string     `xml:",chardata"`
	AddressRequired string     `xml:"addressRequired"`
	Address         AddressXML `xml:"address"`
	Type            string     `xml:"type"`
	Cost            string     `xml:"cost"`
}

// BankXML Bank
type BankXML struct {
	XMLName xml.Name `xml:"bank"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"name"`
}

// InstallmentXML Installment
type InstallmentXML struct {
	XMLName                       xml.Name `xml:"installment"`
	Text                          string   `xml:",chardata"`
	Quantity                      string   `xml:"quantity"`
	Value                         string   `xml:"value"`
	NoInterestInstallmentQuantity string   `xml:"noInterestInstallmentQuantity"`
}

// HolderXML Holder
type HolderXML struct {
	XMLName   xml.Name     `xml:"holder"`
	Text      string       `xml:",chardata"`
	Name      string       `xml:"name"`
	Documents DocumentsXML `xml:"documents"`
	BirthDate string       `xml:"birthDate"`
	Phone     PhoneXML     `xml:"phone"`
}

// CreditCardXML CreditCard
type CreditCardXML struct {
	XMLName        xml.Name       `xml:"creditCard"`
	Text           string         `xml:",chardata"`
	Token          string         `xml:"token"`
	Installment    InstallmentXML `xml:"installment"`
	Holder         HolderXML      `xml:"holder"`
	BillingAddress AddressXML     `xml:"billingAddress"`
}

// PaymentXML Payment
type PaymentXML struct {
	XMLName         xml.Name      `xml:"payment"`
	Text            string        `xml:",chardata"`
	Mode            string        `xml:"mode"`
	Method          string        `xml:"method"`
	Bank            BankXML       `xml:"bank"`
	Sender          SenderXML     `xml:"sender"`
	Currency        string        `xml:"currency"`
	NotificationURL string        `xml:"notificationURL"`
	Items           ItemsXML      `xml:"items"`
	ExtraAmount     string        `xml:"extraAmount"`
	Reference       string        `xml:"reference"`
	Shipping        ShippingXML   `xml:"shipping"`
	CreditCard      CreditCardXML `xml:"creditCard"`
}
