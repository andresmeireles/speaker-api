package cli

import (
	"fmt"
	"os"
	"reflect"

	"github.com/andresmeireles/speaker/internal/cli/commands"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	"github.com/spf13/cobra"
)

func Commands(sl servicelocator.ServiceLocator) {
	cmd := &cobra.Command{}
	commands := []any{
		commands.MigrateUp,
		commands.MigrateDown,
		commands.SetAppKey,
		commands.CreateUser,
		commands.ListUser,
		commands.ShowNumberOfUnusedDependencies,
	}
	resolvedCommands := []*cobra.Command{}

	for _, v := range commands {
		resolvedCommands = append(resolvedCommands, resolve(sl, v))
	}

	cmd.AddCommand(resolvedCommands...)

	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func resolve(sl servicelocator.ServiceLocator, commandFunction any) *cobra.Command {
	command := reflect.TypeOf(commandFunction)
	commandKind := command.Kind()
	isFunc := commandKind == reflect.Func

	if !isFunc {
		panic("command is not a function: " + commandKind.String())
	}

	outValue := command.Out(0).String()
	if outValue != "*cobra.Command" {
		panic("command function must return a cobra.Command: " + commandKind.String())
	}

	numOfParams := command.NumIn()
	commandParams := []reflect.Value{}

	for i := 0; i < numOfParams; i++ {
		name := command.In(i).String()
		param := sl.Get(name)

		commandParams = append(commandParams, reflect.ValueOf(param))
	}

	r, ok := reflect.ValueOf(commandFunction).Call(commandParams)[0].Interface().(*cobra.Command)
	if !ok {
		panic("command function must return a *cobra.Command: " + commandKind.String())
	}

	return r
}
