package gochallenges

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ChatGptResponse struct {
	Answers []string `json:"answers"`
}

func ChatGptEndpoint() {
	http.HandleFunc("/chatgpt", chatGptHandler)
	http.ListenAndServe(":8080", nil)
}

func chatGptHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Missing 'query' parameter in request", http.StatusBadRequest)
		return
	}

	chatGptResponse, err := GetChatGptResponse(query)
	if err != nil {
		http.Error(w, "Failed to get response from Chat GPT API: "+err.Error(), http.StatusInternalServerError)
		return
	}

	answers := strings.Join(chatGptResponse.Answers, ", ")
	fmt.Fprintf(w, "Answers: %s", answers)
}

func GetChatGptResponse(query string) (ChatGptResponse, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.chat-gpt.com/answers", nil)
	q := req.URL.Query()
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return ChatGptResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ChatGptResponse{}, err
	}

	var chatGptResponse ChatGptResponse
	err = json.Unmarshal(body, &chatGptResponse)
	if err != nil {
		return ChatGptResponse{}, err
	}

	return chatGptResponse, nil
}
