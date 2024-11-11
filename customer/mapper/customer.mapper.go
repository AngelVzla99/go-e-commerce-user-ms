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

func MapCustomerDbeToDto(dbe repository.CustomerDbe) dto.CustomerDto {
	return dto.CustomerDto{
		ID:    dbe.ID,
		Name:  dbe.Name,
		Email: dbe.Email,
		Phone: dbe.Phone,
	}
}
