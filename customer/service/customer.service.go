package service

import (
	"fmt"

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
	fmt.Println("\n\nCreateCustomers")
	// Create customer in the database
	customerId, error := s.repository.CreateCustomers(mapper.MapCustomerDtoToDbe(dto))
	if error != nil {
		return "", error
	}
	fmt.Println("Created in DB with id ", customerId)

	// Send message to Kafka
	s.kafkaClient.Produce(customerId, "customer_created")

	fmt.Println("Message sended to the queue")

	return customerId, nil
}

func (s *CustomerService) GetCustomer(id string) (dto.CustomerDto, error) {
	fmt.Println("\n\nGetCustomer")
	customerDbe, error := s.repository.GetCustomerById(id)
	if error != nil {
		return dto.CustomerDto{}, error
	}

	return mapper.MapCustomerDbeToDto(customerDbe), nil
}
