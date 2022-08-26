/*
Copyright © 2022 David Hay davidhaydev@proton.me
*/
package cmd

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var templates embed.FS

// handleExisting a .gitignore already exists delete it
func handleExisting() {
	_, err := os.Stat(".gitignore")

	if os.IsNotExist(err) {
		return
	}

	var ans string
	fmt.Println(".gitignore detected. Would you like to create a fresh file? [y/n]")
	fmt.Scanln(&ans)

	if ans == "y" {
		fmt.Println("Deleting...")
		os.Remove(".gitignore")
	} else {
		fmt.Println("Appending...")
	}
}

var supportedFiles = []string{"extjs", "zendframework", "sugarcrm", "chefcookbook", "joomla", "cfwheels", "jenv", "virtualenv", "eagle", "objective-c", "java", "nim", "codeki", "rus", "sass", "concrete5", "flexbuilder", "haskell", "gw", "elisp", "otto", "matlab", "dropbo", "redcar", "android", "drupal", "linu", "phalcon", "dm", "backup", "syncthing", "cvs", "labview", "cuda", "coq", "microsoftoffice", "prestashop", "putty", "igorpro", "nanoc", "espresso", "elm", "purescrip", "lua", "magento", "leiningen", "dreamweaver", "unity", "urbogears2", "e", "gitbook", "metals", "mercurial", "smalltalk", "sketchup", "gpg", "ortoisegi", "code", "archives", "bazaar", "contributing", "delphi", "idris", "vim", "expressionengine", "racke", "waf", "erlang", "jdeveloper", "python", "jekyll", "metaprogrammingsystem", "diff", "seamgen", "svn", "grails", "cakephp", "lilypond", "scons", "perl", "scala", "symfony", "ruby", "kate", "ada", "appengine", "libreoffice", "windows", "archlinuxpackages", "rails", "c", "actionscrip", "finale", "zephir", "forcedotcom", "symphonycms", "gcov", "vvvv", "raku", "mercury", "eiffelstudio", "agda", "webmethods", "episerver", "appceleratortitanium", "vagran", "al", "plone", "maven", "jboss", "visualstudio", "dar", "ocaml", "opencar", "eclipse", "laravel", "r", "sb", "opa", "jenkins_home", "ensime", "d", "julia", "processing", "playframework", "elixir", "craftcms", "octave", "scheme", "flaxengine", "images", "slickedi", "erraform", "ly", "virtuoso", "anjuta", "readme", "composer", "node", "ypo3", "visualstudiocode", "godo", "emacs", "oracleforms", "notepadpp", "yeoman", "lithium", "unrealengine", "packer", "ilinxise", "synopsysvcs", "gradle", "autotools", "yii", "wordpress", "patch", "commonlisp", "netbeans", "rhodesrhomobile", "q", "lemonstand", "extpattern", "sublimete", "lazarus", "stella", "go", "cloud9", "calabash", "bricxcc", "fuelphp", "monodevelop", "swif", "redis", "ags", "momentics", "extmate", "darteditor", "ninja", "modelsim", "qooxdoo", "ansible", "cmake", "scrivener", "macos", "psoccreator", "jetbrains", "wincat3", "ros", "ojo", "fancy", "stata", "kicad", "kohana", "sdcc", "codeigniter", "kdevelop4"}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:       "new",
	Short:     "Create a new .gitignore",
	ValidArgs: supportedFiles,
	Long: `
create new .gitignore followed by the types of languages, os, etc. that are relevant.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please enter the types of files your project will contain")
			return
		}

		handleExisting() // check if user has file

		ignoreFile, err := os.OpenFile(".gitignore",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		defer func() {
			if err := ignoreFile.Close(); err != nil {
				panic(err)
			} else {
				fmt.Println("All done, go build something awesome!")
			}
		}()

		for _, v := range args {
			data, err := templates.ReadFile("templates/" + strings.ToLower(v) + ".txt")

			if err == nil {
				if _, err = ignoreFile.Write(data); err != nil {
					panic(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
