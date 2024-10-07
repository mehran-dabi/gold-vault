package clients

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"goldvault/user-service/internal/core/domain/entity"
)

// KavenegarSMSProviderClient handles communication with KavenegarSMSProviderClient API.
type KavenegarSMSProviderClient struct {
	APIKey string
	Sender string
	Host   string
	net    *http.Client
}

// NewKavenegarSMSProvider creates a new instance of KavenegarSMSProviderClient.
func NewKavenegarSMSProvider(apiKey, sender, host string) *KavenegarSMSProviderClient {
	return &KavenegarSMSProviderClient{
		APIKey: apiKey,
		Sender: sender,
		Host:   host,
		net:    &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *KavenegarSMSProviderClient) SendSMS(ctx context.Context, sms entity.SimpleSMS) error {
	// Set the sender if not provided
	if sms.Sender == "" {
		sms.Sender = c.Sender
	}

	if sms.Template == "" {
		sms.Template = "default"
	}
	// Convert the struct to URL query parameters
	queryParams, err := StructToQueryParams(sms)
	if err != nil {
		return err
	}

	request := APIRequest{
		URL:    fmt.Sprintf("%s/%s/verify/lookup.json", c.Host, c.APIKey),
		Method: http.MethodGet,
		Params: queryParams,
	}

	_, err = c.callAPI(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

// StructToQueryParams converts a struct into URL query parameters
func StructToQueryParams(data interface{}) (map[string]string, error) {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	// Ensure the passed interface is a struct
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %s", t.Kind())
	}

	params := make(map[string]string)

	// Iterate through the struct fields
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		// Use the field name as the query parameter key
		key := strings.ToLower(fieldType.Name)

		// Handle different types, converting to string appropriately
		switch fieldValue.Kind() {
		case reflect.String:
			params[key] = fieldValue.String()
		case reflect.Int, reflect.Int64:
			params[key] = fmt.Sprintf("%d", fieldValue.Int())
		default:
			// Handle any other cases (add more as needed)
			params[key] = fmt.Sprintf("%v", fieldValue)
		}
	}

	// Return the encoded query string
	return params, nil
}

// APIRequest is the structure to define API request properties
type APIRequest struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    []byte
	Params  map[string]string // Added for URL parameters
}

// callAPI performs a generic API call.
func (c *KavenegarSMSProviderClient) callAPI(ctx context.Context, request APIRequest) (string, error) {

	// Parse the base URL
	parsedURL, err := url.Parse(request.URL)
	if err != nil {
		return "", err
	}

	// Add URL parameters if provided
	if request.Method == http.MethodGet && len(request.Params) > 0 {
		query := parsedURL.Query()
		for key, value := range request.Params {
			query.Add(key, value)
		}
		parsedURL.RawQuery = query.Encode()
	}

	// Create a new HTTP request
	req, err := http.NewRequest(request.Method, parsedURL.String(), bytes.NewBuffer(request.Body))
	if err != nil {
		return "", err
	}

	// Set headers if provided
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	// Create a new HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check if the status code indicates success (2xx)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", errors.New("API call failed with status: " + resp.Status)
	}

	return string(body), nil
}
