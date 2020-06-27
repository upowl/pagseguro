package transactions

import "encoding/xml"

// Phone Phone
type Phone struct {
	XMLName  xml.Name `xml:"phone"`
	Text     string   `xml:",chardata"`
	AreaCode string   `xml:"areaCode"`
	Number   string   `xml:"number"`
}

// Document Document
type Document struct {
	XMLName xml.Name `xml:"document"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type"`
	Value   string   `xml:"value"`
}

// Documents Documents
type Documents struct {
	XMLName  xml.Name   `xml:"documents"`
	Text     string     `xml:",chardata"`
	Document []Document `xml:"document"`
}

// Sender Sender
type Sender struct {
	XMLName   xml.Name  `xml:"sender"`
	Text      string    `xml:",chardata"`
	Name      string    `xml:"name"`
	Email     string    `xml:"email"`
	Phone     Phone     `xml:"phone"`
	Documents Documents `xml:"documents"`
	Hash      string    `xml:"hash"`
}

// Item Item
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Text        string   `xml:",chardata"`
	ID          string   `xml:"id"`
	Description string   `xml:"description"`
	Quantity    string   `xml:"quantity"`
	Amount      string   `xml:"amount"`
}

// Items Items
type Items struct {
	XMLName xml.Name `xml:"items"`
	Text    string   `xml:",chardata"`
	Item    []Item   `xml:"item"`
}

// Address Address
type Address struct {
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

// Shipping Shipping
type Shipping struct {
	XMLName         xml.Name `xml:"shipping"`
	Text            string   `xml:",chardata"`
	AddressRequired string   `xml:"addressRequired"`
	Address         Address  `xml:"address"`
	Type            string   `xml:"type"`
	Cost            string   `xml:"cost"`
}

// Bank Bank
type Bank struct {
	XMLName xml.Name `xml:"bank"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"name"`
}

// Installment Installment
type Installment struct {
	XMLName                       xml.Name `xml:"installment"`
	Text                          string   `xml:",chardata"`
	Quantity                      string   `xml:"quantity"`
	Value                         string   `xml:"value"`
	NoInterestInstallmentQuantity string   `xml:"noInterestInstallmentQuantity"`
}

// Holder Holder
type Holder struct {
	XMLName   xml.Name  `xml:"holder"`
	Text      string    `xml:",chardata"`
	Name      string    `xml:"name"`
	Documents Documents `xml:"documents"`
	BirthDate string    `xml:"birthDate"`
	Phone     Phone     `xml:"phone"`
}

// CreditCard CreditCard
type CreditCard struct {
	XMLName        xml.Name    `xml:"creditCard"`
	Text           string      `xml:",chardata"`
	Token          string      `xml:"token"`
	Installment    Installment `xml:"installment"`
	Holder         Holder      `xml:"holder"`
	BillingAddress Address     `xml:"billingAddress"`
}

// Payment Payment
type Payment struct {
	XMLName         xml.Name   `xml:"payment"`
	Text            string     `xml:",chardata"`
	Mode            string     `xml:"mode"`
	Method          string     `xml:"method"`
	Bank            Bank       `xml:"bank"`
	Sender          Sender     `xml:"sender"`
	Currency        string     `xml:"currency"`
	NotificationURL string     `xml:"notificationURL"`
	Items           Items      `xml:"items"`
	ExtraAmount     string     `xml:"extraAmount"`
	Reference       string     `xml:"reference"`
	Shipping        Shipping   `xml:"shipping"`
	CreditCard      CreditCard `xml:"creditCard"`
}
