package mms

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"project/status/check"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func MMSChangeCodeToCountry(data []MMSData) []MMSData {
	res := make([]MMSData, 0)

	for _, elem := range data {
		country := check.GetCountryForCode(elem.Country)
		elem.Country = country
		res = append(res, elem)
	}

	return res
}

func CheckMMSData(data []MMSData) []MMSData {
	result := make([]MMSData, 0)

	for _, elem := range data {
		if !check.IsCountry(elem.Country) {
			continue
		}

		if !check.IsBandwidth(elem.Bandwidth) {
			continue
		}

		if !check.IsResponseTime(elem.ResponseTime) {
			continue
		}

		if !check.IsProviderSMSandMMS(elem.Provider) {
			continue
		}

		result = append(result, elem)
	}

	return result
}

func StatusMMS(url string) []MMSData {
	result := make([]MMSData, 0)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error() + `: ` + url)
		return []MMSData{}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []MMSData{}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return []MMSData{}
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return []MMSData{}
	}

	return CheckMMSData(result)
}
