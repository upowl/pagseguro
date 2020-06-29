package transactions

import "encoding/xml"

// paymentMethodXML paymentMethodXML
type paymentMethodXML struct {
	XMLName xml.Name `xml:"paymentMethod"`
	Text    string   `xml:",chardata"`
	Type    string   `xml:"type"`
	Code    string   `xml:"code"`
}

// transactionXML transactionXML
type transactionXML struct {
	XMLName          xml.Name         `xml:"transaction"`
	Text             string           `xml:",chardata"`
	Date             string           `xml:"date"`
	LastEventDate    string           `xml:"lastEventDate"`
	Code             string           `xml:"code"`
	Reference        string           `xml:"reference"`
	RecoveryCode     string           `xml:"recoveryCode"`
	Type             string           `xml:"type"`
	Status           string           `xml:"status"`
	PaymentMethod    paymentMethodXML `xml:"paymentMethod"`
	PaymentLink      string           `xml:"paymentLink"`
	GrossAmount      string           `xml:"grossAmount"`
	DiscountAmount   string           `xml:"discountAmount"`
	FeeAmount        string           `xml:"feeAmount"`
	NetAmount        string           `xml:"netAmount"`
	ExtraAmount      string           `xml:"extraAmount"`
	InstallmentCount string           `xml:"installmentCount"`
	ItemCount        string           `xml:"itemCount"`
	Items            itemsXML         `xml:"items"`
	Sender           senderXML        `xml:"sender"`
	Shipping         shippingXML      `xml:"shipping"`
}
