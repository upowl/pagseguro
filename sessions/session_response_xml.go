package sessions

import "encoding/xml"

// SessionResponseXML SessionResponseXML
type SessionResponseXML struct {
	XMLName xml.Name `xml:"session"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id"`
}

// GetSessionID GetSessionID
func (item *SessionResponseXML) GetSessionID() string {
	return item.ID
}
