# qwik

This is a simple command line utility to query Wikipedia.

![Example gif](example.gif)

### Usage
```shell script
git clone https://github.com/alessandrobessi/qwik.git
cd qwik
go build -o qwik cmd/main.go
sudo mv qwik /usr/bin/ # to make it available from anywhere

qwik cpu # query "cpu" on en.wikipedia.org
qwik -lang [LANG] cpu # query "cpu" on [LANG].wikipedia.org
```
