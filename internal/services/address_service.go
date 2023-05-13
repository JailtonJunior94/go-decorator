package services

import (
	"context"

	"github.com/jailtonjunior94/go-decorator/internal/dtos"
	"github.com/jailtonjunior94/go-decorator/internal/interfaces"
)

type addressService struct {
	address interfaces.AddressService
}

func NewAddressService(address interfaces.AddressService) interfaces.AddressService {
	return &addressService{address: address}
}

func (s *addressService) FetchAddressByZipcode(ctx context.Context, zipcode string) (*dtos.AddressResponse, error) {
	return s.address.FetchAddressByZipcode(ctx, zipcode)
}
