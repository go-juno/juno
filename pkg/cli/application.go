package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-juno/juno/pkg/cli/argv"
	"github.com/go-juno/juno/pkg/cli/flag"
)

var (
	// iApp
	iApp *application
	// Version
	Version = "1.0.0"
	cliName = "juno"
)

func init() {
	iApp = New("app", "0.0.0")
}

// New
func New(name, version string) *application {
	app := &application{
		Name:    name,
		Version: version,
	}
	app.BasePath = argv.Program().Dir
	return app
}

// App
func App() *application {
	return iApp
}

// SetName
func SetName(name string) *application {
	return iApp.SetName(name)
}

// SetVersion
func SetVersion(version string) *application {
	return iApp.SetVersion(version)
}

// SetDebug
func SetDebug(debug bool) *application {
	return iApp.SetDebug(debug)
}

// Use
func Use(h ...HandlerFunc) *application {
	return iApp.Use(h...)
}

// AddCommand
func AddCommand(cmds ...*Command) *application {
	iApp.AddCommand(cmds...)
	return iApp
}

// Run
func Run() {
	iApp.Run()
}

// application
type application struct {
	// 应用名称
	Name string
	// 应用版本
	Version string
	// 应用调试
	Debug bool
	// 基础路径
	BasePath string

	// 是否单命令
	singleton bool
	// 默认命令
	defaultCommand string
	// 命令集合
	commands []*Command
	// handlers
	handlers []HandlerFunc
}

// HandlerFunc
type HandlerFunc func(next func())

// SetName
func (t *application) SetName(name string) *application {
	t.Name = name
	return t
}

// SetVersion
func (t *application) SetVersion(version string) *application {
	t.Version = version
	return t
}

// SetDebug
func (t *application) SetDebug(debug bool) *application {
	t.Debug = debug
	return t
}

// Use
func (t *application) Use(h ...HandlerFunc) *application {
	t.handlers = append(t.handlers, h...)
	return t
}

// AddCommand
func (t *application) AddCommand(cmds ...*Command) *application {
	t.commands = append(t.commands, cmds...)
	// init
	for _, c := range t.commands {
		if c.Singleton {
			t.singleton = true
		}
		if c.Default {
			t.defaultCommand = c.Name
		}
	}
	if t.singleton {
		argv.Parse(true)
		flag.Parse()
	}
	return t
}

// Run 执行
func (t *application) Run() {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case *NotFoundError, *UnsupportedError:
				fmt.Println(err)
				return
			default:
				panic(err)
			}
		}
	}()

	if len(t.commands) == 0 {
		panic(errors.New("command cannot be empty"))
	}

	command := argv.Command()
	if command == "" {
		if flag.Match("h", "help").Bool(false) {
			t.globalHelp()
			return
		}
		if flag.Match("v", "version").Bool(false) {
			t.version()
			return
		}

		options := flag.Options().Map()
		if len(options) == 0 {
			if t.defaultCommand != "" && len(os.Args) == 1 {
				os.Args = append(os.Args, t.defaultCommand)
				argv.Parse()
				flag.Parse()
				t.Run()
			} else {
				t.globalHelp()
			}
			return
		}

		if t.singleton {
			t.call()
			return
		}

		f := ""
		for k := range options {
			f = k
			break
		}
		panic(NewNotFoundError(fmt.Errorf("flag provided but not defined: '%s', see '%s --help'.", f, cliName)))
	} else if flag.Match("help").Bool(false) {
		t.commandHelp()
		return
	}
	t.call()
}

func (t *application) getCommand(n string) *Command {
	var cmd *Command
	if t.singleton {
		// 单命令
		for _, c := range t.commands {
			if c.Singleton {
				cmd = c
				break
			}
		}
		if cmd == nil {
			panic(errors.New("singleton command not found"))
		}
	} else {
		for _, c := range t.commands {
			if c.Name == n {
				cmd = c
				break
			}
		}
	}
	return cmd
}

// 执行命令
func (t *application) call() {
	// 命令行选项效验
	t.validateOptions()

	// 提取命令
	command := argv.Command()
	cmd := t.getCommand(command)
	if cmd == nil {
		panic(NewNotFoundError(fmt.Errorf("'%s' is not command, see '%s --help'.", command, cliName)))
	}
	if cmd.RunF == nil && cmd.RunI == nil {
		panic(fmt.Errorf("'%s' command RunF & RunI field is empty", cmd.Name))
	}

	// 执行命令
	exec := func() {
		rf := cmd.RunF
		if rf != nil {
			rf()
			return
		}
		ri := cmd.RunI
		if ri != nil {
			ri.Main()
			return
		}
	}
	if len(t.handlers) > 0 {
		tmp := append(t.handlers, cmd.handlers...)
		for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
			tmp[i], tmp[j] = tmp[j], tmp[i]
		}
		var next func()
		for k, f := range tmp {
			if k == 0 {
				n := exec
				c := f
				next = func() {
					c(n)
				}
				if len(tmp) == 1 {
					c(n)
				}
			} else if len(tmp)-1 == k {
				f(next)
			} else {
				n := next
				c := f
				next = func() {
					c(n)
				}
			}
		}
	} else {
		exec()
	}
}

