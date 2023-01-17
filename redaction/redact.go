package redaction

import (
	"encoding/json"
	"encoding/xml"
	"reflect"

	"sync"
)

var (
	mu sync.RWMutex
	// To keep a count of fields which have redact tag
	redactFieldCount int
	// To keep a track of whether redaction worked fine
	hasRedactionError bool
)

func getPointerToStruct(obj interface{}) interface{} {
	// check if the value is pointer
	reqAddrValue := reflect.ValueOf(obj)
	if reqAddrValue.Kind() == reflect.Ptr {
		return obj
	} else {
		// Create a new instance of the underlying type
		vp := reflect.New(reflect.TypeOf(obj))

		// Set the original value to the new instance
		vp.Elem().Set(reflect.ValueOf(obj))

		// Gives a pointer to struct
		return vp.Interface()
	}
}

func RedactJson(req interface{}) string {
	if req == nil {
		return ""
	}

	out := getPointerToStruct(req)

	// Declare original values slice
	originalValues := make([]interface{}, 0)

	redact(out, &originalValues, true, false)

	// Create a json redaction copy
	var jsonRedaction []byte
	jsonRedaction, _ = json.Marshal(out)

	// Reads Redaction Error State
	hasRedactError := getRedactionErrorState()
	redactionFieldCount := getRedactionFieldCount()

	if redactionFieldCount == 0 {
		// Call alternate flow
		return RedactedLog(req, true)
	} else if redactionFieldCount > 0 && hasRedactError {
		// Get back to original state
		redact(out, &originalValues, false, false)
		return RedactedLog(req, true)
	} else {
		// Get back to original state
		redact(out, &originalValues, false, false)
		// Returns
		return string(jsonRedaction)
	}
}

// First is to check if there's a way to check if at all the input struct has redact tag. " RedactCount == 0"
// RedactCount should only be updated in case of Redaction for all the fields, i.e just once.
// We will need another flag to help us identify if at all the redaction went through properly.

func RedactXml(req interface{}) string {
	if req == nil {
		return "<Redact><error>Empty request</error></Redact>"
	}

	out := getPointerToStruct(req)

	// Declare original values slice
	originalValues := make([]interface{}, 0)

	// Redact the json
	redact(out, &originalValues, true, false)

	// Creates a xml redaction copy
	var xmlRedaction []byte
	xmlRedaction, _ = xml.Marshal(out)

	// Reads Redaction Error State
	hasRedactError := getRedactionErrorState()
	redactionFieldCount := getRedactionFieldCount()

	if redactionFieldCount == 0 {
		// Call alternate flow
		return RedactedLog(req, false)
	} else if redactionFieldCount > 0 && hasRedactError {
		// Get back to original state
		redact(out, &originalValues, false, false)
		// Call alternate flow
		return RedactedLog(req, false)
	} else {
		// Get back to original state
		redact(out, &originalValues, false, false)
		return string(xmlRedaction)
	}
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

	// if field/struct is passed by pointer, then first dereference it to get the
	// underlying value (the pointer must not be pointing to a nil value)
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

			_, shouldRedact := fieldType.Tag.Lookup("redact")
			if shouldRedact && save {
				setRedactionFieldCount()
			}

			if !fValue.IsValid() {
				failedToRedact(shouldRedact)
				continue
			}

			if !fValue.CanAddr() {
				failedToRedact(shouldRedact)
				continue
			}

			if !fValue.Addr().CanInterface() {
				failedToRedact(shouldRedact)
				continue
			}

			redact(fValue.Addr().Interface(), originalValues, save, shouldRedact)
		}
		return
	}

	if requestType.Kind() == reflect.Array || requestType.Kind() == reflect.Slice {
		if isRedact && save {
			setRedactionFieldCount()
		}
		for i := 0; i < inputValue.Len(); i++ {
			arrValue := inputValue.Index(i)
			if !arrValue.IsValid() {
				failedToRedact(isRedact)
				continue
			}

			if !arrValue.CanAddr() {
				failedToRedact(isRedact)
				continue
			}

			if !arrValue.Addr().CanInterface() {
				failedToRedact(isRedact)
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
	} else if inputValue.Kind() == reflect.Int {
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
	} else if inputValue.Kind() == reflect.Int32 {
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
	} else if inputValue.Kind() == reflect.Int64 {
		if save {
			*originalValues = append(*originalValues, inputValue.Int())
			if isRedact {
				inputValue.SetInt(0)
			}
		} else {
			inputValue.SetInt((*originalValues)[0].(int64))
			*originalValues = (*originalValues)[1:]
		}
	} else if inputValue.Kind() == reflect.Float64 {
		if save {
			*originalValues = append(*originalValues, inputValue.Float())
			if isRedact {
				inputValue.SetFloat(0.0)
			}
		} else {
			inputValue.SetFloat((*originalValues)[0].(float64))
			*originalValues = (*originalValues)[1:]
		}
	} else if inputValue.Kind() == reflect.Float32 {
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

// RedactedLog writes returns a string and
// attempts to redact any PII fields first.
func RedactedLog(value interface{}, isJson bool) string {
	switch value := value.(type) {
	case []byte:
		return string(value)
		//return string(zaputil.RedactJSON(value))
	case string:
		return value
	default:
		return MarshalRedactPayload(value, isJson)
	}
}

// MarshalRedactPayload returns a string after attempting to redact the struct
func MarshalRedactPayload(value interface{}, isJson bool) string {
	var buf []byte
	var err error
	if isJson {
		buf, err = json.Marshal(value)
	} else {
		buf, err = xml.Marshal(value)
	}

	if err != nil {
		return err.Error()
	}

	return string(buf)
}

func failedToRedact(failedToRedact bool) {
	if failedToRedact {
		setRedactionErrorState()
	}
}

// Getter For Redaction Field Count
func getRedactionFieldCount() int {
	mu.RLock()
	defer mu.RUnlock()
	me := redactFieldCount
	return me
}

// Setter For Redaction Field Count
func setRedactionFieldCount() {
	mu.Lock()
	defer mu.Unlock()
	redactFieldCount++
}

// Getter For Redaction Error State
func getRedactionErrorState() bool {
	mu.RLock()
	defer mu.RUnlock()
	me := hasRedactionError
	return me
}

// Setter For Redaction Error State
func setRedactionErrorState() {
	mu.Lock()
	defer mu.Unlock()
	hasRedactionError = true
}
