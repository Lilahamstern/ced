docker build -t bec/rabbitmq .

docker tag bec/rabbitmq hamsterapps.net:5000/rabbitmq:latest

docker push hamsterapps.net:5000/rabbitmq:latest