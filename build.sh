protoc --go_out=./ --go-grpc_out=./ *.proto
git pull
export latest="$(git describe --tags $(git rev-list --tags --max-count=1))"
curl https://proxy.golang.org/github.com/moph-gateway/his-proto/@v/$latest.info
