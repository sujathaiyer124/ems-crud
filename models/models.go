package models

type Employee struct {
	ID        int     `json:"id,omitempty"`
	FirstName string  `json:"firstname" validate:"required,max=10,min=3"`
	LastName  string  `json:"lastname" validate:"required,max=10,min=3"`
	Email     string  `json:"email" validate:"required,email"`
	Password  string  `json:"password" validate:"required"`
	PhoneNo   string  `json:"phoneno" validate:"required,len=10"`
	Role      string  `json:"role" validate:"required"`
	Salary    float64 `json:"salary" validate:"required,gt=0"`
}


