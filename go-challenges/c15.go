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

// To keep a count of fields which have redact tag
var RedactCount int

func Redact(req interface{}) string {
	if req == nil {
		return ""
	}

	out := to_struct_ptr(req)

	// Declare original values slice
	originalValues := make([]interface{}, 0)

	// Redact the json
	redact(out, &originalValues, true, false)

	// Create a json redaction copy
	var jsonRedaction []byte
	jsonRedaction, _ = json.Marshal(out)
	fmt.Println(string(jsonRedaction))

	if RedactCount > 0 && string(jsonRedaction) == "" {
		return "{'RedactError':'Failed to redact the struct.'}"
	}

	// Get back to original state
	redact(out, &originalValues, false, false)

	// Returns
	return string(jsonRedaction)
}

func RedactXml(req interface{}) string {
	if req == nil {
		return "<Redact><error>Empty request</error></Redact>"
	}

	// Declare original values slice
	originalValues := make([]interface{}, 0)

	// Redact the json
	redact(req, &originalValues, true, false)

	// Creates a xml redaction copy
	var xmlRedaction []byte
	xmlRedaction, _ = xml.Marshal(req)

	if RedactCount > 0 && string(xmlRedaction) == "" {
		return "<RedactError> Failed to redact the struct.<RedactError>"
	}

	// Get back to original state
	redact(req, &originalValues, false, false)

	// Returns
	return string(xmlRedaction)
}

func redact(req interface{}, originalValues *[]interface{}, save bool, isRedact bool) {

	// if target is not pointer, then immediately return
	// modifying struct's field requires addressable object
	reqAddrValue := reflect.ValueOf(req)
	if reqAddrValue.Kind() != reflect.Ptr {
		return
	}

	inputValue := reqAddrValue.Elem()
	if !inputValue.IsValid() {
		return
	}

	requestType := inputValue.Type()

	// If the field/struct is passed by pointer, then first dereference it to get the
	// underlying value (the pointer must not be pointing to a nil value).
	if inputValue.Kind() == reflect.Ptr && !inputValue.IsNil() {
		inputValue = inputValue.Elem()
		if !inputValue.IsValid() {
			return
		}

		requestType = inputValue.Type()
	}

	if requestType.Kind() == reflect.Struct {
		// If target is a struct then recurse on each of its field.
		for i := 0; i < requestType.NumField(); i++ {
			fieldType := requestType.Field(i)
			fValue := inputValue.Field(i)
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
			if shouldRedact {
				RedactCount++
			}
			redact(fValue.Addr().Interface(), originalValues, save, shouldRedact)
		}
		return
	}

	if requestType.Kind() == reflect.Array || requestType.Kind() == reflect.Slice {
		if isRedact {
			RedactCount++
		}
		for i := 0; i < inputValue.Len(); i++ {
			arrValue := inputValue.Index(i)
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
	if inputValue.CanSet() && inputValue.Kind() == reflect.String && !inputValue.IsZero() {
		if save {
			*originalValues = append(*originalValues, inputValue.String())
			if isRedact {
				inputValue.SetString("********")
			}
		} else {
			inputValue.SetString((*originalValues)[0].(string))
			*originalValues = (*originalValues)[1:]
		}
	} else if inputValue.CanInt() && inputValue.Kind() == reflect.Int {
		if save {
			*originalValues = append(*originalValues, inputValue.Int())
			if isRedact {
				inputValue.SetInt(0)
			}
		} else {
			temp := int((*originalValues)[0].(int64))
			inputValue.Set(reflect.ValueOf(temp))
			*originalValues = (*originalValues)[1:]
		}
	} else if inputValue.CanInt() && inputValue.Kind() == reflect.Int32 {
		if save {
			*originalValues = append(*originalValues, inputValue.Int())
			if isRedact {
				inputValue.SetInt(0)
			}
		} else {
			temp := int32((*originalValues)[0].(int64))
			inputValue.Set(reflect.ValueOf(temp))
			*originalValues = (*originalValues)[1:]
		}
	} else if inputValue.CanInt() && inputValue.Kind() == reflect.Int64 {
		if save {
			*originalValues = append(*originalValues, inputValue.Int())
			if isRedact {
				inputValue.SetInt(0)
			}
		} else {
			inputValue.SetInt((*originalValues)[0].(int64))
			*originalValues = (*originalValues)[1:]
		}
	} else if inputValue.CanFloat() && inputValue.Kind() == reflect.Float64 {
		if save {
			*originalValues = append(*originalValues, inputValue.Float())
			if isRedact {
				inputValue.SetFloat(0.0)
			}
		} else {
			inputValue.SetFloat((*originalValues)[0].(float64))
			*originalValues = (*originalValues)[1:]
		}
	} else if inputValue.CanFloat() && inputValue.Kind() == reflect.Float32 {
		if save {
			*originalValues = append(*originalValues, inputValue.Float())
			if isRedact {
				inputValue.SetFloat(0.0)
			}
		} else {
			temp := float32((*originalValues)[0].(float64))
			inputValue.Set(reflect.ValueOf(temp))
			*originalValues = (*originalValues)[1:]
		}
	}
}

func to_struct_ptr(obj interface{}) interface{} {

	// Create a new instance of the underlying type
	vp := reflect.New(reflect.TypeOf(obj))

	vp.Elem().Set(reflect.ValueOf(obj))

	return vp.Interface()
}
