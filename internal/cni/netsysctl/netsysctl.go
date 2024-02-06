package netsysctl

import (
	"fmt"
	"strconv"

	"github.com/phoracek/passt-kubevirt-network-binding-plugin/internal/cni/sysctl"
)

type sysControl struct{}

var sysCtl = sysctl.New()

func New() sysControl {
	return sysControl{}
}

func (_ sysControl) IPv4SetPingGroupRange(from, to int) error {
	return sysCtl.SetSysctl(sysctl.PingGroupRange, fmt.Sprintf("%d %d", from, to))
}

func (_ sysControl) IPv4SetUnprivilegedPortStart(port int) error {
	return sysCtl.SetSysctl(sysctl.UnprivilegedPortStart, strconv.Itoa(port))
}
