package main

type Config struct {
	Profilesettings Profilesettings `json:"profileSettings"`
}
type Profilesettings struct {
	Randomphone     bool `json:"randomPhone"`
	Randomfirstname bool `json:"randomFirstname"`
	Randomlastname  bool `json:"randomLastname"`
	Randomline2     bool `json:"randomLine2"`
}
type Aycd struct {
	Name                          string          `json:"name"`
	Size                          string          `json:"size"`
	Profilegroup                  string          `json:"profileGroup"`
	Billingaddress                Billingaddress  `json:"billingAddress"`
	Shippingaddress               Shippingaddress `json:"shippingAddress"`
	Paymentdetails                Paymentdetails  `json:"paymentDetails"`
	Samebillingandshippingaddress bool            `json:"sameBillingAndShippingAddress"`
	Onlycheckoutonce              bool            `json:"onlyCheckoutOnce"`
	Matchnameoncardandaddress     bool            `json:"matchNameOnCardAndAddress"`
}
type Billingaddress struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Line1    string `json:"line1"`
	Line2    string `json:"line2"`
	Line3    string `json:"line3"`
	Postcode string `json:"postCode"`
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
}
type Shippingaddress struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Line1    string `json:"line1"`
	Line2    string `json:"line2"`
	Line3    string `json:"line3"`
	Postcode string `json:"postCode"`
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
}
type Paymentdetails struct {
	Nameoncard   string `json:"nameOnCard"`
	Cardtype     string `json:"cardType"`
	Cardnumber   string `json:"cardNumber"`
	Cardexpmonth string `json:"cardExpMonth"`
	Cardexpyear  string `json:"cardExpYear"`
	Cardcvv      string `json:"cardCvv"`
}

type Hawkmesh struct {
	Name                   string     `json:"Name"`
	Shipping               Shipping   `json:"Shipping"`
	Email                  string     `json:"Email"`
	Favorite               bool       `json:"Favorite"`
	Billing                Billing    `json:"Billing"`
	Billingsameasshipping  bool       `json:"BillingSameAsShipping"`
	Paymenttype            string     `json:"PaymentType"`
	Creditcard             Creditcard `json:"CreditCard"`
	Onecheckout            bool       `json:"OneCheckout"`
	Randomemail            bool       `json:"RandomEmail"`
	Randomfirstname        bool       `json:"RandomFirstname"`
	Randomphone            bool       `json:"RandomPhone"`
	Randomsecondaryaddress bool       `json:"RandomSecondaryAddress"`
	Randomlastname         bool       `json:"RandomLastname"`
}
type Shipping struct {
	Firstname    string      `json:"Firstname"`
	Lastname     string      `json:"Lastname"`
	Country      string      `json:"Country"`
	Addressline  string      `json:"AddressLine"`
	Addressline2 string      `json:"AddressLine2"`
	City         string      `json:"City"`
	Postalcode   string      `json:"Postalcode"`
	Phone        string      `json:"Phone"`
	State        interface{} `json:"State"`
	Housenumber  string      `json:"HouseNumber"`
}
type Billing struct {
	Firstname    string      `json:"Firstname"`
	Lastname     string      `json:"Lastname"`
	Country      string      `json:"Country"`
	Addressline  string      `json:"AddressLine"`
	Addressline2 string      `json:"AddressLine2"`
	City         string      `json:"City"`
	Postalcode   string      `json:"Postalcode"`
	Phone        string      `json:"Phone"`
	State        interface{} `json:"State"`
	Housenumber  string      `json:"HouseNumber"`
}
type Creditcard struct {
	Holder   string `json:"Holder"`
	Number   string `json:"Number"`
	Expmonth int    `json:"ExpMonth"`
	Expyear  int    `json:"ExpYear"`
	Cvv      string `json:"Cvv"`
}
