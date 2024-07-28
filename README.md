# Watchtower
Go client for watchtower

## Installation
### Go install
```
go install github.com/Mr-MSA/watchtower@main
watchtower init
watchtower init autocompelete
```
+ set watchtower address with scheme in `~/.watchtower/.env`

### Manual:
```
git clone https://github.com/Mr-MSA/watchtower
cd watchtower
go build .
./init.sh
watchtower init autocompelete
```
### Help
```
watchtower help
watchtower help flags
```

### Update Watch
```
<<<<<<< HEAD
go install github.com/Mr-MSA/watchtower@v1.1.0
=======
go install github.com/Mr-MSA/watchtower@latest
>>>>>>> cb427464da4859867f124e5992f93a1fb12b7fcb
watchtower update 
```

### Reinstall for new version
```bash
rm -rf ~/.watch-client
go install github.com/Mr-MSA/watchtower@latest
watchtower init 
# ---- Set IP, username and password at ~/.watchtower/.env
watchtower init autocomplete
zsh
watchtower help
```

### Active Autocomplete
Execute following commands:
```
watchtower update
watchtower init autocomplete
```
Add following commands in zshrc or zsh profile and source file:
```
fpath=(~/.watchtower/ $fpath)
autoload -Uz compinit
compinit
```
### Configuration
+ Set Username and Password in `~/watchtower/.env`

### Flags
```
Flags:
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
   --status string (filter by status)
   --title string (filter by title)

   --exclude-domain string (exclude a domain from results)
   --exclude-provider string (exclude a provider from results)
   --exclude-scope string (exclude a scope from results)
```
