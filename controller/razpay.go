package controller

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"

	"github.com/razorpay/razorpay-go"
)

func Executerazorpay() (string, error) {

	client := razorpay.NewClient("rzp_test_BBPTzEI8vRyL49", "1isUkJ6PXLSJcwgXhiQAeTFO")

	data := map[string]interface{}{
		"amount":   500,
		"currency": "INR",
		"receipt":  "101",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return "", errors.New("Payment not initiated")
	}
	razorId, _ := body["id"].(string)
	return razorId, nil
}

func RazorPaymentVerification(sign, orderId, paymentId string) error {
	signature := sign
	secret := "1isUkJ6PXLSJcwgXhiQAeTFO"
	data := orderId + "|" + paymentId

	h := hmac.New(sha256.New, []byte(secret))

	_, err := h.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(signature)) != 1 {
		return errors.New("Payment failed")
	} else {
		return nil
	}
}
