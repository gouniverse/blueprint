package cmds

import (
	// "project/internal/tasks"
	"project/config"
	"strings"

	"github.com/gouniverse/utils"
	"github.com/mingrammer/cfmt"
)

// ExecuteCommand executes a task from the command line
func ExecuteCommand(name string, args []string) bool {
	argumentsMap := utils.ArgsToMap(args)
	cfmt.Infoln("Executing command: ", name, " with arguments: ", argumentsMap)

	if name == "list" {
		for index, taskHandler := range config.Cms.TaskStore.TaskHandlerList() {
			cfmt.Warningln(utils.ToString(index+1) + ". Task with alias: " + taskHandler.Alias())
			cfmt.Infoln("    - Human Frienly Title: " + taskHandler.Title())
			cfmt.Infoln("    - Human Frienly Description: " + taskHandler.Description())
		}
		return true
	}

	for _, taskHandler := range config.Cms.TaskStore.TaskHandlerList() {
		if strings.EqualFold(unifyName(taskHandler.Alias()), unifyName(name)) {
			taskHandler.SetOptions(argumentsMap)
			taskHandler.Handle()
			return true
		}
	}

	cfmt.Errorln("Unrecognized command: ", name)
	return false
}

func unifyName(name string) string {
	name = strings.ReplaceAll(name, "-", "")
	name = strings.ReplaceAll(name, "_", "")
	return name
}

// var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
// var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// func toSnakeCase(str string) string {
// 	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
// 	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
// 	return strings.ToLower(snake)
// }

// func toSlugCase(str string) string {
// 	snake := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
// 	snake = matchAllCap.ReplaceAllString(snake, "${1}-${2}")
// 	return strings.ToLower(snake)
// }
