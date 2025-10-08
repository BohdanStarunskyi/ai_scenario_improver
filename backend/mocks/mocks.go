package mocks

import "errors"

type NetworkManagerSuccessMock struct {
	SendRequestFunc func(prompt string) (string, error)
}

func (n NetworkManagerSuccessMock) SendRequest(prompt string) (string, error) {
	return "Test Response", nil
}

type NetworkManagerErrorMock struct {
	SendRequestFunc func(prompt string) (string, error)
}

func (n NetworkManagerErrorMock) SendRequest(prompt string) (string, error) {
	return "", errors.New("Test Error")
}
