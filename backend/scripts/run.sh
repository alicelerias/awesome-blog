echo "Executing test"

set -e

SCRIPTS="./scripts"
BUILD_DIR="$SCRIPTS/build"
BUILD_OUTPUT="$BUILD_DIR/program"
GO_ENV=production
GOOS=linux
GOARCH=amd64
DOWNLOADS_DIR="$SCRIPTS/downloads"
mkdir -p $DOWNLOADS_DIR

echo Checking GO Lang is Installed


# set +e
# go version
# if [ $? -eq 127 ] ; then
#   echo Installing Go lang
#   mkdir -p $DOWNLOADS_DIR
#   exit 1
#   wget https://go.dev/dl/go1.20.3.linux-amd64.tar.gz -O $DOWNLOADS_DIR/go.tar.gz
#   tar -xzf $DOWNLOADS_DIR/go.tar.gz -C $DOWNLOADS_DIR
# fi

# PATH+=":$DOWNLOADS_DIR/go/bin"

set -e

echo `go version`

echo Install dependencies
go mod tidy

echo Running Format check
$SCRIPTS/format.sh

echo Running Tests
go test ./...

echo Building

mkdir -p $BUILD_DIR

go build -ldflags="-s -w" -o $BUILD_OUTPUT

chmod +x $BUILD_OUTPUT

echo Deploy

scp $BUILD_DIR/program pantheon:/home/yuriharrison/app
scp $BUILD_DIR/.production.env pantheon:/home/yuriharrison/app/.env


ssh pantheon 'cd /home/yuriharrison/app && . .env && ./program'

set +e

echo Done!
