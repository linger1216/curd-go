package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"strings"
)

func spaceValid(input string) error {
	if strings.Contains(input, " ") {
		return fmt.Errorf("has space")
	}
	return nil
}

type CliConfig struct {
	departmentName  string
	projectName     string
	applicationName string
	maintainer      string
	postgresUrl     string
	localCache      bool
}

func NewCliConfig() *CliConfig {
	return &CliConfig{departmentName: "location"}
}

func main() {
	//cliConfig := NewCliConfig()
	//var err error
	//
	//cliConfig.projectName, err = inputCli("projectName", "DefaultProjectName", spaceValid)
	//if err != nil || len(cliConfig.projectName) == 0 {
	//	return
	//}
	//
	//cliConfig.applicationName, err = inputCli("applicationName", "DefaultApplicationName", spaceValid)
	//if err != nil || len(cliConfig.applicationName) == 0 {
	//	return
	//}
	//
	//cliConfig.maintainer, err = inputCli("maintainer", "DefaultMaintainer", spaceValid)
	//if err != nil || len(cliConfig.maintainer) == 0 {
	//	return
	//}
	//
	//cliConfig.postgresUrl, err = inputCli("postgresUrl", "", nil)
	//if err != nil || len(cliConfig.postgresUrl) == 0 {
	//	return
	//}
	//
	//cliConfig.localCache, err = ConfirmCli("need local cache")
	//if err != nil {
	//	return
	//}
	//
	//confirm, err := ConfirmCli("all correct")
	//if err != nil {
	//	return
	//}
	//
	//if confirm {
	//	fmt.Println(cliConfig)
	//}
	//
	//cliConfig.postgresUrl = "postgres://lid.guan:@localhost:15432/zhigan?sslmode=disable"

	//pg := meta.NewPostgresMeta(utils.NewPostgres(&utils.PostgresConfig{
	//	Uri:     cliConfig.postgresUrl,
	//	MaxIdle: 0,
	//	MaxOpen: 0,
	//}))
	//
	//pg.GetInfo()

}

func inputCli(label, def string, validFn func(input string) error) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validFn,
		Default:  def,
	}
	return prompt.Run()
}

func selectCli(label string, items ...string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, r, e := prompt.Run()
	return r, e
}

func ConfirmCli(label string) (bool, error) {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	r, err := prompt.Run()
	if err != nil {
		return false, err
	}
	if strings.ToLower(r) == "y" {
		return true, nil
	}
	return false, nil
}
