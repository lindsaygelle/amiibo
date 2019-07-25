package amiibo

func NewAPI() {}

type api interface{}

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
