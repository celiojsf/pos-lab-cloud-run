package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type CEPService struct {
	client *http.Client
}

type ViaCEPResponse struct {
	CEP        string `json:"cep"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Erro       bool   `json:"erro"`
}

func NewCEPService() *CEPService {
	return &CEPService{
		client: &http.Client{},
	}
}

func (s *CEPService) ValidateCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}

func (s *CEPService) GetCityFromCEP(cep string) (string, error) {
	if !s.ValidateCEP(cep) {
		return "", fmt.Errorf("invalid zipcode")
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := s.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var cepData ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&cepData); err != nil {
		return "", err
	}

	if cepData.Erro {
		return "", fmt.Errorf("can not find zipcode")
	}

	return fmt.Sprintf("%s, %s", cepData.Localidade, cepData.UF), nil
}
