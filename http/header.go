package http

type Headers map[string]string

func (pointer *Headers) Add(key, value string) bool {
	_, ok := pointer.Get(key)
	ok = (ok != true)
	if ok {
		(*pointer)[key] = value
	}
	return ok
}

func (pointer *Headers) Assign(h *Headers) *Headers {
	h.Each(func(key, value string) {
		pointer.Add(key, value)
	})
	return pointer
}

func (pointer *Headers) Del(key string) bool {
	delete(*pointer, key)
	ok := (pointer.Has(key) == false)
	return ok
}

func (pointer *Headers) Each(f func(key, value string)) {
	for key, value := range *pointer {
		f(key, value)
	}
}

func (pointer *Headers) Get(key string) (string, bool) {
	header, ok := pointer.Get(key)
	return header, ok
}

func (pointer *Headers) Has(key string) bool {
	_, ok := pointer.Get(key)
	return ok
}

func (pointer *Headers) Intersection(h *Headers) bool {
	for key := range *h {
		if pointer.Has(key) {
			return true
		}
	}
	return false
}

func (pointer *Headers) Len() int {
	return len(*pointer)
}

func (pointer *Headers) Set(key, value string) *Headers {
	(*pointer)[key] = value
	return pointer
}
