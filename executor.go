package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"
)

type Executor interface {
	Execute(pwdId int, dutyCycle int) error
}

type IPMIExecutor struct{}

func (i *IPMIExecutor) Execute(pwdId int, dutyCycle int) error {
	hexPwdId := fmt.Sprintf("%#02x", pwdId)
	hexDutyCycle := fmt.Sprintf("%#02x", dutyCycle)
	rawCodes := []string{"raw", "0x2e", "0x44", "0xfd", "0x19", "0x00", hexPwdId, "0x01", hexDutyCycle}
	log.Println("IPMI executor: ipmitool", strings.Join(rawCodes, " "))
	_, err, exitCode := RunCommand("ipmitool", rawCodes...)
	if exitCode != 0 {
		return errors.New(err)
	}
	return nil
}

type FakeExecutor struct {
}

func (f *FakeExecutor) Execute(pwdId int, dutyCycle int) error {
	log.Printf("Fake executor: CMD=ipmitool raw 0x2e 0x44 0xfd 0x19 0x00 %#02x 0x01 %#02x", pwdId, dutyCycle)
	return nil
}

func RunCommand(name string, args ...string) (stdout string, stderr string, exitCode int) {
	var outBuffer, errBuffer bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	err := cmd.Run()
	stdout = outBuffer.String()
	stderr = errBuffer.String()

	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		}
	} else {
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	return
}
