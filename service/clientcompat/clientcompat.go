package clientcompat

import "trojan-panel/model/constant"

// SupportsV2Ray reports whether a node type can be exported to V2Ray-style
// base64 subscriptions. NaiveProxy is intentionally limited to sing-box.
func SupportsV2Ray(nodeTypeId *uint) bool {
	return nodeTypeId != nil && *nodeTypeId != constant.NaiveProxy
}
