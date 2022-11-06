package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    bool   `json:"gender"`
	Age       int8   `json:"age"`
	Address   string `json:"address"`
	//Sku    string  `json:"sku" validate:"required,sku"` //the tag "sku" is there so we can add custom validation
}

type Users []*User

func (p *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *User) ToBson() (doc *bson.D, err error) {
	data, err := bson.Marshal(p)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

//Function to validate the incoming object from front.
//NOTE: if the tag "sku" is not present in the struct anotations we will get an error

//func (p *User) Validate() error {
//	validate := validator.New()
//
//	err := validate.RegisterValidation("sku", validateSKU)
//	if err != nil {
//		return err
//	}
//
//	return validate.Struct(p)
//}

//We use a regex to validate a custom look of sku-s
//For eg: abc-abc-abc

//func validateSKU(fl validator.FieldLevel) bool {
//	re := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]")
//	matches := re.FindAllString(fl.Field().String(), -1)
//
//	if len(matches) != 1 {
//		return false
//	}
//
//	return true
//}
