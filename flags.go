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
	Tag             string
	PublicTarget    string
	ExcludeScope    string
	Status          string
	Title           string
	ExcludeDomain   string
	ExcludeProvider string
	Provider        string
	Date            string
	Method          string
	Compare         string
	Body            string
	BodyFile        string
	ResponseHeaders string
	Technology      string
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

	intelFlags.StringVar(&args.ResponseHeaders, "response-headers", "", "")
	intelFlags.StringVar(&args.Technology, "tech", "", "")
	intelFlags.StringVar(&args.Tag, "tag", "", "")
	intelFlags.StringVar(&args.Provider, "provider", "", "set providers")
	intelFlags.StringVar(&args.Status, "status", "", "match status")
	intelFlags.StringVar(&args.Title, "title", "", "match title")
	intelFlags.StringVar(&args.Date, "date", "", "set date")
	intelFlags.StringVar(&args.ExcludeProvider, "exclude-provider", "", "exclude provider from result")
	intelFlags.StringVar(&args.ExcludeScope, "exclude-scope", "", "exclude scope from result")
	intelFlags.StringVar(&args.ExcludeDomain, "exclude-domain", "", "exclude domain from result")
	intelFlags.StringVar(&args.Method, "method", "", "http request method")
	intelFlags.StringVar(&args.Body, "body", "", "request body")
	intelFlags.StringVar(&args.PublicTarget, "public-target", "", "add public target by name")
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

	if len(args) == 2 && args[1] == "version" {
		fmt.Println(version)
		os.Exit(0)
	} else if len(args) == 2 && args[1] == "flags" {
		fmt.Printf(`Flags:
   --body "bodystring" (request body)
   --body-file "filename" (request body file name)
   --public-target "program_name" (add public target by name)
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
   --status string (filter by status)
   --title string (filter by title)
   --tech (filter by technology)
		
   --exclude-domain string (exclude a domain from results)
   --exclude-provider string (exclude a provider from results)
   --exclude-scope string (exclude a scope from results)
`)
	} else {
		fmt.Printf(`watchtower help flags
watchtower help version

watchtower init
watchtower init autocomplete

watchtower update

watchtower get single target {{target_name}}
watchtower get single subdomain {{subdomain}}
watchtower get single live {{domain}}
watchtower get single http {{subdomain}}

watchtower get subdomains domain {{domain}}
watchtower get subdomains scope {{scope}}
watchtower get subdomains all

watchtower get lives scope {{scope}}
watchtower get lives domain {{domain}}
watchtower get lives all

watchtower get http scope {{scope}}
watchtower get http domain {{domain}}
watchtower get http all

watchtower get latest subdomains domain {{domain}}
watchtower get latest subdomains scope {{scope}}
watchtower get latest subdomains all

watchtower get latest lives domain {{domain}}
watchtower get latest lives scope {{scope}}
watchtower get latest lives all

watchtower get latest http domain {{domain}}
watchtower get latest http scope {{scope}}
watchtower get latest http all

watchtower get targets list
watchtower get targets public all
watchtower get targets public platform {{platform}}

watchtower get fresh subdomains all 
watchtower get fresh subdomains scope {{scope}} 
watchtower get fresh subdomains domain {{domain}} 

watchtower get fresh lives all 
watchtower get fresh lives scope {{scope}} 
watchtower get fresh lives domain {{domain}} 

watchtower get fresh http all 
watchtower get fresh http scope {{scope}} 
watchtower get fresh http domain {{domain}} 

watchtower get statistics sqs
watchtower get technologies list

watchtower regexp list
watchtower regexp apply -body-file body.txt
watchtower regexp test scope {{scope}} -body-file body.txt
watchtower regexp test all  -body-file body.txt

watchtower orch clean database scope {{scope}}
watchtower orch clean database all 
watchtower orch clean tags scope {{scope}}
watchtower orch clean tags all
watchtower orch clean ips scope {{scope}}
watchtower orch clean ips all
watchtower orch clean domains scope {{scope}}
watchtower orch clean domains all
watchtower orch passive enum all 
watchtower orch passive enum scope {{scope}}
watchtower orch resolution all
watchtower orch resolution scope {{scope}}
watchtower orch http all
watchtower orch http scope {{scope}}
watchtower orch targets update

watchtower put resolution 
watchtower put subdomain {{scope}}
watchtower put orch resolution
watchtower put orch http

watchtower target delete {{target_name}}
watchtower target create

(use -body-file body.txt to set body)
`)
	}
	os.Exit(0)
}
