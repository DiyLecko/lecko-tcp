# lecko-tcp
Simple TCP Server with goRecycleBuffer

## What is this?
Very Simple TCP Server with [DiyLecko/goRecycleBuffer](https://github.com/DiyLecko/goRecycleBuffer)

## Installation
`git clone https://github.com/DiyLecko/lecko-tcp` and `cd lecko-tcp & go build main.go` and run.

## How to use?
Arbiter will receive Packets and handle at `handleRequest()` in `arbiter/arbiter.go` file.

## Why create this?
I created [DiyLecko/goRecycleBuffer](https://github.com/DiyLecko/goRecycleBuffer) for better using Buffer. So maybe this is a kind of example for goRecycleBuffer. goRecycleBuffer recycle buffer for high performance and effective memory usage in your program.

If you want to make simple TCP Server, then use this project!

clone -> add your packet handling in `handleRequest()` -> run -> done!

Thanks.

## TODO
1. Add Packet Handler with Google Protobuf
2. Add Client Example
