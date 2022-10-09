### Load Balancer - Go Lang


Fire-up multiple HTTP servers (Flask Required):

```
python src/server.py "Server-Name" "Port No."
for i in {1..5}; do python src/server.py "Server-$i" "500$i" & done
```

While importing `gocron` do this after writing the import statement:

```
go mod init examplem/load-balancer 
go mod tidy
```

Get the loadbalancer online:

```
go run src/loadbalancer.go
```

Land multiple requests on the loadbalancer using `curl` command:

```
for i in {1..10}; do curl 127.0.0.1:8000; done
```