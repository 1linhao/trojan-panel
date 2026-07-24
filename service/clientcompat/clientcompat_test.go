package clientcompat

import (
	"testing"
	"trojan-panel/model/constant"
)

func TestSupportsV2Ray(t *testing.T) {
	tests := []struct {
		name       string
		nodeTypeId *uint
		want       bool
	}{
		{name: "missing node type", nodeTypeId: nil, want: false},
		{name: "xray", nodeTypeId: uintPointer(constant.Xray), want: true},
		{name: "trojan-go", nodeTypeId: uintPointer(constant.TrojanGo), want: true},
		{name: "hysteria", nodeTypeId: uintPointer(constant.Hysteria), want: true},
		{name: "hysteria2", nodeTypeId: uintPointer(constant.Hysteria2), want: true},
		{name: "naiveproxy", nodeTypeId: uintPointer(constant.NaiveProxy), want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SupportsV2Ray(tt.nodeTypeId); got != tt.want {
				t.Fatalf("SupportsV2Ray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func uintPointer(value uint) *uint {
	return &value
}
