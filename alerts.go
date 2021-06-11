package alphasoc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alphasoc/alphasoc-go/models"
)

// Alerts retrieves alerts from alphasoc API.
// Follow parameter is used to retrieve new alerts since last request.
// If response is different than 200 OK or response cannot be decoded
// alphasoc.Error is returned, which contains error Message and response StatusCode.
func (c *Client) Alerts(ctx context.Context, follow string) (*models.Alerts, error) {
	req, err := c.prepareRequest(ctx, "GET", alertsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	if follow != "" {
		query := req.URL.Query()
		query.Add("follow", follow)
		req.URL.RawQuery = query.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	switch resp.StatusCode {
	case http.StatusOK:
		alerts := &models.Alerts{}
		err := decoder.Decode(alerts)
		if err != nil {
			return nil, APIError{fmt.Sprintf("parsing alerts: %v", err), resp.StatusCode}
		}

		return alerts, nil

	case http.StatusBadRequest:
		errMsg := &models.ErrorMessage{}
		err := decoder.Decode(errMsg)
		if err != nil {
			return nil, APIError{fmt.Sprintf("status bad request, json decode error %v", err), resp.StatusCode}
		}

		if errMsg.Message == nil {
			return nil, APIError{"status bad request", resp.StatusCode}
		}

		return nil, APIError{*errMsg.Message, resp.StatusCode}

	case http.StatusUnauthorized:
		errMsg := &models.ErrorMessage{}
		err := decoder.Decode(errMsg)
		if err != nil {
			return nil, APIError{fmt.Sprintf("status unauthorized, json decode error %v", err), resp.StatusCode}
		}

		if errMsg.Message == nil {
			return nil, APIError{"status unauthorized", resp.StatusCode}
		}

		return nil, APIError{*errMsg.Message, resp.StatusCode}

	case http.StatusForbidden:
		errMsg := &models.ErrorMessage{}
		err := decoder.Decode(errMsg)
		if err != nil {
			return nil, APIError{fmt.Sprintf("status forbidden, json decode error %v", err), resp.StatusCode}
		}

		if errMsg.Message == nil {
			return nil, APIError{"status forbidden", resp.StatusCode}
		}

		return nil, APIError{*errMsg.Message, resp.StatusCode}

	case http.StatusTooManyRequests:
		errMsg := &models.ErrorMessage{}
		err := decoder.Decode(errMsg)
		if err != nil {
			return nil, APIError{fmt.Sprintf("status too many requests, json decode error %v", err), resp.StatusCode}
		}

		if errMsg.Message == nil {
			return nil, APIError{"status too many requests", resp.StatusCode}
		}

		return nil, APIError{*errMsg.Message, resp.StatusCode}

	default:
		return nil, APIError{"unexpected response", resp.StatusCode}
	}
}
