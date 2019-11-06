docker build -t bec/projectservice .

docker tag bec/projectservice hamsterapps.net:5000/projectservice:latest

docker push hamsterapps.net:5000/projectservice:latest