package clients

import (
	"encoding/json"
	"fmt"
	"g37-lanchonete/internal/domain/ports"
	"g37-lanchonete/internal/domain/services/dto"
)

type mercadoPagoBroker struct {
	httpClient HttpClient
	brokerPath string
}

func NewMercadoPagoBroker(httpClient HttpClient, brokerPath string) ports.PaymentBroker {
	return mercadoPagoBroker{
		httpClient: httpClient,
		brokerPath: brokerPath,
	}
}

func (b mercadoPagoBroker) GeneratePaymentQRCode(request dto.PaymentQRCodeRequest) (dto.PaymentQRCodeResponse, error) {
	reqBody, err := json.Marshal(&request)
	if err != nil {
		return dto.PaymentQRCodeResponse{}, fmt.Errorf("failed to marshal payment qrcode request, error: %v", err)
	}

	response, err := b.httpClient.DoPost(b.brokerPath, reqBody)
	if err != nil {
		return dto.PaymentQRCodeResponse{}, fmt.Errorf("failed to call mercado pago broker, error: %v", err)
	}
	defer response.Body.Close()

	var paymentQRCodeResponse dto.PaymentQRCodeResponse
	err = json.NewDecoder(response.Body).Decode(&paymentQRCodeResponse)
	if err != nil {
		return dto.PaymentQRCodeResponse{}, fmt.Errorf("failed to decode mercado pago response, error: %v", err)
	}

	return paymentQRCodeResponse, nil
}
