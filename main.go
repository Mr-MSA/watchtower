package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const version = "1.0.3"

func main() {

	// get home dir
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// validate args
	if len(os.Args[1:]) == 0 {
		fmt.Println("Execute: watch help")
		os.Exit(0)
	}

	// get arguments
	args := dropFlags(os.Args[1:])

	// rand string
	var ver = String(10)

	// init

	if args[0] == "init" {
		if _, err := os.Stat(homedir + "/.watch-client"); os.IsNotExist(err) {
			if err := os.Mkdir(homedir+"/.watch-client/", os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}

		if len(args) == 2 && args[1] == "autocomplete" {
			if _, err := os.Stat(homedir + "/.watch-client/completions/"); os.IsNotExist(err) {
				if err := os.Mkdir(homedir+"/.watch-client/completions", os.ModePerm); err != nil {
					log.Fatal(err)
				}
			}

			if err := downloadFile(homedir+"/.watch-client/completions/_watch", "https://raw.githubusercontent.com/Mr-MSA/Watch/main/_watch?"+ver); err != nil {
				fmt.Println(err)
			}
			if err := downloadFile(homedir+"/.watch-client/init-autocomplete.sh", "https://raw.githubusercontent.com/Mr-MSA/Watch/main/init-autocomplete.sh?"+ver); err != nil {
				fmt.Println(err)
			}
			cmd := exec.Command("zsh", homedir+"/.watch-client/init-autocomplete.sh")
			_, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			}
		} else if len(args) == 1 {

			if err := downloadFile(homedir+"/.watch-client/.env", "https://raw.githubusercontent.com/Mr-MSA/Watch/main/.env"); err != nil {
				fmt.Println(err)
			}

			if err := downloadFile(homedir+"/.watch-client/structure.json", "https://raw.githubusercontent.com/Mr-MSA/Watch/main/structure.json?"+ver); err != nil {
				fmt.Println(err)
			}
		}
		os.Exit(0)

	} else if args[0] == "update" {
		if err := downloadFile(homedir+"/.watch-client/structure.json", "https://raw.githubusercontent.com/Mr-MSA/Watch/main/structure.json?"+ver); err != nil {
			fmt.Println(err)
		}
		fmt.Println("structure.json updated!")
		os.Exit(0)
	} else {
		// check if config dir exists
		if _, err := os.Stat(homedir + "/.watch-client"); os.IsNotExist(err) {
			fmt.Println("Path " + homedir + "/.watch-client not found! please execute 'watch init' ")
			os.Exit(0)
		}
	}

	// validate baseurl
	if envVariable("baseURL") == "WATCH_SERVER" {
		fmt.Println("Please set watch server address at " + homedir + "/.watch-client/.env")
		os.Exit(0)
	}

	// read and parse configurations
	config := ReadJSON(homedir + "/.watch-client/structure.json")

	// show help
	showHelp(args, config)

	// set variables
	var api string
	var count = 1

	// find endpoint
	var conf map[string]interface{} = config
	for i := 0; i <= len(args); i++ {

		if conf[args[i]] == nil {
			break
		}

		if fmt.Sprintf("%T", conf[args[i]]) == "string" {
			api = conf[args[i]].(string)
			count += i
			break
		} else {
			conf = conf[args[i]].(map[string]interface{})
		}
	}

	// validate api
	if api == "" {
		fmt.Println("Command/API not found")
		os.Exit(0)
	}

	// check has arg
	if strings.Contains(api, "{{arg}}") {
		count++
	}

	// parse api endpoint
	api = strings.Replace(api, "{{arg}}", args[len(args)-1], -1)
	api = strings.Replace(api, "{{base}}", envVariable("baseURL"), -1)

	// parse flags
	flags := os.Args[1:][count:]
	var flagArgs intelArgs

	intelCommand := flag.NewFlagSet("main", flag.ContinueOnError)

	defineIntelArgumentFlags(intelCommand, &flagArgs)

	if err := intelCommand.Parse(flags); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(2)
	}

	// set body
	var body string
	if flagArgs.Body != "" {

		body = flagArgs.Body

	} else if flagArgs.BodyFile != "" {

		// read body file
		fileContent, err := ioutil.ReadFile(flagArgs.BodyFile)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(0)
		}

		body = string(fileContent)
	}

	// append ?
	if !strings.Contains(api, "?") {
		api = api + "?"
	}

	// set default request methods
	if flagArgs.Method == "" {
		flagArgs.Method = setMethod(args)
	}

	// set loop
	if args[0] == "get" {
		if args[1] == "lives" || args[1] == "fresh" || args[1] == "subdomains" || args[1] == "latest" || args[1] == "http" {
			flagArgs.Loop = true
		}
	}

	// set endpoint by flags

	if flagArgs.Count {
		api = fmt.Sprintf("%s&count=true", api)
	}
	if flagArgs.CDN {
		api = fmt.Sprintf("%s&cdn=true", api)
	}
	if flagArgs.Internal {
		api = fmt.Sprintf("%s&internal=true", api)
	}
	if flagArgs.NoLimit {
		api = fmt.Sprintf("%s&no_limit=true", api)
		flagArgs.Loop = false
	}
	if flagArgs.Total {
		api = fmt.Sprintf("%s&total=true", api)
	}
	if flagArgs.JSON {
		api = fmt.Sprintf("%s&json=true", api)
		flagArgs.Loop = false
	}
	if flagArgs.Provider != "" {
		api = fmt.Sprintf("%s&provider=%s", api, flagArgs.Provider)
	}
	if flagArgs.Title != "" {
		api = fmt.Sprintf("%s&title=%s", api, flagArgs.Title)
	}
	if flagArgs.Status != "" {
		api = fmt.Sprintf("%s&status=%s", api, flagArgs.Status)
	}
	if flagArgs.Date != "" {
		api = fmt.Sprintf("%s&date=%s", api, flagArgs.Date)
	}
	if flagArgs.ExcludeDomain != "" {
		api = fmt.Sprintf("%s&exclude_domain=%s", api, flagArgs.ExcludeDomain)
	}
	if flagArgs.ExcludeScope != "" {
		api = fmt.Sprintf("%s&exclude_scope=%s", api, flagArgs.ExcludeScope)
	}
	if flagArgs.ExcludeProvider != "" {
		api = fmt.Sprintf("%s&exclude_provider=%s", api, flagArgs.ExcludeProvider)
	}

	// limit res
	if flagArgs.Limit {
		flagArgs.Loop = false
	}

	var out string
	if flagArgs.Loop && !flagArgs.Count {

		// get count of results
		count, err := strconv.Atoi(MakeHttpRequest(api+"&count=true", flagArgs, body))
		if err != nil {
			fmt.Println("Can't convert string to integer")
		}

		// get results
		for i := 0; i <= ((count / 1000) + 1); i++ {
			resp := MakeHttpRequest(api+"&limit=1000&page="+strconv.Itoa(i), flagArgs, body)

			if flagArgs.Compare == "" {
				fmt.Print(resp)
			} else {
				out += resp

			}
		}
		fmt.Printf("\n")

	} else {

		// send http request to api endpoint
		resp := MakeHttpRequest(api, flagArgs, body)
		if flagArgs.Compare == "" {
			fmt.Print(resp)
		} else {
			out = resp
		}
		fmt.Printf("\n")
	}

	if flagArgs.Compare != "" {

		f, err := os.Create(".tmp")
		if err != nil {
			log.Fatal(err)
		}

		_, err2 := f.WriteString(out)
		if err2 != nil {
			log.Fatal(err2)
		}
		f.Close()

		var f1, f2 string
		f1 = ".tmp"
		f2 = flagArgs.Compare

		if flagArgs.ReverseCompare {
			f1 = flagArgs.Compare
			f2 = ".tmp"
		}

		if _, err := os.Stat(flagArgs.Compare); err != nil {
			fmt.Println("Compare file does not exist!")
			return
		}

		cmd := exec.Command("grep", "-Fxvf", f1, f2)
		stdout, _ := cmd.Output()

		fmt.Print(string(stdout))
	}

}
