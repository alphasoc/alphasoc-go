package alphasoc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetAlerts(follow string) (*Alerts, error) {
	return c.GetAlertsCtx(context.Background(), follow)
}

// GetAlertsCtx retrieves alerts from alphasoc API.
// Follow parameter is used to retrieve new alerts since last request.
// If response is different than 200 OK or response cannot be decoded
// alphasoc.Error is returned, which contains error Message and response StatusCode.
func (c *Client) GetAlertsCtx(ctx context.Context, follow string) (*Alerts, error) {
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
		alerts := &Alerts{}
		err := decoder.Decode(alerts)
		if err != nil {
			return nil, APIError{fmt.Sprintf("parsing alerts: %v", err), resp.StatusCode}
		}

		return alerts, nil

	case http.StatusBadRequest:
		errMsg := &ErrorMessage{}
		err := decoder.Decode(errMsg)
		if err != nil {
			return nil, APIError{fmt.Sprintf("status bad request, json decode error %v", err), resp.StatusCode}
		}

		if errMsg.Message == nil {
			return nil, APIError{"status bad request", resp.StatusCode}
		}

		return nil, APIError{*errMsg.Message, resp.StatusCode}

	case http.StatusUnauthorized:
		errMsg := &ErrorMessage{}
		err := decoder.Decode(errMsg)
		if err != nil {
			return nil, APIError{fmt.Sprintf("status unauthorized, json decode error %v", err), resp.StatusCode}
		}

		if errMsg.Message == nil {
			return nil, APIError{"status unauthorized", resp.StatusCode}
		}

		return nil, APIError{*errMsg.Message, resp.StatusCode}

	case http.StatusForbidden:
		errMsg := &ErrorMessage{}
		err := decoder.Decode(errMsg)
		if err != nil {
			return nil, APIError{fmt.Sprintf("status forbidden, json decode error %v", err), resp.StatusCode}
		}

		if errMsg.Message == nil {
			return nil, APIError{"status forbidden", resp.StatusCode}
		}

		return nil, APIError{*errMsg.Message, resp.StatusCode}

	case http.StatusTooManyRequests:
		errMsg := &ErrorMessage{}
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
