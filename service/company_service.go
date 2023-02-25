package service

import (
	"boilerplate/persistence/dtos"
	"boilerplate/repository"
	"context"
)

type ICompanyService interface {
	Create(ctx context.Context, input *dtos.CompanyRequestDTO) error
}

type CompanyService struct {
	companyRepository repository.ICompanyRepository
}

// TODO - change implementation to return pointer to service and make receivers pointer based for lite container

func NewCompanyService(userRepository repository.ICompanyRepository) ICompanyService {
	return CompanyService{
		companyRepository: userRepository,
	}
}

func (u CompanyService) Create(ctx context.Context, input *dtos.CompanyRequestDTO) error {

	return nil
}
