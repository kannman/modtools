package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/kannman/modtools/cli"
	"github.com/kannman/modtools/file"
	"github.com/kannman/modtools/output"
	"github.com/spf13/cobra"
)

var (
	rootCmd     = &cobra.Command{Use: "modtools"}
	flagVerbose bool
)

var pathCmd = &cobra.Command{
	Use:   "path [module]",
	Short: "get path for current version of module",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		listjson := cli.ReadListJSON()

		mods := file.ParseDepsListJSON(listjson)

		for _, mod := range mods {
			if mod.Path != args[0] {
				continue
			}
			moddir := mod.Dir

			if mod.Replace != nil {
				moddir = mod.Replace.Dir
			}
			cmd.Println(moddir)
			return
		}

		output.Error("module '%s' not found", args[0])
	},
}

var vendCmd = &cobra.Command{
	Use:   "vendor",
	Short: "copy files matching glob pattern to ./vendor (**/*.c **/*.h **/*.proto)",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		listJSON := cli.ReadListJSON()

		modules := file.ParseDepsListJSON(listJSON)
		copyCount := 0
		for _, mod := range modules[1:] {
			modFiles := cli.BuildModVendorList(args, mod.Dir)

			for _, vendorFile := range modFiles {
				x := strings.Index(vendorFile, mod.Dir)
				if x < 0 {
					output.Error("vendor file '%s' doesn't belong to mod '%s', strange.", vendorFile, mod.Path)
				}

				localFile := path.Join("vendor", mod.Path, vendorFile[len(mod.Dir):])
				os.MkdirAll(filepath.Dir(localFile), os.ModePerm)

				if flagVerbose {
					output.Print("copy %s %s", vendorFile, localFile)
				}

				cli.CopyFile(vendorFile, localFile)
				copyCount++
			}
		}

		output.Info("%d files copied", copyCount)
	},
}

func main() {
	rootCmd.AddCommand(
		pathCmd,
		vendCmd,
	)

	rootCmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "verbose output")
	rootCmd.Execute()
}
