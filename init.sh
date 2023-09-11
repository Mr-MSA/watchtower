go build . 

mkdir -p ~/.watch-client

cp structure.json ~/.watch-client
[ -d ~/.watch-client/.env ] && cp .env ~/.watch-client

sudo mv watchtower /usr/local/bin

echo "Execute 'watchtower help'"
