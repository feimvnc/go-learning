go build
./app

open url http://localhost:8888

alias docker='lima nerdctl'
docker build -t app .
docker image ls 

alias docker='lima nerdctl'
docker run -p 8888:8888 app 
listening ... http://localhost:8888

curl http://localhost:8888
Hello, world

## check lima vm status and start 
limactl list
limactl stop default
limactl start default
limactl list

docker run -p 8888:8888 app
# open another terminal and run curl 
curl http://localhost:8888
Hello, world

# if below error, stop and start lima default vm, start container again 
# FATA[0000] failed to create shim task: 
# OCI runtime create failed: runc create failed

lima stop default 
lima start default 


