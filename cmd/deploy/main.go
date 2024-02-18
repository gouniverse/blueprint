package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"

	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/utils"
	"github.com/mingrammer/cfmt"
	"github.com/samber/lo"
	"github.com/sfreiberg/simplessh"
)

func main() {
	timestamp := carbon.Now(carbon.UTC).Format("Ymd_His")
	buildExecutablePath := "tmp/application_deploy_" + timestamp
	sshKey := "{{ SSHKEY }}.prv"
	sshUser := "{{ SSHUSER }}"
	sshHost := "{{ SSHHOST }}"
	sshLogin := sshUser + "@" + sshHost
	remDeployDir := "/home/{{ USER }}/{{ APP_DIR }}"
	remTempDeployName := "temp_deploy_" + timestamp
	pm2ProcessName := "{{ PROCESSNAME }}"

	cfmt.Infoln("1. Building executable...")

	err := buildExecutable(buildExecutablePath)

	if err != nil {
		log.Fatal(err)
		return
	}

	cfmt.Infoln("2. Uploading executable...")

	cmd := `scp -o stricthostkeychecking=no -i ` + privateKeyPath(sshKey) + ` ` + buildExecutablePath + ` ` + sshLogin + `:` + remDeployDir + `/` + remTempDeployName

	cfmt.Infoln(" - Executing:" + cmd)
	utils.ExecLine(cmd)

	cfmt.Infoln("3. Replace current executable...")

	cmds := []struct {
		cmd      string
		required bool
	}{
		{
			cmd:      `chmod 750 ` + remDeployDir + `/` + remTempDeployName,
			required: true,
		},
		{
			cmd:      `mv ` + remDeployDir + `/application  ` + remDeployDir + `/` + timestamp + `_backup_application`,
			required: true,
		},
		{
			cmd:      `mv ` + remDeployDir + `/` + remTempDeployName + `  ` + remDeployDir + `/application`,
			required: true,
		},
		{
			cmd:      `mv ` + remDeployDir + `/application.error.log ` + remDeployDir + `/` + timestamp + `_backup_application.error.log`,
			required: false,
		},
		{
			cmd:      `mv ` + remDeployDir + `/application.log ` + remDeployDir + `/` + timestamp + `_backup_application.log`,
			required: false,
		},
		{
			cmd:      `pm2 restart ` + pm2ProcessName,
			required: true,
		},
	}

	for _, entry := range cmds {
		cfmt.Infoln(" - Executing:" + entry.cmd)
		output, error := ssh(sshHost, sshUser, sshKey, entry.cmd)

		if error != nil {
			cfmt.Errorln("  - Error:", error)
			cfmt.Errorln("  - Output: ", output)
			if entry.required {
				return // stop on first error, if required
			}
		}

		cfmt.Successln("  - Output: ", lo.Ternary(output == "", "no output", output))
	}

	cfmt.Infoln("Deployed!")

}

func buildExecutable(pathExec string) error {
	newEnv := os.Environ()
	newEnv = append(newEnv, "GOOS=linux")
	newEnv = append(newEnv, "GOARCH=amd64")
	newEnv = append(newEnv, "CGO_ENABLED=0")

	cmd := exec.Command("go", "build", "-ldflags", "-s -w", "-v", "-o", pathExec, "main.go")
	cmd.Env = newEnv
	out, err := cmd.CombinedOutput()

	if err != nil {
		cfmt.Errorln(string(out))
	} else {
		cfmt.Successln(string(out))
	}

	return err
}

func privateKeyPath(sshKey string) string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	homeDirectory := user.HomeDir
	privateKeyPath := homeDirectory + "/.ssh/" + sshKey
	return privateKeyPath
}

func ssh(sshHost, sshUser, sshKey, cmd string) (output string, err error) {
	client, err := simplessh.ConnectWithKeyFile(sshHost+":22", sshUser, privateKeyPath(sshKey))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	outputBytes, err := client.Exec(cmd)

	if err != nil {
		return string(outputBytes), err
	}

	return string(outputBytes), nil
}
