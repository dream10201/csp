package cmd

import (
	"bytes"
	"github.com/hirochachacha/go-smb2"
	"log"
	"net"
	"os/exec"
	"strings"
)

const (
	listAll = "pdbedit -L"
	server  = "127.0.0.1:445"
)

// Check 检查用户是否存在
func Check(user string) bool {
	if len(user) <= 0 {
		return false
	}
	var buffer bytes.Buffer
	ec := exec.Command("sh", "-c", listAll)
	ec.Stdout = &buffer
	err := ec.Run()
	if err != nil {
		log.Println(err.Error())
	}
	lines := strings.Split(buffer.String(), "\n")
	for i := range lines {
		if len(lines[i]) > 0 {
			temp := strings.Split(lines[i], ":")
			if user == temp[0] {
				return true
			}
		}
	}
	return false
}

// ChangePassword 修改密码
func ChangePassword(user, pwd string) bool {
	if len(pwd) <= 0 || len(user) <= 0 || !Check(user) {
		return false
	}
	var buffer bytes.Buffer
	ec := exec.Command("sh", "-c", "printf \"%s\\n%s\\n\" "+pwd+" "+pwd+" | smbpasswd -s "+user)
	ec.Stdout = &buffer
	err := ec.Run()
	if err != nil {
		log.Println(err.Error())
	}
	if len(buffer.String()) <= 0 {
		return true
	}
	return false
}

// CheckOldPassword 检查原密码
func CheckOldPassword(user, oldPwd string) bool {
	if len(oldPwd) <= 0 || len(user) <= 0 || !Check(user) {
		return false
	}
	conn, err := net.Dial("tcp", server)
	if err != nil {
		return false
	}
	defer conn.Close()
	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     user,
			Password: oldPwd,
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		return false
	}
	defer s.Logoff()
	return true
}
