package dto

type Pagination struct {
	Limit  int `json:"limit" validate:"required,gte=1,lte=100"`
	Offset int `json:"offset" validate:"required,gte=0"`
}

func (p *Pagination) Validate() error {
	return GetValidator().Struct(p)
}
