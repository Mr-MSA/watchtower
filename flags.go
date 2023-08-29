package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type intelArgs struct {
	JSON            bool
	Loop            bool
	Count           bool
	CDN             bool
	Total           bool
	Limit           bool
	Internal        bool
	NoLimit         bool
	ReverseCompare  bool
	ExcludeScope    string
	ExcludeDomain   string
	ExcludeProvider string
	Provider        string
	Date            string
	Method          string
	Compare         string
	Body            string
	BodyFile        string
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
	intelFlags.StringVar(&args.Date, "date", "", "set date")
	intelFlags.StringVar(&args.ExcludeProvider, "exclude-provider", "", "exclude provider from result")
	intelFlags.StringVar(&args.ExcludeScope, "exclude-scope", "", "exclude scope from result")
	intelFlags.StringVar(&args.ExcludeDomain, "exclude-domain", "", "exclude domain from result")
	intelFlags.StringVar(&args.Method, "method", "", "http request method")
	intelFlags.StringVar(&args.Body, "body", "", "request body")
	intelFlags.StringVar(&args.BodyFile, "body-file", "", "request body file")
	intelFlags.StringVar(&args.Compare, "compare", "", "compare response")
	intelFlags.BoolVar(&args.ReverseCompare, "rc", false, "reverse compare")
	intelFlags.BoolVar(&args.JSON, "json", false, "show output as json")
	intelFlags.BoolVar(&args.Count, "count", false, "add count=true")
	intelFlags.BoolVar(&args.CDN, "cdn", false, "add cdn=true")
	intelFlags.BoolVar(&args.Internal, "internal", false, "add internal=true")
	intelFlags.BoolVar(&args.NoLimit, "no-limit", false, "add no_limit=true")
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
   --body "bodystring" (request body)
   --body-file "filename" (request body file name)
   --method string (http request method)
		
   --compare "filename" (compare response)
   --rc (reverse compare)
		  
   --limit (limit results)
   --loop (get all pages)
		
   --count (show count of results)
   --json (show output as json)
   --cdn (add cdn=ture)
   --total (add total=true)
   --internal (add internal=true)
   --no-limit (add not_limit=true)
		
   --date string (set date of results)
   --provider string (filter by providers)
		
   --exclude-domain string (exclude a domain from results)
   --exclude-provider string (exclude a provider from results)
   --exclude-scope string (exclude a scope from results)
`)
	} else {
		fmt.Printf(`watch help flags
watch version
watch init
watch update

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

watch get targets list
watch get public targets all
watch get public targets platform {{platform}}

watch get fresh subdomains all 
watch get fresh subdomains scope {{scope}} 
watch get fresh subdomains domain {{domain}} 

watch get fresh lives all 
watch get fresh lives scope {{scope}} 
watch get fresh lives domain {{domain}} 

watch get fresh http all 
watch get fresh http scope {{scope}} 
watch get fresh http domain {{domain}} 

watch get statistics sqs
watch get technologies list

watch regexp list
watch regexp apply -body-file body.txt
watch regexp test scope {{scope}} -body-file body.txt
watch regexp test all  -body-file body.txt

watch orch clean database scope {{scope}}
watch orch clean database all 
watch orch passive enum all 
watch orch passive enum scope {{scope}}
watch orch resolution all
watch orch resolution scope {{scope}}

watch put resolution 
watch put subdomain {{scope name}}
watch put orch resolution
watch put orch http

watch delete target {{target_name}}
watch create target

(use -body-file body.txt to set body)
`)
	}
	os.Exit(0)
}
