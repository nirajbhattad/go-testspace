package gochallenges

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Mask interface {
	MaskJson(req []byte) (res []byte, err error)
}

type TestMask struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Office  string `json:"office"`
}

func MaskingRedaction() {
	testMask := TestMask{
		Name:    "Niraj",
		Phone:   "8660941082",
		Address: "29 Raglan Avenue",
		Office:  "BookMyshow",
	}
	testBuf, _ := json.Marshal(testMask)

	masking := Masking{}
	res, _ := masking.MaskJson(testBuf)
	fmt.Println(string(res))
}

type Masking struct{}

func (M *Masking) MaskJson(req []byte) (res []byte, err error) {

	// Mask Request
	byteRes := RedactJSON(req)

	// Print Masked Response
	fmt.Println(string(byteRes))
	return byteRes, nil
}

// RedactJSON matches keys in JSON and clears their string values.
func RedactJSON(b []byte) []byte {
	for _, rule := range redactKeyRules {
		b = rule.ReplaceAll(b, []byte(`"$1":"*"`))
	}
	return b
}

const redactValue = `".*?[^\\]"`

var redactKeyRules = []*regexp.Regexp{
	regexp.MustCompile(`(?i)"([^"]*?tax[^"]*?)"\s*:\s*` + redactValue),
	regexp.MustCompile(`(?i)"([^"]*?phone[^"]*?)"\s*:\s*` + redactValue),
	regexp.MustCompile(`(?i)"([^"]*?address[^"]*?)"\s*:\s*` + redactValue),
	regexp.MustCompile(`(?i)"([^"]*?office[^"]*?)"\s*:\s*` + redactValue),
}
