package transactions

import "encoding/xml"

// PaymentMethodXML PaymentMethodXML
type PaymentMethodXML struct {
	XMLName xml.Name `xml:"paymentMethod"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type"`
	Code    string   `xml:"code"`
}

// TransactionXML TransactionXML
type TransactionXML struct {
	XMLName          xml.Name         `xml:"transaction"`
	Text             string           `xml:",chardata"`
	Date             string           `xml:"date"`
	LastEventDate    string           `xml:"lastEventDate"`
	Code             string           `xml:"code"`
	Reference        string           `xml:"reference"`
	RecoveryCode     string           `xml:"recoveryCode"`
	Type             string           `xml:"type"`
	Status           string           `xml:"status"`
	PaymentMethod    PaymentMethodXML `xml:"paymentMethod"`
	PaymentLink      string           `xml:"paymentLink"`
	GrossAmount      string           `xml:"grossAmount"`
	DiscountAmount   string           `xml:"discountAmount"`
	FeeAmount        string           `xml:"feeAmount"`
	NetAmount        string           `xml:"netAmount"`
	ExtraAmount      string           `xml:"extraAmount"`
	InstallmentCount string           `xml:"installmentCount"`
	ItemCount        string           `xml:"itemCount"`
	Items            ItemsXML         `xml:"items"`
	Sender           SenderXML        `xml:"sender"`
	Shipping         ShippingXML      `xml:"shipping"`
}
