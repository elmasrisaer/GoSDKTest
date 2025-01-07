package hooks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
}

type Response struct {
	Status  int
	Headers map[string]string
	Body    string
}

type ErrorResponse struct {
	Error string
}

type CustomHook struct {
	currentToken  string
	currentExpiry int64
	mutex         sync.Mutex
}

func NewCustomHook() *CustomHook {
	return &CustomHook{}
}

func (h *CustomHook) getToken(clientID, clientSecret string) (string, int64, error) {
	fullURL := "https://auth.celitech.net/oauth2/token"

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	data := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=client_credentials", clientID, clientSecret)

	req, err := http.NewRequest("POST", fullURL, bytes.NewBufferString(data))
	if err != nil {
		return "", 0, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", 0, errors.New("failed to get token")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return "", 0, err
	}

	if errMsg, exists := response["error"].(string); exists {
		return "", 0, errors.New(errMsg)
	}

	expiresIn, ok := response["expires_in"].(float64)
	if !ok {
		return "", 0, errors.New("invalid expires_in value")
	}

	accessToken, ok := response["access_token"].(string)
	if !ok {
		return "", 0, errors.New("invalid access_token value")
	}

	return accessToken, int64(expiresIn * 1000), nil
}

func (h *CustomHook) BeforeRequest(req *Request, params map[string]string) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	clientID, ok1 := params["client_id"]
	clientSecret, ok2 := params["client_secret"]
	if !ok1 || !ok2 {
		return errors.New("missing client_id or client_secret")
	}

	if h.currentToken == "" || h.currentExpiry < time.Now().UnixMilli() {
		token, expiresIn, err := h.getToken(clientID, clientSecret)
		if err != nil {
			return err
		}
		h.currentToken = token
		h.currentExpiry = time.Now().UnixMilli() + expiresIn
	}

	req.Headers["Authorization"] = fmt.Sprintf("Bearer %s", h.currentToken)
	return nil
}

func (h *CustomHook) AfterResponse(req *Request, resp *Response, params map[string]string) {
	fmt.Printf("AfterResponse: %#v\n", resp)
}

func (h *CustomHook) OnError(req *Request, resp *ErrorResponse, params map[string]string) {
	fmt.Printf("On Error: %#v\n", resp)
}
