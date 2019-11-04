docker service rm rabbitmq
docker service create --name=rabbitmq --replicas=1 --network=bec_network -p 1883:1883 -p 5672:5672 -p 15672:15672 hamsterapps.net:5000/rabbitmq

docker service rm projectservice


docker service create --name=viz --publish=8080:8080/tcp --constraint=node.role==manager --mount=type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock dockersamples/visualizer