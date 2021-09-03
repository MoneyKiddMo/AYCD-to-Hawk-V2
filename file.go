package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func readJson() (*[]Hawkmesh, error) {
	var aycd []Aycd
	var hawkProfiles []Hawkmesh
	var houseNum string
	var billingNum string

	aycdFile, err := os.Open("aycd.json")
	if err != nil {
		color.Red("[ERROR] file.readJson Error Opening aycd.json: %s", err)
		return nil, err
	}
	color.HiGreen("[SUCCESS] Successfully Opened AYCD Import.")
	defer aycdFile.Close()

	aycdData, _ := ioutil.ReadAll(aycdFile)

	json.Unmarshal(aycdData, &aycd)

	for _, p := range aycd {
		nameArray := strings.Fields(p.Shippingaddress.Name)
		billingNameArray := strings.Fields(p.Billingaddress.Name)

		re := regexp.MustCompile("[0-9]+")
		houseNumber := re.FindAllString(p.Shippingaddress.Line1, -1)
		if len(houseNumber) > 0 {
			houseNum = houseNumber[0]
		} else {
			houseNumber2 := re.FindAllString(p.Shippingaddress.Line2, -1)
			if len(houseNumber2) > 0 {
				houseNum = houseNumber2[0]
			} else {
				color.Red("[ERROR] Error Parsing House Number, Setting as 0...")
				houseNum = "0"
			}
		}

		billingNumber := re.FindAllString(p.Billingaddress.Line1, -1)
		if len(billingNumber) > 0 {
			billingNum = billingNumber[0]
		} else {
			billingNumber2 := re.FindAllString(p.Billingaddress.Line2, -1)
			if len(billingNumber2) > 0 {
				billingNum = billingNumber2[0]
			} else {
				color.Red("[ERROR] Error Parsing House Number, Setting as 0...")
				houseNum = "0"
			}
		}

		hawkProfiles = append(hawkProfiles, Hawkmesh{
			Name: p.Name,
			Shipping: Shipping{
				Firstname:    nameArray[0],
				Lastname:     nameArray[1],
				Country:      countries[p.Shippingaddress.Country],
				Addressline:  p.Shippingaddress.Line1,
				Addressline2: p.Shippingaddress.Line2,
				City:         p.Shippingaddress.City,
				Postalcode:   p.Shippingaddress.Postcode,
				Phone:        p.Shippingaddress.Phone,
				State:        stateparser[p.Shippingaddress.State],
				Housenumber:  houseNum,
			},
			Email:    p.Shippingaddress.Email,
			Favorite: false,
			Billing: Billing{
				Firstname:    billingNameArray[0],
				Lastname:     billingNameArray[1],
				Country:      countries[p.Billingaddress.Country],
				Addressline:  p.Billingaddress.Line1,
				Addressline2: p.Billingaddress.Line2,
				City:         p.Billingaddress.City,
				Postalcode:   p.Billingaddress.Postcode,
				Phone:        p.Billingaddress.Phone,
				State:        stateparser[p.Billingaddress.State],
				Housenumber:  billingNum,
			},
			Billingsameasshipping: p.Samebillingandshippingaddress,
			Paymenttype:           "CreditCard",
			Creditcard: Creditcard{
				Holder:   p.Paymentdetails.Nameoncard,
				Number:   p.Paymentdetails.Cardnumber,
				Expmonth: expmonth[p.Paymentdetails.Cardexpmonth],
				Expyear:  expyear[p.Paymentdetails.Cardexpyear],
				Cvv:      p.Paymentdetails.Cardcvv,
			},
			Onecheckout:            p.Onlycheckoutonce,
			Randomemail:            false,
			Randomfirstname:        false,
			Randomphone:            false,
			Randomsecondaryaddress: false,
			Randomlastname:         false,
		})
	}

	return &hawkProfiles, nil
}

func writeHawk(hawk *[]Hawkmesh) error {

	jsonHawk, err := json.MarshalIndent(hawk, "", "")
	if err != nil {
		color.Red("[ERROR] file.writeHawk.json.MarshalIndent Error Marshalling to Hawk File: %s", err)
		return err
	}

	ioutil.WriteFile("hawkexport.json", jsonHawk, 0644)
	color.HiGreen("[SUCCESS] Successfully Converted Profiles.")

	return nil
}
