package clientcompat

import (
	"testing"

	"trojan-panel/model/constant"
)

func TestHysteria2ShareServer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		portHopping string
		client      string
		wantServer  string
		wantQuery   string
	}{
		{
			name:       "without port hopping",
			wantServer: "example.com:29999",
		},
		{
			name:        "official URI keeps multi-port authority",
			portHopping: "20000-30000,40000",
			wantServer:  "example.com:20000-30000,40000",
		},
		{
			name:        "V2Ray uses parseable port and mport extension",
			portHopping: "20000-30000,40000",
			client:      constant.ClientV2Ray,
			wantServer:  "example.com:29999",
			wantQuery:   "&mport=20000-30000%2C40000",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			server, query := Hysteria2ShareServer("example.com", 29999, test.portHopping, test.client)
			if server != test.wantServer {
				t.Fatalf("Hysteria2ShareServer() server = %q, want %q", server, test.wantServer)
			}
			if query != test.wantQuery {
				t.Fatalf("Hysteria2ShareServer() query = %q, want %q", query, test.wantQuery)
			}
		})
	}
}
