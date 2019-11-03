docker build -t bec/projectservice .

docker tag bec/projectservice 192.168.1.70:5000/projectservice:latest

docker push 192.168.1.70:5000/projectservice:latest