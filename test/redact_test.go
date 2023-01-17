package redact

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRedactSimple tests redaction on a simple struct with json
func TestRedactSimple(t *testing.T) {

	// Simple struct
	type User struct {
		Username      string   `json:"userName" `
		Password      string   `json:"password" redact:""`
		DbSecrets     []string `json:"dbSecrets"`
		SIN           *int     `json:"sin" redact:""`
		SSN           int      `json:"ssn" redact:""`
		AccountNum    int32    `json:"accountNum" redact:""`
		AccountNumber []int64  `json:"accountNumber" redact:""`
		AccountPer    float32  `json:"accountPer" redact:""`
		AccountsPer   float64  `json:"accountsPer" redact:""`
	}
	sin := 1234568
	expectedSin := 0

	user := User{
		Username:      "Redaction",
		Password:      "codepassword",
		DbSecrets:     []string{"db_secret_1", "db_secret_2"},
		SIN:           &sin,
		SSN:           1234567,
		AccountNum:    1234,
		AccountNumber: []int64{888888, 123456},
		AccountPer:    23.056,
		AccountsPer:   12345.5678,
	}

	userRedacted := &User{
		Username:      "Redaction",
		Password:      "********",
		DbSecrets:     []string{"db_secret_1", "db_secret_2"},
		SIN:           &expectedSin,
		SSN:           0,
		AccountNum:    0,
		AccountNumber: []int64{0, 0},
		AccountPer:    0.0,
		AccountsPer:   0.0,
	}

	validateJsonRedaction(t, user, userRedacted)
}

// TestRedactSimpleXml tests redaction on a simple struct with xml
func TestRedactSimpleXml(t *testing.T) {

	// Simple struct
	type User struct {
		Username  string   `xml:"userName" `
		Password  string   `xml:"password" redact:""`
		DbSecrets []string `xml:"dbSecrets"`
	}

	user := &User{
		Username:  "Redaction",
		Password:  "codepassword",
		DbSecrets: []string{"db_secret_1", "db_secret_2"},
	}

	userRedacted := &User{
		Username:  "Redaction",
		Password:  "********",
		DbSecrets: []string{"db_secret_1", "db_secret_2"},
	}

	validateXmlRedaction(t, user, userRedacted)
}

// TestRedactNested tests redaction on a nested struct with json.
func TestRedactNested(t *testing.T) {

	// Simple struct
	type User struct {
		Username  string   `json:"userName" `
		Password  string   `json:"password" redact:""`
		DbSecrets []string `json:"dbSecrets" redact:""`
	}

	// Nested struct
	type Users struct {
		Secret   string   `json:"secret" redact:""`
		Keys     []string `json:"keys" redact:""`
		UserInfo []User   `json:"userInfo"`
	}

	users := &Users{
		Secret: "secret_sshhh",
		Keys:   []string{"key_1", "key_2", "key_3"},
		UserInfo: []User{
			{
				Username:  "Redaction Test",
				Password:  "Redaction_Password",
				DbSecrets: []string{"Redaction_db_secret_1", "Redaction_db_secret_2"},
			},
			{
				Username:  "Redaction Test 2",
				Password:  "Redaction_Password",
				DbSecrets: []string{"Redaction_db_secret_1", "Redaction_db_secret_2"},
			},
		},
	}

	usersRedacted := &Users{
		Secret: "********",
		Keys:   []string{"********", "********", "********"},
		UserInfo: []User{
			{
				Username:  "Redaction Test",
				Password:  "********",
				DbSecrets: []string{"********", "********"},
			},
			{
				Username:  "Redaction Test 2",
				Password:  "********",
				DbSecrets: []string{"********", "********"},
			},
		},
	}

	validateJsonRedaction(t, users, usersRedacted)
}

// TestRedactNestedXml tests redaction on a nested struct with xml.
func TestRedactNestedXml(t *testing.T) {

	// Simple struct
	type User struct {
		Username  string   `xml:"userName" `
		Password  string   `xml:"password" redact:""`
		DbSecrets []string `xml:"dbSecrets" redact:""`
	}

	// Nested struct
	type Users struct {
		Secret   string   `xml:"secret" redact:""`
		Keys     []string `xml:"keys" redact:""`
		UserInfo []User   `xml:"userInfo"`
	}

	users := &Users{
		Secret: "secret_sshhh",
		Keys:   []string{"key_1", "key_2", "key_3"},
		UserInfo: []User{
			{
				Username:  "Redaction Test",
				Password:  "Redaction_Password",
				DbSecrets: []string{"Redaction_db_secret_1", "Redaction_db_secret_2"},
			},
			{
				Username:  "Redaction Test 2",
				Password:  "Redaction_Password",
				DbSecrets: []string{"Redaction_db_secret_1", "Redaction_db_secret_2"},
			},
		},
	}

	usersRedacted := &Users{
		Secret: "********",
		Keys:   []string{"********", "********", "********"},
		UserInfo: []User{
			{
				Username:  "Redaction Test",
				Password:  "********",
				DbSecrets: []string{"********", "********"},
			},
			{
				Username:  "Redaction Test 2",
				Password:  "********",
				DbSecrets: []string{"********", "********"},
			},
		},
	}

	validateXmlRedaction(t, users, usersRedacted)
}

