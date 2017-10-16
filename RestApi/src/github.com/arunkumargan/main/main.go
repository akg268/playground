package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

//Homeloans struct
type Homeloans struct {
	ProductName                string  `json:"productName"`
	PrincipalBalance           float64 `json:"principalBalance"`
	AvailableCreditLimit       float64 `json:"availableCreditLimit"`
	ProductTypeCode            string  `json:"productTypeCode"`
	NextPaymentDueDate         string  `json:"nextPaymentDueDate"`
	InternetStopCode           string  `json:"internetStopCode"`
	PaidInFullStopCode         string  `json:"paidInFullStopCode"`
	DaysPastDue                string  `json:"daysPastDue"`
	ForeclosureStopCode        string  `json:"foreclosureStopCode"`
	BadCheckCode               string  `json:"badCheckCode"`
	AccountOpenDate            string  `json:"accountOpenDate"`
	AccountNumber              string  `json:"accountNumber"`
	IsRecurringPaymentDraft    bool    `json:"isRecurringPaymentDraft"`
	GracePeriodDays            int     `json:"gracePeriodDays"`
	PropertyStateCode          string  `json:"propertyStateCode"`
	IsRecurringPaymentEligible bool    `json:"isRecurringPaymentEligible"`
}

//Accounts struct
type Accounts struct {
	AccountNumber string `json:"accountNumber"`
}

//Homeloanslob Handler function
func Homeloanslob(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method received is :", r.Method)
	//have to add insecure skip verify until we get a client cert for the wiremock.
	client := &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	urlstring, _ := url.Parse("https://entitlements-wiremock-east.cloud.capitalone.com/loans-homeloans-accounts-web/loans/home-loans/accounts/")
	accrefid := "LpLECMx7YAF6KBLuI2dtMiJJ3%2BbzlHHXFDSfwqIg8Ps%3D"
	encodedaccrefid := url.URL{Path: accrefid}
	endpointurl := urlstring.String() + encodedaccrefid.String()
	fmt.Println("calling mock service:", endpointurl)
	req, _ := http.NewRequest("GET", endpointurl, nil)
	req.Header.Set("Accept", "application/json;v=3")
	req.Header.Set("Api-Key", "MOBILE")
	req.Header.Set("Content-Type", "application/json;v=3")
	fmt.Println("calling wiremock:", req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	fmt.Println("response code:", resp.StatusCode)

	switch resp.StatusCode {
	case 200:
		//Another way of fabricating resp
		//data, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	fmt.Println(err)
		//	}
		//	w.Write(data)
		var homeloans = new(Homeloans)
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(bytes.NewBuffer(data).String())
		json.Unmarshal(data, &homeloans)
		json.NewEncoder(w).Encode(&homeloans)
	case 404:
		w.WriteHeader(http.StatusNotFound)
		data, _ := ioutil.ReadAll(resp.Body)
		w.Write(data)
	case 500:
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := ioutil.ReadAll(resp.Body)
		w.Write(data)
	default:
		w.WriteHeader(http.StatusNoContent)
	}

}

//Postretail function to make a post call to retail lob wiremock
func Postretail(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method received is :", r.Method)
	client := &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	urlstring := "https://entitlements-wiremock-east.cloud.capitalone.com/int-retail-accounts/integration/retail/accounts"
	fmt.Println("calling mock service:", urlstring)
	accounts1 := Accounts{AccountNumber: "1205899801"}
	accounts2 := Accounts{AccountNumber: "1357699213"}
	rdata := []Accounts{accounts1, accounts2}
	reqdata, _ := json.Marshal(&rdata)
	req, _ := http.NewRequest("POST", urlstring, bytes.NewBuffer(reqdata))
	req.Header.Set("Accept", "application/json;v=3")
	req.Header.Set("Api-Key", "MOBILE")
	req.Header.Set("Content-Type", "application/json;v=3")

	fmt.Println("calling wiremock:", req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	fmt.Println("response code:", resp.StatusCode)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(data)
}

func main() {
	fmt.Println("this is a main function")
	go http.HandleFunc("/homeloanslob", Homeloanslob)
	go http.HandleFunc("/retaillob", Postretail)
	err := http.ListenAndServe(":8191", nil)
	//err := http.ListenAndServeTLS(":8443", "eil-qaid.pem", "eil-qaid.key", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Service is Running at port 8191")
}
