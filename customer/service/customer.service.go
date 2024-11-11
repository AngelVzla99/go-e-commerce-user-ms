package service

import (
	"github.com/AngelVzla99/e-commerce-user-ms/customer/dto"
	"github.com/AngelVzla99/e-commerce-user-ms/customer/mapper"
	"github.com/AngelVzla99/e-commerce-user-ms/customer/repository"
	kafka_client "github.com/AngelVzla99/e-commerce-user-ms/integrations/kafka"
)

func NewCustomerService() *CustomerService {
	return &CustomerService{
		repository:  repository.NewCustomerRepository(),
		kafkaClient: kafka_client.NewKafkaClient(),
	}
}

type CustomerService struct {
	repository  *repository.CustomerRepository
	kafkaClient *kafka_client.KafkaClient
}

func (s *CustomerService) CreateCustomers(dto dto.CreateCustomerDto) (string, error) {
	// Create customer in the database
	customerId, error := s.repository.CreateCustomers(mapper.MapCustomerDtoToDbe(dto))
	if error != nil {
		return "", error
	}

	// Send message to Kafka
	s.kafkaClient.Produce(customerId, "customer_created")

	return customerId, nil
}

func (s *CustomerService) GetCustomer(id string) (dto.CustomerDto, error) {
	customerDbe, error := s.repository.GetCustomerById(id)
	if error != nil {
		return dto.CustomerDto{}, error
	}

	return mapper.MapCustomerDbeToDto(customerDbe), nil
}
