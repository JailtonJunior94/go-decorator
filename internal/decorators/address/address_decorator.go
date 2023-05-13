package decorator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jailtonjunior94/go-decorator/internal/dtos"
	"github.com/jailtonjunior94/go-decorator/internal/interfaces"
	"github.com/jailtonjunior94/go-decorator/pkg/caching"

	"github.com/go-redis/redis/v8"
)

type addressDecorator struct {
	cache caching.ICache
	inner interfaces.AddressService
}

func NewAddressDecorator(cache caching.ICache, inner interfaces.AddressService) interfaces.AddressService {
	return &addressDecorator{
		cache: cache,
		inner: inner,
	}
}

func (d *addressDecorator) FetchAddressByZipcode(ctx context.Context, zipcode string) (*dtos.AddressResponse, error) {
	value, err := d.cache.Get(ctx, zipcode)
	if errors.Is(err, redis.Nil) {
		address, err := d.inner.FetchAddressByZipcode(ctx, zipcode)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		b, _ := json.Marshal(&address)
		if err := d.cache.Set(ctx, zipcode, string(b), 30*time.Second); err != nil {
			log.Println(err)
			return nil, err
		}

		fmt.Println("address found through the postal service")
		return address, nil
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	address := &dtos.AddressResponse{}
	json.Unmarshal([]byte(*value), address)

	fmt.Println("address found through cache")
	return address, nil
}
