package usecase

import (
	"avyaas/internal/domain/presenter"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func verifyKhaltiPayment(pIDx string) (success bool, res *presenter.VerifyPaymentStatus, err error) {
	verificationUrl := fmt.Sprintf("%v", viper.GetString("paymentGateways.khalti.verifyUrl"))

	payloadStruct := struct {
		PIDx string `json:"pidx"`
	}{
		PIDx: pIDx,
	}

	payload, err := json.Marshal(payloadStruct)
	if err != nil {
		return false, nil, err
	}

	req, err := http.NewRequest("POST", verificationUrl, bytes.NewBuffer(payload))

	req.Header.Set("Authorization", "Key "+viper.GetString("paymentGateways.khalti.secretKey"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, res, err
	}

	defer resp.Body.Close()

	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	return true, res, nil
}
