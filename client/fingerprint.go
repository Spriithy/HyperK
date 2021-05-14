package client

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

var SysInfo *systemInfo

type systemInfo struct {
	OS       string
	Version  string
	HostName string
	UserName *user.User
	CpuCount int
}

func init() {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd := exec.Command("cmd", "ver")
	cmd.Stdin = strings.NewReader("")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	osStr := strings.Replace(stdout.String(), "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)

	tmp1 := strings.Index(osStr, "[Version")
	tmp2 := strings.Index(osStr, "]")

	var ver string
	if tmp1 == -1 || tmp2 == -1 {
		ver = "unknown"
	} else {
		ver = osStr[tmp1+9 : tmp2]
	}

	SysInfo = &systemInfo{
		OS:       runtime.GOOS,
		Version:  ver,
		CpuCount: runtime.NumCPU(),
	}

	SysInfo.HostName, _ = os.Hostname()
	SysInfo.UserName, _ = user.Current()
}

func (sysInfo *systemInfo) String() string {
	return fmt.Sprintf("OS=%v|Version=%v|Hostname=%v|CPUs=%v|Username=%v", sysInfo.OS, sysInfo.Version, sysInfo.HostName, sysInfo.CpuCount, sysInfo.UserName.Username)
}
