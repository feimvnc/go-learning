source: https://www.youtube.com/watch?v=PAUjYyBfELk
source: https://github.com/jboursiquot/portscan

use net.Dial to build a network scan utility 

# to test local, used below to start a docker ps

limactl start default 
alias docker='lima nerdctl'
docker ps 

# start a mongodb docker instance 
docker run -d --name mongodb -p 27017:27017 mongo


#test 

# basic net.Dial single port 27017
(base) user:basic user$ go run main.go 
2022/11/23 03:32:21 27017 open
(base) user:basic user$ 


# range scan, sequential scan 
(base) user:range-scan user$ go run main.go 
2022/11/23 03:37:08 27000 closed (dial tcp [::1]:27000: connect: connection refused)
2022/11/23 03:37:08 27001 closed (dial tcp [::1]:27001: connect: connection refused)
2022/11/23 03:37:08 27002 closed (dial tcp [::1]:27002: connect: connection refused)
2022/11/23 03:37:08 27003 closed (dial tcp [::1]:27003: connect: connection refused)
2022/11/23 03:37:08 27004 closed (dial tcp [::1]:27004: connect: connection refused)
2022/11/23 03:37:08 27005 closed (dial tcp [::1]:27005: connect: connection refused)
2022/11/23 03:37:08 27006 closed (dial tcp [::1]:27006: connect: connection refused)
2022/11/23 03:37:08 27007 closed (dial tcp [::1]:27007: connect: connection refused)
2022/11/23 03:37:08 27008 closed (dial tcp [::1]:27008: connect: connection refused)
2022/11/23 03:37:08 27009 closed (dial tcp [::1]:27009: connect: connection refused)
2022/11/23 03:37:08 27010 closed (dial tcp [::1]:27010: connect: connection refused)
2022/11/23 03:37:08 27011 closed (dial tcp [::1]:27011: connect: connection refused)
2022/11/23 03:37:08 27012 closed (dial tcp [::1]:27012: connect: connection refused)
2022/11/23 03:37:08 27013 closed (dial tcp [::1]:27013: connect: connection refused)
2022/11/23 03:37:08 27014 closed (dial tcp [::1]:27014: connect: connection refused)
2022/11/23 03:37:08 27015 closed (dial tcp [::1]:27015: connect: connection refused)
2022/11/23 03:37:08 27016 closed (dial tcp [::1]:27016: connect: connection refused)
2022/11/23 03:37:08 27017 open
2022/11/23 03:37:08 27018 closed (dial tcp [::1]:27018: connect: connection refused)
2022/11/23 03:37:08 27019 closed (dial tcp [::1]:27019: connect: connection refused)
(base) user:range-scan user$ 

# concurrent scan , because main goroutine starts and completes first
# port scan goroutines are scheduled only, but did not get change to execute 

(base) user:concurent-scan user$ go run main.go 
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done
2022/11/23 03:45:57 done


# add wait group in goroutine 
(base) user:waitgroup-scan user$ go run main.go 
2022/11/23 04:03:10 27010 closed (dial tcp 127.0.0.1:27010: connect: connection refused)
2022/11/23 04:03:10 27020 closed (dial tcp 127.0.0.1:27020: connect: connection refused)
2022/11/23 04:03:10 27014 closed (dial tcp 127.0.0.1:27014: connect: connection refused)
2022/11/23 04:03:10 27015 closed (dial tcp 127.0.0.1:27015: connect: connection refused)
2022/11/23 04:03:10 27012 closed (dial tcp 127.0.0.1:27012: connect: connection refused)
2022/11/23 04:03:10 27011 closed (dial tcp 127.0.0.1:27011: connect: connection refused)
2022/11/23 04:03:10 27018 closed (dial tcp 127.0.0.1:27018: connect: connection refused)
2022/11/23 04:03:10 27016 closed (dial tcp 127.0.0.1:27016: connect: connection refused)
2022/11/23 04:03:10 27019 closed (dial tcp 127.0.0.1:27019: connect: connection refused)
2022/11/23 04:03:10 27013 closed (dial tcp 127.0.0.1:27013: connect: connection refused)
2022/11/23 04:03:10 27017 open
2022/11/23 04:03:10 done

# include large number of ports 
# error of "socket, too many open files" may display
(base) user:waitgroup-scan user$ go run main.go --from 1000 --to 65535
...
2022/11/23 04:06:47 60110 closed (dial tcp 127.0.0.1:60110: connect: connection refused)
2022/11/23 04:06:47 60103 closed (dial tcp 127.0.0.1:60103: connect: connection refused)
2022/11/23 04:06:47 59885 closed (dial tcp 127.0.0.1:59885: connect: connection refused)
2022/11/23 04:06:51 done


# worker pool, a set of goroutines to pick up and do the work
# this can avoid system resource constraints issue above 
# worker pool helps application be stable 
(base) user:worker-pool-scan user$ go run main.go -ports 27010-27020
27010 closed (dial tcp 127.0.0.1:27010: connect: connection refused)
27014 closed (dial tcp 127.0.0.1:27014: connect: connection refused)
27013 closed (dial tcp 127.0.0.1:27013: connect: connection refused)
27019 closed (dial tcp 127.0.0.1:27019: connect: connection refused)
27012 closed (dial tcp 127.0.0.1:27012: connect: connection refused)
27015 closed (dial tcp 127.0.0.1:27015: connect: connection refused)
27016 closed (dial tcp 127.0.0.1:27016: connect: connection refused)
27020 closed (dial tcp 127.0.0.1:27020: connect: connection refused)
27011 closed (dial tcp 127.0.0.1:27011: connect: connection refused)
27018 closed (dial tcp 127.0.0.1:27018: connect: connection refused)

Results
---x
27017 - open


#semaphore-with-timeout
(base) user:semaphore-with-timeout user$ go run main.go -ports 27000-27020
27000 closed (dial tcp 127.0.0.1:27000: connect: connection refused)
27014 closed (dial tcp 127.0.0.1:27014: connect: connection refused)
27019 closed (dial tcp 127.0.0.1:27019: connect: connection refused)
27002 closed (dial tcp 127.0.0.1:27002: connect: connection refused)
27011 closed (dial tcp 127.0.0.1:27011: connect: connection refused)
27015 closed (dial tcp 127.0.0.1:27015: connect: connection refused)
27012 closed (dial tcp 127.0.0.1:27012: connect: connection refused)

Results
---
27017 - open


# pipeline 
(base) user:network-scan-concurrency user$ cd pipeline/
(base) user:pipeline user$ go run main.go 
completed, check scans.csv for results

 (base) user:pipeline user$ cat scans.csv 
port,open,scanError,scanDuration
27017,true,,246.074µs

# fan out and fan in 

(base) user:fan-out-fan-in user$ go run main.go -ports 27015-27020
0xc00002c240
0xc000024300
main.scanOp{port:27015, open:false, scanErr:"dial tcp 127.0.0.1:27015: connect: connection refused", scanDuration:259903}
main.scanOp{port:27016, open:false, scanErr:"dial tcp 127.0.0.1:27016: connect: connection refused", scanDuration:200537}
main.scanOp{port:27018, open:false, scanErr:"dial tcp 127.0.0.1:27018: connect: connection refused", scanDuration:97975}
main.scanOp{port:27020, open:false, scanErr:"dial tcp 127.0.0.1:27020: connect: connection refused", scanDuration:91425}
main.scanOp{port:27019, open:false, scanErr:"dial tcp 127.0.0.1:27019: connect: connection refused", scanDuration:174316}
main.scanOp{port:27017, open:true, scanErr:"", scanDuration:261235}

