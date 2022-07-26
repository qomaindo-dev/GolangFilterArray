package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	filename = "datajson.json"
)

func main() {

	plan, _ := ioutil.ReadFile(filename)
	var data interface{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dataConvString := string(plan)

	hargaItem := GetHargaItem(dataConvString, "voucher10", "voucher", "Contoh Voucher")
	fmt.Println("Harga perItem: ", hargaItem)
}

func GetHargaItem(dataArray, productCode, productType, productDesc string) float64 {

	productDescReplace := strings.ReplaceAll(productDesc, " ", "_")

	fmt.Println("productType: ", productType)
	fmt.Println("productDescReplace: ", productDescReplace)

	var dataMap map[string]interface{}
	json.Unmarshal([]byte(dataArray), &dataMap)

	dataMap2 := dataMap["data"].(map[string]interface{})

	var totalHarga float64
	feeTransaction := 500.0
	for _, item := range dataMap2["pricelist"].([]interface{}) {
		itemMap := item.(map[string]interface{})
		if itemMap["product_code"] != nil && itemMap["product_code"].(string) != "" {
			if strings.ToLower(itemMap["product_code"].(string)) == strings.ToLower(productCode) {
				totalHarga += itemMap["product_price"].(float64) + feeTransaction
			} else {
				continue
			}
		}
	}

	return totalHarga
}
