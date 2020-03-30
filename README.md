# install Go

```
$ sudo apt update
$ sudo apt install golang-go
```

# install local package (only first time)

```
$ ./get-packages.sh
$ mkdir in out-enc out-dec
```

# build local

```
$ ./go-build.sh
```

# build docker

```
$ ./docker-build.sh
```

# run docker

```
$ ./docker-run.sh
```

(avec mapping in/, out-dec/ et out-enc/)

# How to test

## encode

```

curl -i -X POST -H "Content-Type: multipart/form-data" -F "file=@./enc/60.pdf" http://localhost:4450/encode

HTTP/1.1 100 Continue

HTTP/1.1 200 OK
Content-Type: application/json;charset=UTF-8
Date: Fri, 27 Mar 2020 15:34:31 GMT
Content-Length: 95

```

```yaml
{
  "name": "files/encoded/60.pdf.enc",
  "size": 176553984,
  "created_at": "2020-03-27T15:34:31.7252805Z",
}
```

## get encoded file

```
$  curl localhost:4450/files/encoded/60.pdf.enc --output enc/60.pdf.enc
```

## decode

```

curl -i -X POST -H "Content-Type: multipart/form-data" -F "file=@./enc/60.pdf.enc" http://localhost:4450/decode

HTTP/1.1 100 Continue

HTTP/1.1 200 OK
Content-Type: application/json;charset=UTF-8
Date: Fri, 27 Mar 2020 15:34:31 GMT
Content-Length: 95
```

```yaml
{
  "name": "files/decoded/60.pdf.enc.dec",
  "size": 176553984,
  "created_at": "2020-03-27T15:34:31.7252805Z",
}
```

## get decoded file

```
$  curl localhost:4450/files/decoded/60.pdf.enc.dec --output enc/60.pdf.enc.dec
```
