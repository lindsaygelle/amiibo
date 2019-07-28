package amiibo

var (
	_ api = (*API)(nil)
)

// NewAPI returns a new Amiibo API pointer.
func NewAPI() *API {
	return &API{}
}

// API defines the methods for an API struct.
type api interface {
	Get() error
}

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
		r, err := unmarshallRawPayload(b)
		if err != nil {
			return err
		}
		pointer.rawPayload = r
	}
	return nil
}
