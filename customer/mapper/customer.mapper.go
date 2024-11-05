package mapper

import (
	"github.com/AngelVzla99/e-commerce-user-ms/customer/dto"
	"github.com/AngelVzla99/e-commerce-user-ms/customer/repository"
)

func MapCustomerDtoToDbe(dto dto.CreateCustomerDto) repository.CreateCustomerDbe {
	return repository.CreateCustomerDbe{
		Name:  dto.Name,
		Email: dto.Email,
		Phone: dto.Phone,
	}
}
