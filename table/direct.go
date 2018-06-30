package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	//reportsURL  = "https://api.direct.yandex.com/json/v5/reports"
	reportsURL  = "https://api-sandbox.direct.yandex.com/json/v5/reports"
	token       = "AQAAAAAVc3FOAAUU82sT6LAmkE7yq_zbH_EAsfQ"
	clientLogin = "climbingspb"
)

const (
	headers = `
{
	"Authorization": "Bearer "` + token + `,
	"Client-Login": ` + clientLogin + `,
	"Accept-Language": "ru",
	"processingMode": "auto"
}
`
	body = `
{
	"params": {
		"SelectionCriteria": {
			"DateFrom": "2018-06-10",
			"DateTo": "2018-06-28"
		},
		"FieldNames": [
			"Date",
			"CampaignName",
			"LocationOfPresenceName",
			"Impressions",
			"Clicks",
			"Cost"
		],
		"ReportName": "Тестовый отчёт",
		"ReportType": "CAMPAIGN_PERFORMANCE_REPORT",
		"DateRangeType": "CUSTOM_DATE",
		"Format": "TSV",
		"IncludeVAT": "NO",
		"IncludeDiscount": "NO"
	}
}
`
)

func getDirectReport(tryOffline bool) string {
	var jsonStr = []byte(body)

	req, err := http.NewRequest("POST", reportsURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Client-Login", clientLogin)
	req.Header.Set("Accept-Language", "ru")
	req.Header.Set("processingMode", "auto")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	if (resp.StatusCode == 201 || resp.StatusCode == 202) && tryOffline {
		seconds, err := strconv.Atoi(resp.Header.Get("RetryIn"))
		if err != nil {
			seconds = 1
		}
		time.Sleep(time.Duration(seconds) * time.Second)

		return getDirectReport(false)
	}

	return string(body)

}
