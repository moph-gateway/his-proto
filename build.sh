echo -n "version vx.x.x";
read;
protoc --go_out=./ --go-grpc_out=./ *.proto
git add .
git commit -m "${REPLY}"
git pull 
git push origin main
git tag ${REPLY}
git push origin ${REPLY}
export latest="$(git describe --tags $(git rev-list --tags --max-count=1))"
curl https://proxy.golang.org/github.com/moph-gateway/his-proto/@v/$latest.info
