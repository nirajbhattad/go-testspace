package gochallenges

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Redaction2 interface {
	MarshalJson() (res []byte, err error)
}

// RedactTest will be the struct tag to be used to select te fiedls to redact
type RedactTest struct {
	SSN            string         `json:"ssn" redact:"ssn"`
	UserName       string         `json:"userName"`
	Password       string         `json:"password" redact:"password"`
	Address        string         `json:"address"`
	RedactInternal RedactInternal `json:"redactInternal"`
}

// Redact will be the struct tag to be used to select te fiedls to redact
type RedactInternal struct {
	SSN      string `json:"ssn" redact:"ssn"`
	UserName string `json:"userName"`
	Password string `json:"password" redact:"password"`
	Address  string `json:"address"`
}

// Generic implementation for Redaction
func (u *RedactTest) MarshalJSON() ([]byte, error) {

	uType := reflect.TypeOf(u).Elem()

	var fields []reflect.StructField
	for i := 0; i < uType.NumField(); i++ {
		f := uType.Field(i)
		if dbkey, ok := f.Tag.Lookup("redact"); ok && dbkey != "" {
			f.Tag = reflect.StructTag(fmt.Sprintf(`json:"%s"`, dbkey))
			f.Anonymous = true
		}
		fields = append(fields, f)
	}
	fmt.Println(fields)

	newType := reflect.StructOf(fields)
	fmt.Println(newType)

	initInstance := reflect.TypeOf(u).Elem()

	origInstance := reflect.ValueOf(u)
	if origInstance.Kind() == reflect.Ptr {
		origInstance = origInstance.Elem()
	}
	count := origInstance.NumField()
	fmt.Println(count)

	fmt.Println(initInstance)
	fmt.Println(origInstance)

	newTypeInstance := reflect.New(initInstance).Elem()
	fmt.Println(newTypeInstance)

	// Copy over the values
	for i := 0; i < origInstance.NumField(); i++ {
		fmt.Println(origInstance.Field(i))
		if dbkey, ok := initInstance.Field(i).Tag.Lookup("redact"); ok && dbkey != "" {
			origInstance.Field(i).SetString("******")
			newTypeInstance.Field(i).Set(origInstance.Field(i))
			fmt.Println(newTypeInstance)
		} else {
			newTypeInstance.Field(i).Set(origInstance.Field(i))
		}
	}
	fmt.Println(newTypeInstance)
	return json.Marshal(newTypeInstance.Interface())
}

// How the client can call the struct
func TestMarshalJson() {

	redactInternal := RedactInternal{
		SSN:      "123456789",
		UserName: "Mounika",
		Password: "123456789",
		Address:  "123456789",
	}

	// Input Object
	testMask := RedactTest{
		SSN:            "123456789",
		UserName:       "Mounika",
		Password:       "123456789",
		Address:        "123456789",
		RedactInternal: redactInternal,
	}

	jsonByte, err := json.Marshal(&testMask)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonByte))
}

/*

// Get all the fields of struct using reflection
// Select all the fields with tag redact using reflection
// Copy over the values from original object and redact the fields with redact tag
// Return the marshal of new object with redacted values

	uType := reflect.TypeOf(req).Elem()

	var fields []reflect.StructField

	// Scan and collect all fields converting any that have a "db" tag
	for i := 0; i < uType.NumField(); i++ {
		f := uType.Field(i)
		if dbkey, ok := f.Tag.Lookup("db"); ok && dbkey != "" {
			f.Tag = reflect.StructTag(fmt.Sprintf(`json:"%s"`, dbkey))
		}
		fields = append(fields, f)
	}

	newType := reflect.StructOf(fields)
	newTypeInstance := reflect.New(newType).Elem()
	origInstance := reflect.ValueOf(req).Elem()

	// Copy over the values
	for i := 0; i < origInstance.NumField(); i++ {
		newTypeInstance.Field(i).Set(origInstance.Field(i))
	}

	return json.Marshal(newTypeInstance.Interface())
*/
