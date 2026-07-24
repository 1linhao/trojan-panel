package clientcompat

import (
	"encoding/base64"
	"encoding/json"
)

const (
	v2RayNConfigTypeHysteria2 = 7
	v2RayNConfigVersion       = 4
)

type V2RayNHysteria2Config struct {
	Remarks        string
	Address        string
	Port           uint
	Password       string
	SNI            string
	AllowInsecure  bool
	SalamanderPass string
	UpMbps         int
	DownMbps       int
	Ports          string
	HopInterval    string
}

type v2RayNProtocolExtra struct {
	SalamanderPass string `json:"SalamanderPass,omitempty"`
	UpMbps         int    `json:"UpMbps"`
	DownMbps       int    `json:"DownMbps"`
	Ports          string `json:"Ports,omitempty"`
	HopInterval    string `json:"HopInterval,omitempty"`
}

type v2RayNHysteria2Profile struct {
	ConfigType     int                 `json:"ConfigType"`
	ConfigVersion  int                 `json:"ConfigVersion"`
	Remarks        string              `json:"Remarks"`
	Address        string              `json:"Address"`
	Port           uint                `json:"Port"`
	Password       string              `json:"Password"`
	StreamSecurity string              `json:"StreamSecurity"`
	AllowInsecure  string              `json:"AllowInsecure"`
	SNI            string              `json:"Sni,omitempty"`
	ProtoExtra     v2RayNProtocolExtra `json:"ProtoExtraObj"`
}

// V2RayNHysteria2URI uses v2rayN's inner URI format because its standard
// hysteria2:// parser intentionally omits per-node bandwidth and hop interval.
func V2RayNHysteria2URI(config V2RayNHysteria2Config) (string, error) {
	allowInsecure := "false"
	if config.AllowInsecure {
		allowInsecure = "true"
	}
	profile := v2RayNHysteria2Profile{
		ConfigType:     v2RayNConfigTypeHysteria2,
		ConfigVersion:  v2RayNConfigVersion,
		Remarks:        config.Remarks,
		Address:        config.Address,
		Port:           config.Port,
		Password:       config.Password,
		StreamSecurity: "tls",
		AllowInsecure:  allowInsecure,
		SNI:            config.SNI,
		ProtoExtra: v2RayNProtocolExtra{
			SalamanderPass: config.SalamanderPass,
			UpMbps:         config.UpMbps,
			DownMbps:       config.DownMbps,
			Ports:          config.Ports,
			HopInterval:    config.HopInterval,
		},
	}
	content, err := json.Marshal(profile)
	if err != nil {
		return "", err
	}
	return "v2rayn://hysteria2/" + base64.RawURLEncoding.EncodeToString(content), nil
}
