package usecase

import (
	"avyaas/internal/domain/presenter"
	"bytes"
	"io"
	"net/http"

	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

func (uCase *usecase) InitiateKhaltiPayment(request presenter.InitateKhaltiPaymentRequest) (*presenter.InitiateKhaltiPaymentResponse, map[string]string) {
	errMap := make(map[string]string)

	initiateUrl := fmt.Sprintf("%v", viper.GetString("paymentGateways.khalti.initiateUrl"))

	payloadStruct := struct {
		ReturnUrl         string `json:"return_url"`
		WebsiteUrl        string `json:"website_url"`
		Amount            int    `json:"amount"`
		PurchaseOrderID   string `json:"purchase_order_id"`
		PurchaseOrderName string `json:"purchase_order_name"`
	}{
		ReturnUrl:         request.ReturnUrl,
		WebsiteUrl:        request.WebsiteUrl,
		Amount:            request.Amount,
		PurchaseOrderID:   request.PurchaseOrderID,
		PurchaseOrderName: request.PurchaseOrderName,
	}

	payload, err := json.Marshal(payloadStruct)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	req, err := http.NewRequest("POST", initiateUrl, bytes.NewBuffer(payload))
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	req.Header.Set("Authorization", "Key "+viper.GetString("paymentGateways.khalti.secretKey"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	defer resp.Body.Close()

	var responseData *presenter.InitiateKhaltiPaymentResponse

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		errMap["error"] = err.Error()
		return nil, errMap
	}

	return responseData, nil
}
