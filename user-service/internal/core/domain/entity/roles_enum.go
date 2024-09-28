package entity

type Roles string

const (
	RoleAdmin    Roles = "admin"
	RoleCustomer Roles = "customer"
)

func (r Roles) String() string {
	return string(r)
}

func (r Roles) IsValid() bool {
	switch r {
	case RoleAdmin, RoleCustomer:
		return true
	default:
		return false
	}
}
