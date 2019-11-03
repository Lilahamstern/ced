docker build -t bec/rabbitmq .

docker tag bec/rabbitmq 192.168.1.70:5000/rabbitmq:latest

docker push 192.168.1.70:5000/rabbitmq:latest