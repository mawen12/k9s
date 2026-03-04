// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package cmd

import (
	"fmt"

	"github.com/derailed/k9s/internal/color"
	"github.com/spf13/cobra"
)

//
//

/*
versionCmd 代表 version 命令

	k9s version
		访问完整版本信息

	k9s version -s
		访问精简版本信息

	k9s version --short
		访问精简版本信息
*/
func versionCmd() *cobra.Command {
	// 由命令行传递，代表是否打印彩色 Logo
	var short bool

	command := cobra.Command{
		Use:   "version",
		Short: "Print version/build info",
		Long:  "Print version/build information",
		Run: func(*cobra.Command, []string) {
			printVersion(short)
		},
	}

	command.PersistentFlags().BoolVarP(&short, "short", "s", false, "Prints K9s version info in short format")

	return &command
}

func printVersion(short bool) {
	const fmat = "%-20s %s\n"
	var outputColor color.Paint

	if short {
		outputColor = -1
	} else {
		outputColor = color.Cyan
		printLogo(outputColor)
	}
	printTuple(fmat, "Version", version, outputColor)
	printTuple(fmat, "Commit", commit, outputColor)
	printTuple(fmat, "Date", date, outputColor)
}

func printTuple(fmat, section, value string, outputColor color.Paint) {
	if outputColor != -1 {
		_, _ = fmt.Fprintf(out, fmat, color.Colorize(section+":", outputColor), value)
		return
	}
	_, _ = fmt.Fprintf(out, fmat, section, value)
}
