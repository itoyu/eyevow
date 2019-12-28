export GOOS=linux
export GOARCH=amd64

TARGET=stg

if test $1; then
  TARGET=$1
fi

go build -o api .

if test $TARGET = "prod"; then
  scp api ubuntu@suichu.net:eyevow
  ssh ubuntu@suichu.net "sudo mv eyevow /usr/local/bin/"
  ssh ubuntu@suichu.net "sudo service eyevow restart"
else
  scp api ubuntu@work.suichu.net:eyevow
  ssh ubuntu@work.suichu.net "sudo mv eyevow /usr/local/bin/"
  ssh ubuntu@work.suichu.net "sudo service eyevow_api restart"
fi

rm api