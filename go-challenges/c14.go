package gochallenges

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
)

// Redact will be the struct tag to be used to select the fields to redact
type JsonRedact struct {
	SSN        string `json:"ssn" redact:""`
	UserName   string `json:"userName" `
	Password   string `json:"password" redact:""`
	Address    string `json:"address"`
	AccountNum int    `json:"accountNum" redact:""`
}

type XmlRedact struct {
	Height float32 `redact:""`
}

type RedactError struct {
	Msg string `json:"msg" redact:""`
}

func DebugRedaction() {

	// redactJson := JsonRedact{
	// 	SSN:        "123456789",
	// 	UserName:   "Mounika",
	// 	Password:   "123456789",
	// 	Address:    "123456789",
	// 	AccountNum: 1234567,
	// } // json struct pointer

	// xml struct pointer
	xmlRedact := &XmlRedact{Height: 5.5}

	// errorMessage := RedactError{Msg: "Unimplemented Type"}
	// print, _ := json.Marshal(errorMessage)
	// fmt.Println(string(print))

	// var tempInt interface{}
	// json.Unmarshal([]byte("{'RedactError':'Unimplemented Type'}"), &tempInt)
	// fmt.Println(tempInt)

	redactedXml := PrintXML(xmlRedact)
	fmt.Println(string(redactedXml))
}

// redact as json string
func Print(v interface{}) string {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	//pass interface and identify type
	switch val.Kind() {
	case reflect.Struct:
		redacted := RedactionToJson(v)
		return string(redacted)

	default:
		return "{'RedactError':'Unimplemented Type'}"
	}
}

// redact as XML string
func PrintXML(v interface{}) string {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	//pass interface and identify type
	switch val.Kind() {
	case reflect.Struct:
		redacted := RedactionToXml(v)
		return string(redacted)

	default:
		return fmt.Sprintf("<RedactError> Unimplemented Type< RedactError>")
	}
}

func RedactionToJson(req interface{}) (redacted []byte) {

	// if target is not pointer, then immediately return
	// modifying struct's field requires addressable object
	addrValue := reflect.ValueOf(req)
	if addrValue.Kind() != reflect.Ptr {
		return []byte("struct type not a pointer, not addressable")
	}

	// access the values stored in struct using reflection
	inputValue := addrValue.Elem()
	if !inputValue.IsValid() {
		return []byte("struct type does not have a valid value")
	}
	fmt.Println(inputValue)

	// Find the struct type
	inputType := inputValue.Type()
	fmt.Println(inputType)

	// New Instance Of Same Type
	outputInstance := reflect.New(inputType).Elem()

	// Loops over the struct fields, finds the redact tags, redacts the values.
	for i := 0; i < inputType.NumField(); i++ {
		fieldType := inputType.Field(i)
		fieldValue := inputValue.Field(i)

		if !fieldValue.IsValid() {
			continue
		}

		if !fieldValue.CanAddr() {
			// Cannot take pointer of this field, so can't scrub it.
			continue
		}

		if !fieldValue.Addr().CanInterface() {
			// This is an unexported or private field which begins with lowercase
			continue
		}

		if _, ok := fieldType.Tag.Lookup("redact"); ok {
			if fieldValue.CanSet() {
				// Checks For String Kind
				if fieldValue.Kind() == reflect.String && !fieldValue.IsZero() {
					outputInstance.Field(i).Set(inputValue.Field(i))
					outputInstance.Field(i).SetString("******")
					// Checks For Integer Kind
				} else if fieldValue.Kind() == reflect.Int && !fieldValue.IsZero() {
					outputInstance.Field(i).Set(inputValue.Field(i))
					outputInstance.Field(i).SetInt(00000)
				} else if fieldValue.Kind() == reflect.Float32 && !fieldValue.IsZero() {
					outputInstance.Field(i).Set(inputValue.Field(i))
					outputInstance.Field(i).SetFloat(00000)
				} else {
					continue
				}
			}
		} else {
			outputInstance.Field(i).Set(inputValue.Field(i))
		}
	}

	redacted, _ = json.Marshal(outputInstance.Interface())
	redactedXml, _ := xml.Marshal(outputInstance.Interface())

	fmt.Println(redactedXml)
	// Returns the redacted string.
	return redacted
}

func RedactionToXml(req interface{}) (redacted []byte) {

	// if target is not pointer, then immediately return
	// modifying struct's field requires addressable object
	addrValue := reflect.ValueOf(req)
	if addrValue.Kind() != reflect.Ptr {
		return []byte("struct type not a pointer, not addressable")
	}

	// access the values stored in struct using reflection
	inputValue := addrValue.Elem()
	if !inputValue.IsValid() {
		return []byte("struct type does not have a valid value")
	}
	fmt.Println(inputValue)

	// Find the struct type
	inputType := inputValue.Type()
	fmt.Println(inputType)

	// New Instance Of Same Type
	outputInstance := reflect.New(inputType).Elem()

	// Loops over the struct fields, finds the redact tags, redacts the values.
	for i := 0; i < inputType.NumField(); i++ {
		fieldType := inputType.Field(i)
		fieldValue := inputValue.Field(i)

		if !fieldValue.IsValid() {
			continue
		}

		if !fieldValue.CanAddr() {
			// Cannot take pointer of this field, so can't scrub it.
			continue
		}

		if !fieldValue.Addr().CanInterface() {
			// This is an unexported or private field which begins with lowercase
			continue
		}

		if _, ok := fieldType.Tag.Lookup("redact"); ok {
			if fieldValue.CanSet() {
				// Checks For String Kind
				if fieldValue.Kind() == reflect.String && !fieldValue.IsZero() {
					outputInstance.Field(i).Set(inputValue.Field(i))
					outputInstance.Field(i).SetString("******")
					// Checks For Integer Kind
				} else if fieldValue.Kind() == reflect.Int && !fieldValue.IsZero() {
					outputInstance.Field(i).Set(inputValue.Field(i))
					outputInstance.Field(i).SetInt(00000)
				} else if fieldValue.Kind() == reflect.Float32 && !fieldValue.IsZero() {
					outputInstance.Field(i).Set(inputValue.Field(i))
					outputInstance.Field(i).SetFloat(0.0)
				} else {
					continue
				}
			}
		} else {
			outputInstance.Field(i).Set(inputValue.Field(i))
		}
	}

	redacted, _ = xml.Marshal(outputInstance.Interface())
	// Returns the redacted string.
	return redacted
}
