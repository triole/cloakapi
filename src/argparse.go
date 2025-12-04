package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
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
	Conf        string `help:"path to config file" short:"c" default:"${configFile}"`
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
		key := line[0]
		val := line[1]
		ret += fmt.Sprintf("\n%40s: %s", key, val)
	}
	return
}

func sortedIterator(mp map[string]string) (arr []string) {
	for key := range mp {
		arr = append(arr, key)
	}
	sort.Strings(arr)
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
			"configFile":       "conf.toml",
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
