package helpers

import "fmt"

func Pointer[V int | string](v V) *V {
	return &v
}

func ValidateArgs(args ...string) error {
	// This should probably live somewhere else
	var nilCount int

	for _, arg := range args {
		if arg == "" {
			nilCount++
		}
	}

	if nilCount != len(args)-1 {
		return fmt.Errorf("provide either the --pat or --bearer argument")
	}

	return nil
}