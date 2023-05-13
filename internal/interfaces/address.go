package interfaces

import (
	"context"

	"github.com/jailtonjunior94/go-decorator/internal/dtos"
)

type AddressService interface {
	FetchAddressByZipcode(ctx context.Context, zipcode string) (*dtos.AddressResponse, error)
}
