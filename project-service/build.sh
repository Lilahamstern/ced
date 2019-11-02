set GOOS=linux
go build -o projectservice-linux-amd64
set GOOS=windows

docker build -t bec/projectservice .

docker service create --name=projectservice --replicas=1 --network=bec_overlay -p=5050:5050 bec/projectservice