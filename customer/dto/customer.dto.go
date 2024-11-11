package dto

type AddressDto struct {
	Street  string `json:"street" binding:"required"`
	City    string `json:"city" binding:"required"`
	State   string `json:"state" binding:"required"`
	Country string `json:"country" binding:"required"`
	ZipCode string `json:"zipCode" binding:"required"`
}

type CreateCustomerDto struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	// Address AddressDto `json:"address" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type CustomerDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Address AddressDto `json:"address"`
	Phone string `json:"phone"`
}
