docker service create --name=projectservice --replicas=1 --network=bec_network -p=5050:5050 localhost:5000/projectservice

docker service create --name=viz --publish=8080:8080/tcp --constraint=node.role==manager --mount=type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock dockersamples/visualizer