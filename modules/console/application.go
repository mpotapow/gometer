package console

import (
	"fmt"
	"gometer/modules/console/contracts"
	"sort"
	"strings"
)

// Application ...
type Application struct {
	name      string
	version   string
	formatter *Formatter
	commands  map[string]contracts.Command
}

// GetApplicationInstance ...
func GetApplicationInstance() contracts.Application {
	return &Application{
		formatter: GetFormatterInstance(),
		commands: map[string]contracts.Command{
			"list": GetCommandInstance("list", "Lists commands"),
		},
	}
}

// SetName ...
func (a *Application) SetName(name string) {
	a.name = name
}

// SetVersion ...
func (a *Application) SetVersion(version string) {
	a.version = version
}

// Add ...
func (a *Application) Add(command contracts.Command) {
	a.commands[command.GetName()] = command
}

// Run ...
func (a *Application) Run(args []string) {

	if len(args) == 1 {
		a.showCommandList()
		return
	}

	options := args[2:]
	commandName := args[1:2][0]

	if commandName == "list" {
		a.showCommandList()
		return
	}

	if c, ok := a.commands[commandName]; ok {
		if a.hasHelpOption(options) {
			a.showCommandHelp(c)
		}

		a.setOptions(c, options)
		c.Handle(a.formatter)
	}
}

func (a *Application) showCommandList() {

	a.formatter.Writeln(fmt.Sprintf("%s <green>%s</>", a.name, a.version))
	a.formatter.NewLine()

	a.formatter.Writeln("<yellow>Options:</>")
	a.formatter.Text("<green>--help</>  Display this help message")
	a.formatter.NewLine()

	a.formatter.Writeln("<yellow>Usage:</>")
	a.formatter.Text("command [options] [arguments]")
	a.formatter.NewLine()

	var maxNameLen int
	var sortGroup []string

	var availableCommand []contracts.Command
	groupedCommand := make(map[string][]contracts.Command)

	for _, c := range a.commands {
		if len(c.GetName()) > maxNameLen {
			maxNameLen = len(c.GetName())
		}

		if strings.Index(c.GetName(), ":") != -1 {
			groupName := strings.Split(c.GetName(), ":")
			if len(groupName) == 2 {
				if _, ok := groupedCommand[groupName[0]]; !ok {
					sortGroup = append(sortGroup, groupName[0])
				}
				groupedCommand[groupName[0]] = append(groupedCommand[groupName[0]], c)
				continue
			}
		}

		availableCommand = append(availableCommand, c)
	}

	a.formatter.Writeln("<yellow>Available commands:</>")
	sort.Sort(CommandByName(availableCommand))
	for _, c := range availableCommand {
		a.printCommandInfo(c, maxNameLen)
	}

	sort.Strings(sortGroup)
	for _, group := range sortGroup {
		a.formatter.Writeln(fmt.Sprintf(" <yellow>%s:</>", group))
		sort.Sort(CommandByName(groupedCommand[group]))
		for _, c := range groupedCommand[group] {
			a.printCommandInfo(c, maxNameLen)
		}
	}
}

func (a *Application) showCommandHelp(c contracts.Command) {

	a.formatter.Writeln("<yellow>Description:</>")
	a.formatter.Text(c.GetDescription())
	a.formatter.NewLine()

	a.formatter.Writeln("<yellow>Usage:</>")
	a.formatter.Text(fmt.Sprintf("%s [options] %s", c.GetName(), a.formatArgsForPrint(c.GetArguments())))
	a.formatter.NewLine()

	if len(c.GetArguments()) > 0 {
		a.formatter.Writeln("<yellow>Arguments:</>")

		var maxArgLen int
		for _, arg := range c.GetArguments() {
			if len(arg.GetName()) > maxArgLen {
				maxArgLen = len(arg.GetName())
			}
		}

		for _, arg := range c.GetArguments() {
			a.printArgumentInfo(arg, maxArgLen)
		}
	}

	if len(c.GetOptions()) > 0 {
		a.formatter.Writeln("<yellow>Options:</>")

		var maxOptionLen int
		options := make([]contracts.Option, 0, len(c.GetOptions()))
		for _, option := range c.GetOptions() {
			nameLen := len(option.GetName())
			if option.IsFillable() {
				nameLen = nameLen*2 + 3
			}
			if nameLen > maxOptionLen {
				maxOptionLen = nameLen
			}
			options = append(options, option)
		}

		sort.Sort(OptionByName(options))
		for _, option := range options {
			a.printOptionInfo(option, maxOptionLen)
		}
	}
}

func (a *Application) printCommandInfo(c contracts.Command, offsetSize int) {
	separator := strings.Repeat(" ", offsetSize-len(c.GetName()))
	a.formatter.Writeln(fmt.Sprintf("  <green>%s</> %s %s", c.GetName(), separator, c.GetDescription()))
}

func (a *Application) printArgumentInfo(arg contracts.Argument, offsetSize int) {
	separator := strings.Repeat(" ", offsetSize-len(arg.GetName()))
	a.formatter.Writeln(fmt.Sprintf("  <green>%s</> %s %s", arg.GetName(), separator, arg.GetDescription()))
}

func (a *Application) printOptionInfo(option contracts.Option, offsetSize int) {
	name := option.GetName()
	if option.IsFillable() {
		name += fmt.Sprintf("[=%s]", strings.ToUpper(option.GetName()))
	}
	separator := strings.Repeat(" ", offsetSize-len(name))
	a.formatter.Writeln(fmt.Sprintf("  <green>--%s</> %s %s", name, separator, option.GetDescription()))
}

func (a *Application) hasHelpOption(options []string) bool {
	for _, v := range options {
		if v == "--help" {
			return true
		}
	}
	return false
}

func (a *Application) formatArgsForPrint(args []contracts.Argument) string {

	result := ""
	for _, arg := range args {
		result += fmt.Sprintf("[<%s>]", arg.GetName()) + " "
	}
	return result
}

func (a *Application) setOptions(command contracts.Command, options []string) {

	var args []string
	optionList := command.GetOptions()

	for _, option := range options {
		if strings.Index(option, "--") != -1 {
			value := ""
			option = strings.TrimLeft(option, "--")
			if strings.Index(option, "=") != -1 {
				optionData := strings.Split(option, "=")
				option = optionData[0]
				value = optionData[1]
			}

			if optionObj, ok := optionList[option]; ok {
				if optionObj.IsFillable() {
					optionObj.SetValue(value)
				} else {
					optionObj.SetValue(true)
				}
			}
			continue
		}
		args = append(args, option)
	}

	for i, argument := range command.GetArguments() {
		if len(args) > i {
			argument.SetValue(args[i])
		} else {
			break
		}
	}
}
