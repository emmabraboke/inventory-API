package paymentService

import (
	"bytes"
	"encoding/json"
	"inventory/internals/entity/transactionEntity"
	"io"
	"net/http"
)

type paymentSrv struct {
	apiKey string
	apiUrl string
}

type PaymentService interface {
	CreatePayment(req *transactionEntity.PayStackReq) (*transactionEntity.PayStackRes, error)
}

func NewPaymentSrv(apiKey, apiUrl string) PaymentService {
	return &paymentSrv{apiKey: apiKey, apiUrl: apiUrl}
}

func (t *paymentSrv) CreatePayment(paystackReq *transactionEntity.PayStackReq) (*transactionEntity.PayStackRes, error) {

	bodyBytes, err := json.Marshal(paystackReq)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", t.apiUrl, bytes.NewReader(bodyBytes))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+t.apiKey)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resByte, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var paystackRes transactionEntity.PayStackRes

	err = json.Unmarshal(resByte, &paystackRes)

	if err != nil {
		return nil, err
	}

	return &paystackRes, nil
}

func (t *paymentSrv) Webhook() {

}
