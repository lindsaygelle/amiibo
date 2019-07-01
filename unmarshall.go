package amiibo

import "encoding/json"

func unmarshall(content *[]byte) (*rawPayload, error) {
	r := &rawPayload{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
