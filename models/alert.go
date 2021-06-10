package models

import (
	"encoding/json"
	"errors"
)

// DNS returns DnsEvent from Alert if EventType is "dns"
func (a *Alert) DNS() (*DnsEvent, error) {
	if a.EventType == nil || *a.EventType != "dns" {
		return nil, errors.New("alert EventType different than dns")
	}

	dns := &DnsEvent{}

	err := json.Unmarshal(a.Event, dns)
	if err != nil {
		return nil, err
	}

	return dns, nil
}

// IP returns IpEvent from Alert if EventType is "ip"
func (a *Alert) IP() (*IpEvent, error) {
	if a.EventType == nil || *a.EventType != "ip" {
		return nil, errors.New("alert EventType different than ip")
	}

	ip := &IpEvent{}

	err := json.Unmarshal(a.Event, ip)
	if err != nil {
		return nil, err
	}

	return ip, nil
}

// HTTP returns HttpEvent from Alert if EventType is "http"
func (a *Alert) HTTP() (*HttpEvent, error) {
	if a.EventType == nil || *a.EventType != "http" {
		return nil, errors.New("alert EventType different than http")
	}

	http := &HttpEvent{}

	err := json.Unmarshal(a.Event, http)
	if err != nil {
		return nil, err
	}

	return http, nil
}

// TLS returns TlsEvent from Alert if EventType is "tls"
func (a *Alert) TLS() (*TlsEvent, error) {
	if a.EventType == nil || *a.EventType != "tls" {
		return nil, errors.New("alert EventType different than tls")
	}

	tls := &TlsEvent{}

	err := json.Unmarshal(a.Event, tls)
	if err != nil {
		return nil, err
	}

	return tls, nil
}
