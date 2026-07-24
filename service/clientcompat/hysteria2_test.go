package clientcompat

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
)

func TestV2RayNHysteria2URI(t *testing.T) {
	t.Parallel()

	uri, err := V2RayNHysteria2URI(V2RayNHysteria2Config{
		Remarks:        "hy2-test",
		Address:        "example.com",
		Port:           29999,
		Password:       "test-password",
		SNI:            "sni.example.com",
		AllowInsecure:  true,
		SalamanderPass: "obfs-password",
		UpMbps:         30,
		DownMbps:       50,
		Ports:          "20000-30000",
		HopInterval:    "30",
	})
	if err != nil {
		t.Fatalf("V2RayNHysteria2URI() error = %v", err)
	}
	const prefix = "v2rayn://hysteria2/"
	if !strings.HasPrefix(uri, prefix) {
		t.Fatalf("V2RayNHysteria2URI() = %q, want prefix %q", uri, prefix)
	}
	content, err := base64.RawURLEncoding.DecodeString(strings.TrimPrefix(uri, prefix))
	if err != nil {
		t.Fatalf("decode inner URI: %v", err)
	}
	var profile v2RayNHysteria2Profile
	if err := json.Unmarshal(content, &profile); err != nil {
		t.Fatalf("unmarshal inner profile: %v", err)
	}
	if profile.ConfigType != v2RayNConfigTypeHysteria2 || profile.ConfigVersion != v2RayNConfigVersion {
		t.Fatalf("profile version/type = %d/%d", profile.ConfigVersion, profile.ConfigType)
	}
	if profile.Address != "example.com" || profile.Port != 29999 || profile.Password != "test-password" {
		t.Fatalf("profile endpoint/auth was not preserved: %+v", profile)
	}
	if profile.AllowInsecure != "true" || profile.SNI != "sni.example.com" {
		t.Fatalf("profile TLS settings were not preserved: %+v", profile)
	}
	if profile.ProtoExtra.UpMbps != 30 || profile.ProtoExtra.DownMbps != 50 {
		t.Fatalf("profile bandwidth = %d/%d", profile.ProtoExtra.UpMbps, profile.ProtoExtra.DownMbps)
	}
	if profile.ProtoExtra.Ports != "20000-30000" || profile.ProtoExtra.HopInterval != "30" {
		t.Fatalf("profile port hopping was not preserved: %+v", profile.ProtoExtra)
	}
}
