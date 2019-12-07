export GOOS=linux
export GOARCH=amd64

go build -o api .

scp api ubuntu@work.suichu.net:eyevow
ssh ubuntu@work.suichu.net "sudo mv eyevow /usr/local/bin/"

rm api