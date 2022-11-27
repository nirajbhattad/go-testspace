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
	Accounts   []int  `json:"accounts" redact:""`
	Redacts    []XmlRedact
}

type XmlRedact struct {
	Height float32 `redact:""`
	Weight float32 `redact:""`
}

type RedactError struct {
	Msg string `json:"msg" redact:""`
}

func DebugRedaction() {
	// // json struct pointer
	// redactJson := JsonRedact{

	// 	Accounts: []int{1, 2, 3, 4},
	// }

	// // // xml struct pointer
	// // xmlRedact := &XmlRedact{Height: 5.5}

	// redactedJson := Print(redactJson)
	// fmt.Println(string(redactedJson))

	accounts := []string{"1", "2", "3", "4"}
	redactedSlice := RedactSlice(&accounts)
	fmt.Println(string(redactedSlice))

	// redactedXml := PrintXML(xmlRedact)
	// fmt.Println(string(redactedXml))
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
		return "<RedactError> Unimplemented Type< RedactError>"
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
					outputInstance.Field(i).Set(fieldValue)
					outputInstance.Field(i).SetString("******")
					// Checks For Integer Kind
				} else if fieldValue.Kind() == reflect.Int && !fieldValue.IsZero() {
					outputInstance.Field(i).Set(fieldValue)
					outputInstance.Field(i).SetInt(00000)
				} else if fieldValue.Kind() == reflect.Float32 && !fieldValue.IsZero() {
					outputInstance.Field(i).Set(fieldValue)
					outputInstance.Field(i).SetFloat(00000)
				} else if fieldValue.Kind() == reflect.Array || fieldValue.Kind() == reflect.Slice {
					// New Instance Of Same Type
					arrayInstance := reflect.New(fieldValue.Type()).Elem()
					for i := 0; i < fieldValue.Len(); i++ {
						arrValue := fieldValue.Field(i)
						if !arrValue.IsValid() {
							continue
						}

						if !arrValue.CanAddr() {
							continue
						}

						if !arrValue.Addr().CanInterface() {
							continue
						}

						if arrValue.Kind() == reflect.String && !arrValue.IsZero() {
							arrayInstance.Field(i).Set(arrValue)
							arrayInstance.Field(i).SetString("******")
							// Checks For Integer Kind
						} else if arrValue.Kind() == reflect.Int && !arrValue.IsZero() {
							arrayInstance.Field(i).Set(arrValue)
							arrayInstance.Field(i).SetInt(00000)
						} else if arrValue.Kind() == reflect.Float32 && !arrValue.IsZero() {
							arrayInstance.Field(i).Set(arrValue)
							arrayInstance.Field(i).SetFloat(00000)
						}
					}
					outputInstance.Field(i).Set(fieldValue)
					outputInstance.Field(i).Set(arrayInstance)
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

	for i := 0; i < inputType.NumField(); i++ {
		fieldType := inputType.Field(i)
		fieldValue := inputValue.Field(i)

		if !fieldValue.IsValid() {
			continue
		}

		if !fieldValue.CanAddr() {
			continue
		}

		if !fieldValue.Addr().CanInterface() {
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

func RedactSlice(req interface{}) (redacted []byte) {

	addrValue := reflect.ValueOf(req)
	if addrValue.Kind() != reflect.Ptr {
		return
	}

	inputValue := addrValue.Elem()
	if !inputValue.IsValid() {
		return
	}
	fmt.Println(inputValue)

	inputType := inputValue.Type()
	fmt.Println(inputType)

	var inputInstance []interface{}
	if inputType.Kind() == reflect.Array || inputType.Kind() == reflect.Slice {
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

			target := arrValue.Addr().Interface()
			fmt.Println(target)

			fieldAddrValue := reflect.ValueOf(target)
			if fieldAddrValue.Kind() != reflect.Ptr {
				return
			}
			fmt.Println(fieldAddrValue)

			fieldTargetValue := fieldAddrValue.Elem()
			if !fieldTargetValue.IsValid() {
				return
			}
			fmt.Println(fieldTargetValue)

			fieldTargetType := fieldTargetValue.Type()
			fmt.Println(fieldTargetType)

			if fieldTargetType.Kind() == reflect.Ptr && !fieldTargetValue.IsNil() {
				fieldTargetValue = fieldTargetValue.Elem()
				if !fieldTargetValue.IsValid() {
					return
				}

				fieldTargetType = fieldTargetValue.Type()
			}
			fmt.Println(fieldTargetValue)

			fieldInstance := reflect.New(fieldTargetType).Elem()
			if fieldTargetType.Kind() == reflect.String && !arrValue.IsZero() {
				fieldInstance.Set(fieldTargetValue)
				fieldInstance.SetString("******")
				// Checks For Integer Kind
			} else if fieldTargetType.Kind() == reflect.Int && !arrValue.IsZero() {
				fieldInstance.Set(fieldTargetValue)
				fieldInstance.SetInt(00000)
			} else if fieldTargetType.Kind() == reflect.Float32 && !arrValue.IsZero() {
				fieldInstance.Set(fieldTargetValue)
				fieldInstance.SetFloat(00000)
			}
			inputInstance = append(inputInstance, fieldInstance.Interface())
			fmt.Println(fieldInstance.Interface())
			fmt.Println(inputInstance)
		}
	}

	redacted, _ = json.Marshal(inputInstance)
	fmt.Println(string(redacted))
	return redacted
}
