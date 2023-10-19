package models

type Employee struct {
	ID        int     `json:"id,omitempty"`
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	PhoneNo   string  `json:"phoneno"`
	Role      string  `json:"role"`
	Salary    float64 `json:"salary"`
	//DateOfBirth time.Time
}
