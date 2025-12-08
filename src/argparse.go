package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	BUILDTAGS      string
	appName        = "cloakAPI"
	appDescription = "an abstraction layer for the Keycloak Admin API"
	appMainversion = "0.1"

	commands = tCommands{
		List: tCommandsList{
			FedIDs:            "federated user ids:fed",
			Users:             "users:usr",
			UserAttributes:    "user attributes:att",
			IdentityProviders: "identity providers:idp",
		},
	}
)

type tCommands struct {
	List tCommandsList
}

type tCommandsList struct {
	FedIDs            string
	Users             string
	UserAttributes    string
	IdentityProviders string
}

var cli struct {
	Action      string `kong:"-" enum:"conf,exec,calc" default:"conf"`
	Conf        string `help:"config file detection expression" required:"" short:"c"`
	Output      string `help:"output format" short:"o" enum:"json,toml,yaml,table" default:"table"`
	LogFile     string `help:"log file" default:"/dev/stdout"`
	LogLevel    string `help:"log level" default:"info" enum:"trace,debug,info,error"`
	LogNoColors bool   `help:"disable output colours, print plain text"`
	LogJSON     bool   `help:"enable json log, instead of text one"`
	DryRun      bool   `help:"dry run, just print operations that would run" short:"n"`
	VersionFlag bool   `help:"display version" short:"V"`

	Ls struct {
		Entity string `help:"entity to list" arg:"" enum:"${listCommandEnums}"`
	} `cmd:"" help:"list entities, available commands: ${listCommands}"`

	Tpl struct {
		String string `help:"execute template string" short:"s"`
		File   string `help:"template file path" short:"f"`
	} `cmd:"" help:"execute template string or load from file"`

	Var struct {
	} `cmd:"" help:"list available template variables"`
}

func getCommandEnums(cmds any) (ret string) {
	v := reflect.ValueOf(cmds)
	for i := 0; i < v.NumField(); i++ {
		cmd := getCommand(v.Field(i).String())
		ret += cmd + ","
	}
	ret = strings.TrimSuffix(ret, ",")
	return
}

func getCommand(str string) (ret string) {
	arr := strings.Split(str, ":")
	if len(arr) > 1 {
		ret = arr[1]
	}
	return
}

func pprintCommandList(cmds any) (ret string) {
	v := reflect.ValueOf(cmds)
	for i := 0; i < v.NumField(); i++ {
		line := strings.Split(v.Field(i).String(), ":")
		// key := line[0]
		val := line[1]
		ret += fmt.Sprintf("%s, ", val)
	}
	ret = strings.TrimSuffix(ret, ", ")
	return
}

func parseArgs() {
	listCommands := pprintCommandList(commands.List)
	listCommandEnums := getCommandEnums(commands.List)
	ctx := kong.Parse(&cli,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
		kong.Vars{
			"listCommands":     listCommands,
			"listCommandEnums": listCommandEnums,
		},
	)
	_ = ctx.Run()
	// ctx.FatalIfErrorf(err)
	cli.Action = strings.Split(ctx.Command(), " ")[0]
	if cli.Action == "version" {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "Version: "+appMainversion+".", -1)
	fmt.Printf("%s\n", s)
}
