package openapi

import "encoding/json"

// EventsToBody converts DNS events to /events/dns request body
func (r ApiV1EventsDnsPostRequest) EventsToBody() ([]byte, error) {
	events := []byte{}
	for _, ev := range r.dnsEvents {
		j, err := json.Marshal(ev)
		if err != nil {
			return nil, err
		}

		events = append(events, j...)
	}

	return events, nil
}

// EventsToBody converts HTTP events to /events/http request body
func (r ApiV1EventsHttpPostRequest) EventsToBody() ([]byte, error) {
	events := []byte{}
	for _, ev := range r.httpEvents {
		j, err := json.Marshal(ev)
		if err != nil {
			return nil, err
		}

		events = append(events, j...)
	}

	return events, nil
}

// EventsToBody converts IP events to /events/ip request body
func (r ApiV1EventsIpPostRequest) EventsToBody() ([]byte, error) {
	events := []byte{}
	for _, ev := range r.ipEvents {
		j, err := json.Marshal(ev)
		if err != nil {
			return nil, err
		}

		events = append(events, j...)
	}

	return events, nil
}

// EventsToBody converts Lease events to /events/lease request body
func (r ApiV1EventsLeasePostRequest) EventsToBody() ([]byte, error) {
	events := []byte{}
	for _, ev := range r.leaseEvents {
		j, err := json.Marshal(ev)
		if err != nil {
			return nil, err
		}

		events = append(events, j...)
	}

	return events, nil
}

// EventsToBody converts TLS events to /events/tls request body
func (r ApiV1EventsTlsPostRequest) EventsToBody() ([]byte, error) {
	events := []byte{}
	for _, ev := range r.tlsEvents {
		j, err := json.Marshal(ev)
		if err != nil {
			return nil, err
		}

		events = append(events, j...)
	}

	return events, nil
}