// TestRedactEmptyPointer tests redacting on a empty pointer.
func TestRedactEmptyFields(t *testing.T) {

	// Simple struct
	type User struct {
		Username  string   `json:"userName" redact:""`
		Password  string   `json:"password" redact:""`
		DbSecrets []string `json:"dbSecrets" redact:""`
	}

	user := &User{
		Username:  "",
		Password:  "redactionpassword",
		DbSecrets: []string{},
	}

	userRedacted := &User{
		Username:  "",
		Password:  "********",
		DbSecrets: []string{},
	}

	// Validate input with empty fields
	validateJsonRedaction(t, user, userRedacted)
}

func TestRedactEmptyPointer(t *testing.T) {

	// Simple struct
	type User struct {
		Username  string   `json:"userName" redact:""`
		Password  string   `json:"password" redact:""`
		DbSecrets []string `json:"dbSecrets" redact:""`
	}

	// Validate empty pointer input
	var userEmpty *User
	validateJsonRedaction(t, userEmpty, userEmpty)
}

// TestRedactNil tests redacting on a empty or nil input with json tags.
func TestRedactNil(t *testing.T) {

	t.Helper()

	got := Redact(nil)

	assert.Equal(t, "", got,
		"JSON representation mismatch after redacting fields")
}

// TestRedactNilXml tests redacting on a empty or nil input with xml tags.
func TestRedactNilXml(t *testing.T) {

	// Simple struct
	type User struct {
		Username  string   `xml:"userName" redact:""`
		Password  string   `xml:"password" redact:""`
		DbSecrets []string `xml:"dbSecrets" redact:""`
	}

	// Error struct
	type Redact struct {
		Error string `xml:"error" redact:""`
	}

	testError := Redact{
		Error: "Empty request",
	}

	user := &User{
		Username:  "",
		Password:  "redactionpassword",
		DbSecrets: []string{},
	}

	userRedacted := &User{
		Username:  "",
		Password:  "********",
		DbSecrets: []string{},
	}

	// Validate input with empty fields
	validateXmlRedaction(t, user, userRedacted)

	//  Validate nil input
	validateXmlRedaction(t, nil, testError)
}

// TestRedactNestedNil tests redacting on a nested complex struct with
// some nil, empty and specified sensitive fields.
func TestRedactNestedNil(t *testing.T) {
	// Simple struct
	type User struct {
		Username  string   `json:"userName" `
		Password  string   `json:"password" redact:""`
		DbSecrets []string `json:"dbSecrets" redact:""`
	}

	// Nested struct
	type Users struct {
		Secret   string   `json:"secret" redact:""`
		Keys     []string `json:"keys" redact:""`
		UserInfo []User   `json:"userInfo"`
	}

	users := &Users{
		Secret: "",
		Keys:   nil,
		UserInfo: []User{
			{
				Username:  "Redaction 1",
				Password:  "",
				DbSecrets: []string{"Redaction_db_secret_1", "Redaction_db_secret_2"},
			},
			{
				Username:  "Redaction 2",
				Password:  "Redaction_Password",
				DbSecrets: []string{},
			},
		},
	}

	userScrubbed := &Users{
		Secret: "",
		Keys:   nil,
		UserInfo: []User{
			{
				Username:  "Redaction 1",
				Password:  "",
				DbSecrets: []string{"********", "********"},
			},
			{
				Username:  "Redaction 2",
				Password:  "********",
				DbSecrets: []string{},
			},
		},
	}

	validateJsonRedaction(t, users, userScrubbed)
}

// TestRedactNestedXmlNil tests redacting on a nested complex struct with
// some nil, empty and specified sensitive fields.
func TestRedactNestedXmlNil(t *testing.T) {
	// Simple struct
	type User struct {
		Username  string   `xml:"userName" `
		Password  string   `xml:"password" redact:""`
		DbSecrets []string `xml:"dbSecrets" redact:""`
	}

	// Nested struct
	type Users struct {
		Secret   string   `xml:"secret" redact:""`
		Keys     []string `xml:"keys" redact:""`
		UserInfo []User   `xml:"userInfo"`
	}

	users := &Users{
		Secret: "",
		Keys:   nil,
		UserInfo: []User{
			{
				Username:  "Redaction 1",
				Password:  "",
				DbSecrets: []string{"Redaction_db_secret_1", "Redaction_db_secret_2"},
			},
			{
				Username:  "Redaction 2",
				Password:  "Redaction_Password",
				DbSecrets: []string{},
			},
		},
	}

	userRedacted := &Users{
		Secret: "",
		Keys:   nil,
		UserInfo: []User{
			{
				Username:  "Redaction 1",
				Password:  "",
				DbSecrets: []string{"********", "********"},
			},
			{
				Username:  "Redaction 2",
				Password:  "********",
				DbSecrets: []string{},
			},
		},
	}

	validateXmlRedaction(t, users, userRedacted)
}

// validateJsonRedaction is a helper function to validate redaction functionality on a struct with json tags.
func validateJsonRedaction(t *testing.T, msg, redactedMsg interface{}) {
	t.Helper()

	// Get the redacted string from Redact API.
	got := Redact(msg)

	// Compare it against the given redacted representaation.
	var b []byte
	b, _ = json.Marshal(redactedMsg)
	want := string(b)

	assert.Equal(t, want, got,
		"JSON representation mismatch after redacting fields")
}

// validateXmlRedaction is a helper function to validate redaction functionality on a struct with xml tags.
func validateXmlRedaction(t *testing.T, msg, redactedMsg interface{}) {
	t.Helper()

	// Get the redacted string from Redact API.
	got := RedactXml(msg)

	// Compare it against the given redacted representaation.
	var b []byte
	b, _ = xml.Marshal(redactedMsg)
	want := string(b)

	assert.Equal(t, want, got,
		"XML representation mismatch after redacting fields")
}
