package alphasoc

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func stringPtr(s string) *string {
	return &s
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func int32Ptr(i int32) *int32 {
	return &i
}

func int64Ptr(i int64) *int64 {
	return &i
}

func float64Ptr(f float64) *float64 {
	return &f
}

func TestDNS(t *testing.T) {
	type testCase struct {
		alert            *Alert
		expected         *DnsEvent
		expectedErrorMsg string
	}

	cases := []testCase{
		{&Alert{EventType: nil}, nil, "alert EventType different than dns"},
		{&Alert{EventType: stringPtr("ip")}, nil, "alert EventType different than dns"},
		{&Alert{EventType: stringPtr("dns"), Event: []byte(`{`)}, nil, "unexpected end of JSON input"},
		{
			&Alert{EventType: stringPtr("dns"), Event: []byte(`{"ts":"2021-06-01T08:06:03Z","srcIP":"10.14.1.39","srcHost":"win-3xchk5-lp","srcMac":"da:23:68:50:c4:77","query":"jymfezgszvqsa.net","qtype":"A"}`)},
			&DnsEvent{Ts: timePtr(time.Date(2021, time.June, 1, 8, 6, 3, 0, time.UTC)), SrcIP: stringPtr("10.14.1.39"), SrcHost: stringPtr("win-3xchk5-lp"), SrcMac: stringPtr("da:23:68:50:c4:77"), Query: stringPtr("jymfezgszvqsa.net"), Qtype: stringPtr("A")},
			""},
	}

	for _, test := range cases {
		dns, err := test.alert.DNS()
		if err != nil {
			if test.expectedErrorMsg == "" {
				t.Fatalf("Error unexpected, got %v", err)
			}

			if !strings.Contains(err.Error(), test.expectedErrorMsg) {
				t.Fatalf("Expected error to contain: %v, got %v", test.expectedErrorMsg, err.Error())
			}
		} else {
			if test.expectedErrorMsg != "" {
				t.Fatal("Expected error, got nil")
			}

			if !reflect.DeepEqual(test.expected, dns) {
				t.Fatalf("Expected dns event: %+v, got: %+v, err: %v", test.expected, dns, err)
			}
		}
	}
}

func TestIP(t *testing.T) {
	type testCase struct {
		alert            *Alert
		expected         *IpEvent
		expectedErrorMsg string
	}

	cases := []testCase{
		{&Alert{EventType: nil}, nil, "alert EventType different than ip"},
		{&Alert{EventType: stringPtr("dns")}, nil, "alert EventType different than ip"},
		{&Alert{EventType: stringPtr("ip"), Event: []byte(`{`)}, nil, "unexpected end of JSON input"},
		{
			&Alert{EventType: stringPtr("ip"), Event: []byte(`{"ts":"2021-06-01T10:09:44Z","srcIP":"10.100.90.3","srcPort":64300,"srcHost":"win-z2yyxq-dp","srcMac":"16:c8:60:26:09:a6","destIP":"46.126.206.128","destPort":16879,"proto":"udp","bytesIn":19346,"bytesOut":10540,"app":"ssl3","action":"allowed","duration":0.681071712857}`)},
			&IpEvent{Ts: timePtr(time.Date(2021, time.June, 1, 10, 9, 44, 0, time.UTC)), SrcIP: stringPtr("10.100.90.3"), SrcPort: int32Ptr(64300), SrcHost: stringPtr("win-z2yyxq-dp"), SrcMac: stringPtr("16:c8:60:26:09:a6"), DestIP: stringPtr("46.126.206.128"), DestPort: int32Ptr(16879), Proto: stringPtr("udp"), BytesIn: int64Ptr(19346), BytesOut: int64Ptr(10540), App: stringPtr("ssl3"), Action: stringPtr("allowed"), Duration: float64Ptr(0.681071712857)},
			""},
	}

	for _, test := range cases {
		ip, err := test.alert.IP()

		if err != nil {
			if test.expectedErrorMsg == "" {
				t.Fatalf("Error unexpected, got %v", err)
			}

			if !strings.Contains(err.Error(), test.expectedErrorMsg) {
				t.Fatalf("Expected error to contain: %v, got %v", test.expectedErrorMsg, err.Error())
			}
		} else {
			if test.expectedErrorMsg != "" {
				t.Fatal("Expected error, got nil")
			}

			if !reflect.DeepEqual(test.expected, ip) {
				t.Fatalf("Expected ip event: %+v, got: %+v, err: %v", test.expected, ip, err)
			}
		}
	}
}

func TestHTTP(t *testing.T) {
	type testCase struct {
		alert            *Alert
		expected         *HttpEvent
		expectedErrorMsg string
	}

	cases := []testCase{
		{&Alert{EventType: nil}, nil, "alert EventType different than http"},
		{&Alert{EventType: stringPtr("dns")}, nil, "alert EventType different than http"},
		{&Alert{EventType: stringPtr("http"), Event: []byte(`{`)}, nil, "unexpected end of JSON input"},
		{
			&Alert{EventType: stringPtr("http"), Event: []byte(`{"ts":"2021-06-01T10:09:44Z","srcIP":"10.100.90.3","srcPort":64300,"srcHost":"win-z2yyxq-dp","srcMac":"16:c8:60:26:09:a6","url": "https://www.example.com/path/to/something","bytesIn":19346,"bytesOut":10540,"app":"ssl3","action":"allowed", "method": "GET", "status": 200, "userAgent": "Firefox", "contentType": "json"}`)},
			&HttpEvent{Ts: timePtr(time.Date(2021, time.June, 1, 10, 9, 44, 0, time.UTC)), SrcIP: stringPtr("10.100.90.3"), SrcPort: int32Ptr(64300), SrcHost: stringPtr("win-z2yyxq-dp"), SrcMac: stringPtr("16:c8:60:26:09:a6"), Url: stringPtr("https://www.example.com/path/to/something"), BytesIn: int64Ptr(19346), BytesOut: int64Ptr(10540), App: stringPtr("ssl3"), Action: stringPtr("allowed"), Method: stringPtr("GET"), Status: int64Ptr(200), UserAgent: stringPtr("Firefox"), ContentType: stringPtr("json")},
			""},
	}

	for _, test := range cases {
		ip, err := test.alert.HTTP()

		if err != nil {
			if test.expectedErrorMsg == "" {
				t.Fatalf("Error unexpected, got %v", err)
			}

			if !strings.Contains(err.Error(), test.expectedErrorMsg) {
				t.Fatalf("Expected error to contain: %v, got %v", test.expectedErrorMsg, err.Error())
			}
		} else {
			if test.expectedErrorMsg != "" {
				t.Fatal("Expected error, got nil")
			}

			if !reflect.DeepEqual(test.expected, ip) {
				t.Fatalf("Expected http event: %+v, got: %+v, err: %v", test.expected, ip, err)
			}
		}
	}
}

func TestTLS(t *testing.T) {
	type testCase struct {
		alert            *Alert
		expected         *TlsEvent
		expectedErrorMsg string
	}

	cases := []testCase{
		{&Alert{EventType: nil}, nil, "alert EventType different than tls"},
		{&Alert{EventType: stringPtr("dns")}, nil, "alert EventType different than tls"},
		{&Alert{EventType: stringPtr("tls"), Event: []byte(`{`)}, nil, "unexpected end of JSON input"},
		{
			&Alert{EventType: stringPtr("tls"), Event: []byte(`{"ts":"2021-06-01T10:09:44Z","srcIP":"10.100.90.3","srcPort":64300,"srcHost":"win-z2yyxq-dp","srcMac":"16:c8:60:26:09:a6","destIP":"46.126.206.128","destPort":16879, "certHash": "abcdef1234567890", "issuer": "issuer", "subject": "subject", "ja3": "ja3hash", "ja3s": "ja3s", "validFrom": "2021-06-01T10:09:44Z", "validTo": "2021-06-01T10:09:44Z"}`)},
			&TlsEvent{Ts: timePtr(time.Date(2021, time.June, 1, 10, 9, 44, 0, time.UTC)), SrcIP: stringPtr("10.100.90.3"), SrcPort: int32Ptr(64300), SrcHost: stringPtr("win-z2yyxq-dp"), SrcMac: stringPtr("16:c8:60:26:09:a6"), DestIP: stringPtr("46.126.206.128"), DestPort: int32Ptr(16879), CertHash: stringPtr("abcdef1234567890"), Issuer: stringPtr("issuer"), Subject: stringPtr("subject"), Ja3: stringPtr("ja3hash"), Ja3s: stringPtr("ja3s"), ValidFrom: timePtr(time.Date(2021, time.June, 1, 10, 9, 44, 0, time.UTC)), ValidTo: timePtr(time.Date(2021, time.June, 1, 10, 9, 44, 0, time.UTC))},
			""},
	}

	for _, test := range cases {
		ip, err := test.alert.TLS()

		if err != nil {
			if test.expectedErrorMsg == "" {
				t.Fatalf("Error unexpected, got %v", err)
			}

			if !strings.Contains(err.Error(), test.expectedErrorMsg) {
				t.Fatalf("Expected error to contain: %v, got %v", test.expectedErrorMsg, err.Error())
			}
		} else {
			if test.expectedErrorMsg != "" {
				t.Fatal("Expected error, got nil")
			}

			if !reflect.DeepEqual(test.expected, ip) {
				t.Fatalf("Expected http event: %+v, got: %+v, err: %+v", test.expected, ip, err)
			}
		}
	}
}
