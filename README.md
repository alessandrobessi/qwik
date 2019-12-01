# qwik

This is a simple command line utility to query Wikipedia.

![Example gif](example.gif)

### Usage
```shell script
go build -o qwik cmd/main.go
sudo mv qwik /usr/bin/ # to make it available from anywhere

qwik cpu # query en.wikipedia.org
qwik -lang [LANG] # query [LANG].wikipedia.org
```