// 命令行选项效验
func (t *application) validateOptions() {
	var options []*Option
	if !t.singleton {
		for _, v := range t.commands {
			if v.Name == argv.Command() {
				options = v.Options
				break
			}
		}
	} else {
		for _, v := range t.commands {
			if v.Singleton {
				options = v.Options
				break
			}
		}
	}
	if len(options) == 0 {
		return
	}

	var flags []string
	for _, o := range options {
		for _, v := range o.Names {
			if len(v) == 1 {
				flags = append(flags, fmt.Sprintf("-%s", v))
			} else {
				flags = append(flags, fmt.Sprintf("--%s", v))
			}
		}
	}
	inArray := func(value string, values []string) bool {
		for _, v := range values {
			if v == value {
				return true
			}
		}
		return false
	}
	for f := range flag.Options().Map() {
		if !inArray(f, flags) {

			c := argv.Command()
			if c != "" {
				c = fmt.Sprintf(" %s", c)
			}
			panic(NewNotFoundError(fmt.Errorf("flag provided but not defined: '%s', see '%s%s --help'.", f, cliName, c)))
		}
	}
}

// 全局帮助
func (t *application) globalHelp() {
	command := argv.Command()
	cmd := t.getCommand(command)
	if command != "" && cmd == nil {
		panic(NewNotFoundError(fmt.Errorf("'%s' is not command, see '%s --help'.", command, cliName)))
	}

	if cmd != nil && cmd.Long != "" {
		fmt.Println(cmd.Long)
		fmt.Println("")
	}
	if !t.singleton {
		fmt.Printf("Usage: %s [GLOBAL OPTIONS] COMMAND [ARG...]", cliName)
	} else {
		if cmd != nil && cmd.UsageF != "" {
			fmt.Printf(cmd.UsageF, cliName)
		} else {
			fmt.Printf("Usage: %s [ARG...]", cliName)
		}
	}
	if !t.singleton {
		t.printCommands()
	} else {
		t.printCommandOptions()
	}
	t.printGlobalOptions()
	fmt.Println("")
	fg := ""
	if !t.singleton {
		fg = " COMMAND"
	}
	fmt.Printf("Run '%s%s --help' for more information on a command.", cliName, fg)
	fmt.Println("")
}

// 命令帮助
func (t *application) commandHelp() {
	command := argv.Command()
	cmd := t.getCommand(command)
	if cmd == nil {
		panic(NewNotFoundError(fmt.Errorf("'%s' is not command, see '%s --help'.", command, cliName)))
	}

	if cmd.Long != "" {
		fmt.Println(cmd.Long)
		fmt.Println("")
	}
	if cmd.UsageF != "" {
		fmt.Printf(cmd.UsageF, cliName, command)
	} else {
		fmt.Printf("Usage: %s %s [ARG...]", cliName, command)
	}
	t.printCommandOptions()
	fmt.Println("")
}

// 打印全局选项
func (t *application) printGlobalOptions() {
	tabs := "\t"
	fmt.Println("")
	fmt.Println("Global Options:")
	fmt.Printf("  -h, --help%sPrint usage", tabs)
	fmt.Printf("  -v, --version%sPrint version information", tabs)
}

// 打印命令
func (t *application) printCommands() {
	fmt.Println("")
	fmt.Println("Commands:")
	for _, v := range t.commands {
		command := v.Name
		short := v.Short
		fmt.Printf("  %s\t%s", command, short)
	}
}

// 打印命令选项
func (t *application) printCommandOptions() {
	var options []*Option
	if !t.singleton {
		for _, v := range t.commands {
			if v.Name == argv.Command() {
				options = v.Options
				break
			}
		}
	} else {
		for _, v := range t.commands {
			if v.Singleton {
				options = v.Options
				break
			}
		}
	}
	if len(options) == 0 {
		return
	}

	fmt.Println("")
	fmt.Println("Command Options:")
	for _, o := range options {
		var flags []string
		for _, v := range o.Names {
			if len(v) == 1 {
				flags = append(flags, fmt.Sprintf("-%s", v))
			} else {
				flags = append(flags, fmt.Sprintf("--%s", v))
			}
		}
		fg := strings.Join(flags, ", ")
		fmt.Printf("  %s\t%s", fg, o.Usage)
	}
}

// 版本号
func (t *application) version() {
	appName := t.Name
	appVersion := t.Version
	frameworkVersion := Version
	fmt.Printf("%s %s, framework %s", appName, appVersion, frameworkVersion)
}
