export GOOS=linux
export GOARCH=amd64

go build -o api .

scp api ubuntu@work.suichu.net:eyevow
ssh ubuntu@work.suichu.net "sudo mv eyevow /usr/local/bin/"
ssh ubuntu@work.suichu.net "sudo service eyevow_api restart"

rm api