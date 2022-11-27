package gochallenges

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
)

type RecursionRedact2 struct {
	UserName        string           `json:"userName" `
	Password        string           `json:"password" redact:""`
	RecursionRedact RecursionRedact3 `json:"recursionRedact" redact:""`
}

type RecursionRedact3 struct {
	SSN           *int     `json:"ssn" redact:""`
	AccountNum    *[]int32 `json:"accountNum" redact:""`
	AccountNumber []int64  `json:"accountNumber" redact:""`
	AccountPer    float32  `json:"accountPer" redact:""`
	AccountsPer   float64  `json:"accountsPer" redact:""`
}

func DebugRecursion() {
	ssn := 1234567
	ssnarray := []int32{1234}
	redactJson3 := RecursionRedact3{
		SSN:           &ssn,
		AccountNum:    &ssnarray,
		AccountNumber: []int64{888888, 123456},
		AccountPer:    23.056,
		AccountsPer:   12345.5678,
	}

	redactJson := RecursionRedact2{
		UserName:        "Niraj",
		Password:        "Bhattad",
		RecursionRedact: redactJson3,
	}

	Redact(&redactJson)
}

func Redact(req interface{}) {

	// Redact the json
	originalValues := make([]interface{}, 0)
	redact(req, &originalValues, true, false)

	// Create a json redaction copy
	var jsonRedaction []byte
	jsonRedaction, _ = json.Marshal(req)
	fmt.Println(string(jsonRedaction))

	// Create a xml redaction copy
	var xmlRedactionCopy []byte
	xmlRedactionCopy, _ = xml.Marshal(req)
	fmt.Println(string(xmlRedactionCopy))

	// Get back to original state
	redact(req, &originalValues, false, false)
}

func redact(req interface{}, originalValues *[]interface{}, save bool, isRedact bool) {

	// if target is not pointer, then immediately return
	// modifying struct's field requires addressable object
	reqAddrValue := reflect.ValueOf(req)
	if reqAddrValue.Kind() != reflect.Ptr {
		return
	}

	requestValue := reqAddrValue.Elem()
	if !requestValue.IsValid() {
		return
	}

	requestType := requestValue.Type()

	if requestType.Kind() == reflect.Struct {
		// If target is a struct then recurse on each of its field.
		for i := 0; i < requestType.NumField(); i++ {
			fieldType := requestType.Field(i)
			fValue := requestValue.Field(i)
			if !fValue.IsValid() {
				continue
			}

			if !fValue.CanAddr() {
				// Cannot take pointer of this field, so can't scrub it.
				continue
			}

			if !fValue.Addr().CanInterface() {
				continue
			}
			_, shouldRedact := fieldType.Tag.Lookup("redact")
			redact(fValue.Addr().Interface(), originalValues, save, shouldRedact)
		}
		return
	}

	if requestType.Kind() == reflect.Array || requestType.Kind() == reflect.Slice {
		for i := 0; i < requestValue.Len(); i++ {
			arrValue := requestValue.Index(i)
			if !arrValue.IsValid() {
				continue
			}

			if !arrValue.CanAddr() {
				continue
			}

			if !arrValue.Addr().CanInterface() {

				continue
			}

			redact(arrValue.Addr().Interface(), originalValues, save, isRedact)
		}

		return
	}

	// Base Condition To Return From Recursive Function
	if requestValue.CanSet() && requestValue.Kind() == reflect.String && !requestValue.IsZero() {
		if save {
			*originalValues = append(*originalValues, requestValue.String())
			if isRedact {
				requestValue.SetString("********")
			}
		} else {
			requestValue.SetString((*originalValues)[0].(string))
			*originalValues = (*originalValues)[1:]
		}
	} else if requestValue.CanInt() && requestValue.Kind() == reflect.Int {
		if save {
			*originalValues = append(*originalValues, requestValue.Int())
			if isRedact {
				requestValue.SetInt(0)
			}
		} else {
			temp := int((*originalValues)[0].(int64))
			requestValue.Set(reflect.ValueOf(temp))
			*originalValues = (*originalValues)[1:]
		}
	} else if requestValue.CanInt() && requestValue.Kind() == reflect.Int32 {
		if save {
			*originalValues = append(*originalValues, requestValue.Int())
			if isRedact {
				requestValue.SetInt(0)
			}
		} else {
			temp := int32((*originalValues)[0].(int64))
			requestValue.Set(reflect.ValueOf(temp))
			*originalValues = (*originalValues)[1:]
		}
	} else if requestValue.CanInt() && requestValue.Kind() == reflect.Int64 {
		if save {
			*originalValues = append(*originalValues, requestValue.Int())
			if isRedact {
				requestValue.SetInt(0)
			}
		} else {
			requestValue.SetInt((*originalValues)[0].(int64))
			*originalValues = (*originalValues)[1:]
		}
	} else if requestValue.CanFloat() && requestValue.Kind() == reflect.Float64 {
		if save {
			*originalValues = append(*originalValues, requestValue.Float())
			if isRedact {
				requestValue.SetFloat(0.0)
			}
		} else {
			requestValue.SetFloat((*originalValues)[0].(float64))
			*originalValues = (*originalValues)[1:]
		}
	} else if requestValue.CanFloat() && requestValue.Kind() == reflect.Float32 {
		if save {
			*originalValues = append(*originalValues, requestValue.Float())
			if isRedact {
				requestValue.SetFloat(0.0)
			}
		} else {
			temp := float32((*originalValues)[0].(float64))
			requestValue.Set(reflect.ValueOf(temp))
			*originalValues = (*originalValues)[1:]
		}
	}
}
