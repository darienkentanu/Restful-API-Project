package payment

import (
	"altastore/constants"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type M map[string]interface{}

func RequestBilling(orderID int, amount int) (redirectURL string, err error) {

	url := "https://app.sandbox.midtrans.com/snap/v1/transactions"
	method := "POST"

	// 	payload := strings.NewReader(`{
	//   "transaction_details": {
	//     "order_id": "ORDER-101-12345678",
	//     "gross_amount": 10000
	//   },
	//   "credit_card": {
	//     "secure": true
	//   }
	// }`)

	payload, err := json.Marshal(M{
		"transaction_details": M{"order_id": orderID, "gross_amount": amount},
		"credit_card":         M{"secure": true},
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", constants.SERVER_KEY)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(body))
	temp2 := string(body)
	temp1 := strings.Index(temp2, "https")
	return string(body[temp1 : len(body)-2]), nil
}
