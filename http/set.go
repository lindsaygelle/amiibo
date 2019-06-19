package http

type Set map[string]int

func (pointer *Set) Add(key string) *Set {
	_, ok := pointer.Get(key)
	ok = ok != true
	if ok {
		(*pointer)[key] = 1
	}
	return pointer
}

func (pointer *Set) Assign(keys ...string) *Set {
	for _, key := range keys {
		pointer.Add(key)
	}
	return pointer
}

func (pointer *Set) Del(key string) bool {
	delete(*pointer, key)
	ok := (pointer.Has(key) == false)
	return ok
}

func (pointer *Set) Each(f func(key string)) {
	for key, _ := range *pointer {
		f(key)
	}
}

func (pointer *Set) Get(key string) (string, bool) {
	_, ok := pointer.Get(key)
	return key, ok
}

func (pointer *Set) Has(key string) bool {
	_, ok := pointer.Get(key)
	return ok
}

func (pointer *Set) Len() int {
	return len(*pointer)
}
