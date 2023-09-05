go build . 

mkdir -p ~/.watch-client

cp structure.json ~/.watch-client
[ -d ~/.watch-client/.env ] && cp .env ~/.watch-client

sudo mv watch /usr/local/bin

echo "Execute 'watch help'"


