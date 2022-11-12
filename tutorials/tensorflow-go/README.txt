use lima to run docker commands 

run docker on macos without docker engine 

# run docker command on macos without docker engine
which lima 
lima 
limactl stop default 
limactl start default
alias docker='lima nerdctl'

# after run docker command normally
https://www.tensorflow.org/install/docker


# macpro intel cpu 
docker pull wamuir/golang-tf
docker run -it wamuir/golang-tf -- bash
cat /etc/*issue*
Debian GNU/Linux 11 \n \l
apt-get update 
apt-get install unzip
mkdir -p /model
curl -o /model/inception5h.zip http://download.tensorflow.org/models/inception5h.zip 
unzip /model/inception5h.zip -d /model 
cd /go/src
mkdir imgrecognition 
cat <<eof > main.go
package main

import (
        "fmt"
        tg "github.com/galeone/tfgo"
        tf "github.com/galeone/tensorflow/tensorflow/go"
)

func main() {
        root := tg.NewRoot()
        A := tg.NewTensor(root, tg.Const(root, [2][2]int32{{1, 2}, {-1, -2}}))
        x := tg.NewTensor(root, tg.Const(root, [2][1]int64{{10}, {100}}))
        b := tg.NewTensor(root, tg.Const(root, [2][1]int32{{-10}, {10}}))
        Y := A.MatMul(x.Output).Add(b.Output)
        // Please note that Y is just a pointer to A!

        // If we want to create a different node in the graph, we have to clone Y
        // or equivalently A
        Z := A.Clone()
        results := tg.Exec(root, []tf.Output{Y.Output, Z.Output}, nil, &tf.SessionOptions{})
        fmt.Println("Y: ", results[0].Value(), "Z: ", results[1].Value())
        fmt.Println("Y == A", Y == A) // ==> true
        fmt.Println("Z == A", Z == A) // ==> false
}
eof

go mod init app
go get github.com/galeone/tensorflow/tensorflow/go
go get github.com/galeone/tfgo

go run main.go 

#output 
...
Y:  [[200] [-200]] Z:  [[200] [-200]]
Y == A true
Z == A false


## run with docker 
docker image ls
docker image rm tfapp -f  # remove previous build 
docker build -t tfapp .
# see below result 
# tfapp                    latest                e0f21a9ea742    3 seconds ago    linux/amd64    1.4 GiB    457.3 MiB

docker run tfapp

# output 
2022-11-12 16:44:42.440680: I tensorflow/core/platform/cpu_feature_guard.cc:193] This TensorFlow binary is optimized with oneAPI Deep Neural Network Library (oneDNN) to use the following CPU instructions in performance-critical operations:  AVX2 FMA
To enable them in other operations, rebuild TensorFlow with the appropriate compiler flags.
2022-11-12 16:44:42.443784: I tensorflow/compiler/mlir/mlir_graph_optimization_pass.cc:354] MLIR V1 optimization pass is not enabled
Y:  [[200] [-200]] Z:  [[200] [-200]]
Y == A true
Z == A false
