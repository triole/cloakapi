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
	appName        = "cloakapi"
	appDescription = "an abstraction layer for the Keycloak Admin API"
	appMainversion = "0.1"

	commands = tCommands{
		List: tCommandsList{
			FedIDs:            "fed",
			Users:             "usr",
			UserSessions:      "use",
			UserAttributes:    "att",
			IdentityProviders: "idp",
			AuthFlows:         "flw",
		},
		Remove: tCommandsList{
			Users: "usr",
		},
	}
)

type tCommands struct {
	List   tCommandsList
	Remove tCommandsList
}

type tCommandsList struct {
	FedIDs            string
	Users             string
	UserSessions      string
	UserAttributes    string
	IdentityProviders string
	AuthFlows         string
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
		Entity string `help:"list entity" arg:"" enum:"${listCommands}"`
	} `cmd:"" help:"list entity, available commands: ${listCommands}"`

	Rm struct {
		Entity string `help:"remove entity" arg:"" enum:"${removeCommands}"`
		Target string `help:"remove entity" arg:""`
	} `cmd:"" help:"remove entity, available commands: ${removeCommands}"`

	Tpl struct {
		String string `help:"execute template string" short:"s"`
		File   string `help:"template file path" short:"f"`
	} `cmd:"" help:"execute template string or load from file"`

	Var struct {
	} `cmd:"" help:"list available template variables"`
}

func pprintCommandList(cmds any) (ret string) {
	v := reflect.ValueOf(cmds)
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).String()
		if val != "" {
			ret += fmt.Sprintf("%s, ", val)
		}
	}
	ret = strings.TrimSuffix(ret, ", ")
	return
}

func parseArgs() {
	listCommands := pprintCommandList(commands.List)
	removeCommands := pprintCommandList(commands.Remove)
	ctx := kong.Parse(&cli,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}),
		kong.Vars{
			"listCommands":   listCommands,
			"removeCommands": removeCommands,
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
