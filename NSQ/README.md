# [**NSQ**](https://nsq.io/) - Realtime distributed messaging platform

Docker compose file how to start NSQ
It will start 3 containers:
* NSQD
  * nsqd is the daemon that receives, queues, and delivers messages to clients
  * listen on 4150(clients) and 4151(http API) ports
* NSQLOOKUPD
  * is the daemon that manages topology information. Clients query nsqlookupd to discover nsqd producers for a specific topic and nsqd nodes broadcasts topic and channel information
  * listen on 4160 and 4161 ports
* NSQADMIN
  * is a Web UI to view aggregated cluster stats in realtime and perform various administrative tasks
  * can be reached on port 4171
