package sessions

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
)

// GetToken GetToken
func (item *sessions) GetToken() (*string, error) {
	url := fmt.Sprintf(
		"%s%s?email=%s&token=%s",
		path,
		item.url,
		item.email,
		item.token,
	)

	req, errReq := http.NewRequest("POST", url, nil)
	if errReq != nil {
		return nil, errReq
	}

	clientReq := &http.Client{}

	resp, errRest := clientReq.Do(req)
	if errRest != nil {
		return nil, errRest
	}

	defer resp.Body.Close()

	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		return nil, errBody
	}

	reader := bytes.NewReader(body)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	sessionResponseXML := SessionResponseXML{}
	if err := decoder.Decode(&sessionResponseXML); err != nil {
		return nil, err
	}

	sessionID := sessionResponseXML.ID

	return &sessionID, nil
}
