
Zuad
====

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/zuanet/zuad)

Zuad is the reference full node ZUA implementation written in Go (golang).

## What is Zua Network

ZUA is decentralized cryptocurrency, that is making waves in the world of BlockChain.
This innovative project is based on the Proof of Work consensus alzuaithm and uses the Blake3 hashing function to ensure the security and efficiency of its network.

## Requirements

Go 1.18 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install zuad including all dependencies:

```bash
$ git clone https://github.com/zuanet/zuad
$ cd zuad
$ [go install . ./cmd/...]
$ build.sh

```

- Zuad (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.

# postgres

```bash
$ sudo apt install postgresql postgresql-contrib
$ sudo -u postgres psql
$ CREATE ROLE ZUA WITH LOGIN ENCRYPTED PASSWORD '1';
$ CREATE DATABASE zua OWNER ZUA;
$ Quit psql with \q
```




## Getting Started

Zuad has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
$ cd ~/go/bin
$ ./zuad --utxoindex
```


## Stratum server
$ git clone https://github.com/zuadanet/zua-bridge
...

## Website
Join our website server using the following link: https://ZUA.com/

## Twitter
Join our twitter server using the following link: https://twitter.com/ZuaCurrency

## Discord
Join our discord server using the following link: https://discord.gg/YNYnNN5Pf2

## Telegram
Join our telegram server using the following link: https://t.me/zuacurrency
