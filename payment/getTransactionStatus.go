package payment

import (
	"altastore/constants"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getTransactionStatus(orderID int) (status int, err error) {

	url := "https://api.sandbox.midtrans.com/v2/ORDER-101-1234567/status"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", constants.SERVER_KEY)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println(string(body))
	paymentSuccess := strings.Index(string(body), "settlement")
	if paymentSuccess > 0 {
		return 1, nil
	}
	return 0, nil
}
