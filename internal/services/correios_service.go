package services

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/jailtonjunior94/go-decorator/internal/dtos"
	"github.com/jailtonjunior94/go-decorator/internal/interfaces"
)

type correiosService struct {
	httpClient *http.Client
}

func NewCorreiosService() interfaces.AddressService {
	return &correiosService{httpClient: &http.Client{}}
}

func (s *correiosService) FetchAddressByZipcode(ctx context.Context, zipcode string) (*dtos.AddressResponse, error) {
	payload := `
			<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cli="http://cliente.bean.master.sigep.bsb.correios.com.br/">
				<soapenv:Header/>
				<soapenv:Body>
					<cli:consultaCEP>
						<cep>` + zipcode + `s</cep>
					</cli:consultaCEP>
				</soapenv:Body>
			</soapenv:Envelope>
		`
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente?wsdl", bytes.NewBufferString(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/soap+xml;charset=utf-8")
	req.Header.Set("cache-control", "no-cache")

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res != nil {
		defer res.Body.Close()
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("[ERROR] [StatusCode] [%d] [Detail] [%s]", res.StatusCode, string(b))
	}

	var result *dtos.CorreiosResponse
	err = xml.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	response := dtos.NewAddressResponse(
		result.Body.ConsultaCEPResponse.Return.Cep,
		result.Body.ConsultaCEPResponse.Return.End,
		result.Body.ConsultaCEPResponse.Return.Bairro,
		result.Body.ConsultaCEPResponse.Return.Cidade,
		result.Body.ConsultaCEPResponse.Return.Uf,
	)
	return response, nil
}
