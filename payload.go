package amiibo

func NewPayload(r *RawPayload) *Payload {
	return &Payload{
		Amiibo: NewAmiiboSliceFromRawSlice(r.AmiiboList),
		Items:  NewAmiiboItemSliceFromRawSlice(r.Items)}
}

type Payload struct {
	Amiibo *AmiiboSlice
	Items  *AmiiboItemSlice
}
