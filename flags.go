package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type intelArgs struct {
	JSON           bool
	Loop           bool
	Count          bool
	CDN            bool
	Total          bool
	Limit          bool
	ReverseCompare bool
	Provider       string
	Method         string
	Compare        string
	Body           string
	BodyFile       string
}

func dropFlags(args []string) []string {

	var argsString string = strings.Join(args, " ")

	if strings.Contains(argsString, "--") {

		var re = regexp.MustCompile(`--(.*)`)
		var arr []string

		rep := re.ReplaceAllString(argsString, ``)
		arr = strings.Split(rep, " ")

		return arr[:len(arr)-1]

	}

	return args

}

func defineIntelArgumentFlags(intelFlags *flag.FlagSet, args *intelArgs) {

	intelFlags.StringVar(&args.Provider, "provider", "", "set providers")
	intelFlags.StringVar(&args.Method, "method", "", "http request method")
	intelFlags.StringVar(&args.Body, "body", "", "request body")
	intelFlags.StringVar(&args.BodyFile, "body-file", "", "request body file")
	intelFlags.StringVar(&args.Compare, "compare", "", "compare response")
	intelFlags.BoolVar(&args.ReverseCompare, "rc", false, "reverse compare")
	intelFlags.BoolVar(&args.JSON, "json", false, "show output as json")
	intelFlags.BoolVar(&args.Count, "count", false, "add count=true")
	intelFlags.BoolVar(&args.CDN, "cdn", false, "add cdn=true")
	intelFlags.BoolVar(&args.Total, "total", false, "add total=true")
	intelFlags.BoolVar(&args.Loop, "loop", false, "get all pages")
	intelFlags.BoolVar(&args.Limit, "limit", false, "limit results")

}

func showHelp(args []string, config map[string]interface{}) {

	if args[0] != "help" {
		return
	}

	if len(args) == 2 && args[1] == "flags" {
		fmt.Printf(`Flags:
  --body string
        request body
  --body-file string
        request body file
  --compare string
        compare response
  --rc
        reverse compare
  --count
        add count=true
  --json
        show output as json
  --limit
        limit results
  --loop
        get all pages
  --method string
        http request method
  --provider string
        set providers
  --cdn
        add cdn=ture
  --total
        add total=ture
`)
	} else {
		fmt.Printf(`watch help flags
watch version
watch init

watch get single target {{target_name}}
watch get single subdomain {{subdomain}}
watch get single live {{domain}}
watch get single http {{subdomain}}

watch get subdomains domain {{domain}}
watch get subdomains scope {{scope}}
watch get subdomains all

watch get lives scope {{scope}}
watch get lives domain {{domain}}
watch get lives all

watch get http scope {{scope}}
watch get http domain {{domain}}
watch get http all

watch get latest subdomains domain {{domain}}
watch get latest subdomains scope {{scope}}
watch get latest subdomains all

watch get latest lives domain {{domain}}
watch get latest lives scope {{scope}}
watch get latest lives all

watch get targets
watch get public targets all
watch get public targets platform {{platform}}

watch get fresh
watch get statistics sqs
watch get technologies all

watch regexp list -body-file body.txt
watch regexp apply -body-file body.txt
watch regexp test scope {{scope}}
watch regexp test all

watch orch clean database scope {{scope}}
watch orch clean database all 
watch orch passive enum all 
watch orch passive enum scope {{scope}}
watch orch resolution all
watch orch resolution scope {{scope}}

watch put resolution 
watch put subdomain {{id}}
watch put orch resolution
watch put orch http

watch delete target {{target_name}}
watch create target

(use -body-file body.txt to set body)
`)
	}
	os.Exit(0)
}
