// TODO: change names for better expressiveness
package commands

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/andresmeireles/speaker/internal"
	"github.com/andresmeireles/speaker/internal/tools/servicelocator"
	"github.com/spf13/cobra"
)

func createDI() servicelocator.ServiceLocator {
	sl := servicelocator.NewServiceLocator()
	internal.DIContainer(sl)

	return *sl
}

func showDependencyUses(sl servicelocator.ServiceLocator) map[string]int {
	deps := map[string]int{}
	implementations := []any{}

	for k, v := range sl.GetServices() {
		deps[k] = 0

		implementations = append(implementations, v)
	}

	for _, imp := range implementations {
		ref := reflect.TypeOf(imp)
		if ref.Kind() == reflect.Pointer {
			ref = ref.Elem()
		}

		numOfFields := ref.NumField()

		for i := 0; i < numOfFields; i++ {
			field := ref.Field(i).Type.String()

			if _, ok := deps[field]; ok {
				deps[field]++
			}
		}
	}

	return deps
}

func ShowNumberOfUnusedDependencies() *cobra.Command {
	return &cobra.Command{
		Use:   "deps",
		Short: "Show number of unused dependencies",
		Run: func(cmd *cobra.Command, args []string) {
			sl := createDI()
			result := showDependencyUses(sl)

			for k, v := range result {
				if strings.Contains(k, "Controller") {
					continue
				}

				if v == 0 {
					fmt.Print("!!!!! => ")
				}
				fmt.Println("depedency:", k, "number of uses:", v)
			}
		},
	}
}
