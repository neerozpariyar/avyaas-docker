package usecase

import (
	"avyaas/internal/domain/presenter"

	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

func verifyEsewaPayment(totalAmount int, transactionUUID string, refID string) (success bool, res *presenter.VerifyPaymentStatus, err error) {
	verificationUrl := fmt.Sprintf(
		"%v/?product_code=%v&total_amount=%v&transaction_uuid=%v",
		viper.GetString("paymentGateways.esewa.url"), viper.GetString("paymentGateways.esewa.productID"), totalAmount, transactionUUID)

	response, err := http.Get(verificationUrl)
	if err != nil {
		return false, nil, err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return false, nil, err
	}

	err = json.Unmarshal(responseBody, &res)
	if err != nil {
		return false, nil, err
	}

	res.ReferenceID = refID

	return true, res, nil
}
