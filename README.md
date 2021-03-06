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

# create a module 

- again, used go mod init packagename/modulename to bootstrap it
- then created go file with main function 

# call code from another module 

- use import block to do this
- but cannot find the package
- special instruction to link a local package
```console
go mod edit -replace example.com/greetings=../greetings
```

# error handling 

- import the "errors" package
- this creates an error 
```go
errors.New("error message") 
```
- logger uses "log" package 
- use this to log errors 
```go
log.Fatal(errorObj)
```