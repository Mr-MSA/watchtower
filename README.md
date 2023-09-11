# Watchtower
Go client for watchtower

## Installation
### Go install
```
GOBIN=/usr/local/bin go install github.com/Mr-MSA/watchtower@main
watchtower init
watchtower init autocompelete
```
+ set watchtower address in ~/.watch-client/.env

### Manual:
```
git clone https://github.com/Mr-MSA/watchtower
cd watchtower
go build .
./init.sh
watchtower init autocompelete
```
