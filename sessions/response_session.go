package sessions

import "encoding/xml"

// ResponseSessionXML ResponseSessionXML
type ResponseSessionXML struct {
	XMLName xml.Name `xml:"session"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id"`
}
