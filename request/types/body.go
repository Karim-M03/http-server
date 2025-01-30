package types

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Body struct {
    Data []byte
}

// ParseJSON parses the body as JSON into the provided target
func (b Body) ParseJSON(target interface{}) error {
	if err := json.Unmarshal(b.Data, target); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}
	return nil
}

// ParseXML parses the body as XML into the provided target
func (b Body) ParseXML(target interface{}) error {
	if err := xml.Unmarshal(b.Data, target); err != nil {
		return fmt.Errorf("failed to parse XML: %w", err)
	}
	return nil
}
