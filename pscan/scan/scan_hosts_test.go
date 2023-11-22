package scan_test

import (
	"net"
	"powerful-command-line/pscan/scan"
	"strconv"
	"testing"
)

func TestStateString(t *testing.T) {
	ps := scan.PortState{}
	if ps.Open.String() != "closed" {
		t.Errorf("Expected %q, got %q instead\n", "closed", ps.Open.String())
	}
	ps.Open = true
	if ps.Open.String() != "open" {
		t.Errorf("Expected %q, got %q instead\n", "open", ps.Open.String())
	}
}

func TestRunHostFound(t *testing.T) {
	testCases := []struct {
		name        string
		expectState string
	}{
		{"OpenPort", "open"},
		{"ClosedPort", "closed"},
	}
	host := "localhost"
	hl := &scan.HostsList{}
	hl.Add(host)

	ports := []int{}
	// init ports, 1 open, 1 closed
	for _, tc := range testCases {
		ln, err := net.Listen("tcp", net.JoinHostPort(host, "0"))
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		_, portStr, err := net.SplitHostPort(ln.Addr().String())
		if err != nil {
			t.Fatal(err)
		}
		port, err := strconv.Atoi(portStr)
		if err != nil {
			t.Fatal(err)
		}
		ports = append(ports, port)
		if tc.name == "ClosedPort" {
			ln.Close()
		}
	}
}
