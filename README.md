# runtimebench

`runtimebench` is a tool that measures the performance of the container's low-level runtime.

## Description

There are various low-level runtime now.

* runc: using cgroups and namespaces
* gVisor: hook some systemcall for security
* kata-runtime: run with lightweight vm
* and so on..

`runtimebench` is a tool to easily evaluate low-level runtime of various implementations.

## Demo

```
$ runtimebench run --runtime runc unixbench 
```

## Requirement

* Linux
* Docker
* The low-level runtime you want to test (ex. runc)

## Usage

```
$ runtimebench run <bench-name> -r runc
$ runtimebench list
```

## Install

```
$ curl https://github.com/kontotto/runtimebench/releases/~~/runtimebench.amd64
$ chmod +x runtimebench
$ sudo mv runtimebench /usr/local/bin/runtimebench
```

## License

[MIT](https://github.com/kontotto/runtimebench/blob/master/LICENSE)

## Author

[kontotto](https://github.com/kontotto)