package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/josephe44/chatgtp-api-golang/models"
)

const apiEndpoint = "https://api.openai.com/v1/chat/completions"
const temperature = 0.5
const aiModel = "gpt-3.5-turbo-0301"

func GetAIResponse(apiKey, escapedInput string) (*models.Response, error) {
	payload := models.Request{
		Model:       aiModel,
		Messages:    []models.Message{{Role: "user", Content: escapedInput}},
		Temperature: temperature,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %s", err.Error())
	}

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %s", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var responseObj models.Response
	err = json.NewDecoder(resp.Body).Decode(&responseObj)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err.Error())
	}

	return &responseObj, nil
}
