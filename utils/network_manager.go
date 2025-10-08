package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	apiURL      = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"
	instruction = `You are a professional YouTuber scriptwriter.
Please improve the following script by fixing grammar, style, and flow.
Make it sound natural and conversational, like a charismatic YouTuber speaking on camera.
Add YouTube-style delivery notations such as [Pause], [Emotional], [Excited], [Whisper], or [Joke] where appropriate.
IMPORTANT: Output only the improved script content. Do NOT include any introductory phrases, explanations, or summaries.`
)

type Requester interface {
	SendRequest(prompt string) (string, error)
}

type NetworkManager struct {
	ApiKey string
	Client *http.Client
}

func NewNeworkManager(apiKey string, client *http.Client) (*NetworkManager, error) {
	if apiKey == "" || client == nil {
		return nil, errors.New("couldn't init network manager, check envs")
	}
	manager := NetworkManager{
		ApiKey: apiKey,
		Client: client,
	}
	return &manager, nil
}

func (n NetworkManager) SendRequest(prompt string) (string, error) {
	reqBody := createRequestBody(prompt)

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL+"?key="+n.ApiKey, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := n.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	var apiResp struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	err = json.Unmarshal(respBody, &apiResp)
	if err != nil {
		return "", fmt.Errorf("unmarshal response: %w", err)
	}

	if len(apiResp.Candidates) == 0 || len(apiResp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content in response")
	}

	var improvedScript string
	for _, part := range apiResp.Candidates[0].Content.Parts {
		improvedScript += part.Text
	}

	return improvedScript, nil
}

func createRequestBody(prompt string) map[string]any {
	return map[string]any{
		"contents": []map[string]any{
			{
				"parts": []map[string]string{
					{"text": instruction},
					{"text": prompt},
				},
			},
		},
	}
}
