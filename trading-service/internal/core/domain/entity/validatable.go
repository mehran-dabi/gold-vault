package entity

type Validatable interface {
	Validate() error
}
