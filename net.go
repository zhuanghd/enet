package enet

import (
	"os/exec"
	"strings"
)

// ResolveMetalBox 解析铁盒IP
func ResolveMetalBox() string {
	cmd := exec.Command("bash", "-c", "curl -s https://gitee.com/zhuanghd/config/raw/master/metalbox")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}
