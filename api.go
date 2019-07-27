package amiibo

func NewAPI() {}

// API defines the methods for an API struct.
type api interface{}

// API stores the raw Amiibo & Item information that is contained in the Nintendo XHR HTTP response.
// API provides a single point of reference for data transactions that can be performed on the
// internal raw Amiib & raw Item pointers. 
type API struct {
	rawPayload *RawPayload
}

func (pointer *API) Get() error {
	if pointer.rawPayload == nil {
		b, err := net()
		if err != nil {
			return err
		}
		NewRawPayload(*b...)
	}
	return nil
}
