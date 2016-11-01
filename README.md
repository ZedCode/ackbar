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
