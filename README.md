# ackbar
A tool for sending random data over the network

## What is this?
This is a tool I wrote while investigating browser behavior when unpredictable content is included on a page.

## How to build?
Easy!
```
> git clone git@github.com:ZedCode/ackbar.git
> cd ackbar
> go build .
```

## How to use?
Easy!
```
./ackbar -p 8080
```
Now, in a browser, visit your local IP address and port 8080. Your browser will likely crash.

## Why would I use this?
Because this is written to send random data to any request type made, just include it on a page and see what happens? For example, what happens if you include `<script type='text/javascript' src='http://my-bad-domain-or-ip:8080/random.js'></script>` on a page? What if you use this as the source of an image on an upload form?

This really isn't meant for other people to contribute to, just a little POC for testing how well network services parse their input.

## Example Output
```
$ ./ackbar -p 8080
2016/11/01 12:00:49 Starting Ackbar...
2016/11/01 12:00:49 Starting up listener on 8080. Use ctrl-c to stop.
2016/11/01 12:01:18 IT'S A TRAP!!
2016/11/01 12:01:18 REQUEST FROM: [some IP]:[some port]
2016/11/01 12:01:18 REQUEST CONTENT:  GET /test.png HTTP/1.1
Host: [my IP]:8080
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:49.0) Gecko/20100101 Firefox/49.0
Accept: */*
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate
Referer: http://[a site IP I control]
Connection: keep-alive


2016/11/01 12:01:18 LOCKING S-FOILS IN ATTACK POSITION
2016/11/01 12:01:18 TARGET DOWN: write tcp4 [my IP]:8080->[some IP]:53887: write: broken pipe
2016/11/01 12:01:18 FIRED 2458272 BYTES BEFORE ENDPOINT COWARDLY RETREATED.
```
