# golangpractice

practice repository based on [this link](https://go.dev/doc/tutorial/getting-started#prerequisites)

# installation 

1.  download installer for mac 
2.  open the package, it will install to /usr/local/go 
3.  updates PATH environment variable 
4.  may need to restart terminal 
5.  test with go version 

# dependency tracking

- go.mod defines a module and tracks packages 
- created with cli
```console
go mod init example/hello
```

# call code in external package 

package repository can be found at [this link](pkg.go.dev)

- this seems to automatically parse code for necessary requirements and imports the packages 
```console
go mod tidy
```

- has a timeout
- try to fix with a proxy 
```console
http_proxy=127.0.0.1:4780 https_proxy=127.0.0.1:4780 go mod tidy 
```

- the environment variables are necessary when running, too.