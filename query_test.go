package sampquery

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetServerInfo(t *testing.T) {
	type args struct {
		host          string
		attemptDecode bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{"valid", args{"server.ls-rp.com:7777", false}, ""},
		{"valid", args{"46.174.54.184:7777", false}, ""},
		{"invalid", args{"18.251.83.150:80", false}, "socket read timed out"},
		{"invalid", args{"not a valid url", false}, "failed to resolve host: address not a valid url: missing port in address"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			server, err := GetServerInfo(ctx, tt.args.host, tt.args.attemptDecode)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.NotEmpty(t, server.Address, "server.Address")
				assert.NotEmpty(t, server.Hostname, "server.Hostname")
				assert.NotEmpty(t, server.Gamemode, "server.Gamemode")
				assert.NotZero(t, server.MaxPlayers, "server.MaxPlayers")
			}
			time.Sleep(time.Second) // allow goroutines to run so socket timeout doesn't fire
		})
	}
}

func TestQuery_GetPing(t *testing.T) {
	tests := []struct {
		addr string
	}{
		{"server.ls-rp.com:7777"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			query, err := NewQuery(tt.addr)
			if err != nil {
				t.Error(err)
			}
			gotPing, err := query.GetPing(context.Background())
			assert.NotZero(t, gotPing)
			fmt.Print(gotPing)
		})
	}
}
