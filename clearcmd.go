package logger

import (
	"fmt"
	"path/filepath"

	"github.com/bopher/utils"
	"github.com/spf13/cobra"
)

// ClearCommand get clear logs
func ClearCommand(logPath string) *cobra.Command {
	var cmd = new(cobra.Command)
	cmd.Use = "clear [Directory name or all for clear anything]"
	cmd.Short = "clear log directory"
	cmd.Args = cobra.MinimumNArgs(1)
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if isDir, err := utils.IsDirectory(logPath); err != nil {
			fmt.Printf("failed: %s\n", err.Error())
			return
		} else if !isDir {
			fmt.Printf("failed: directory is invalid or not found!\n")
			return
		}

		if args[0] == "all" {
			dirs, err := utils.GetSubDirectory(logPath)
			if err != nil {
				fmt.Printf("failed: %s\n", err.Error())
				return
			}
			for _, dir := range dirs {
				if err := utils.ClearDirectory(filepath.Join(logPath, dir)); err != nil {
					fmt.Printf("failed: %s\n", err.Error())
					return
				}
			}
		} else {
			if isDir, err := utils.IsDirectory(filepath.Join(logPath, args[0])); err != nil {
				fmt.Printf("failed: %s\n", err.Error())
				return
			} else if !isDir {
				fmt.Printf("failed: %s log directory not found\n", args[0])
				return
			}

			if err := utils.ClearDirectory(filepath.Join(logPath, args[0])); err != nil {
				fmt.Printf("failed: %s\n", err.Error())
				return
			}
		}
		fmt.Printf("cleared!\n")
	}
	return cmd
}
