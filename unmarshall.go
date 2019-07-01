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

func unmarshallAmiibo(content *[]byte) (*Amiibo, error) {
	r := &Amiibo{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func unmarshallItem(content *[]byte) (*Item, error) {
	r := &Item{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func unmarshallRawAmiibo(content *[]byte) (*RawAmiibo, error) {
	r := &RawAmiibo{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func unmarshallRawItem(content *[]byte) (*RawItem, error) {
	r := &RawItem{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
