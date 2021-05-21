> ## Gland feedback
> It's lovely to hear your suggestion，share your thoughts hear: https://github.com/sinngae/gland/issues
>
> **Thank you!**


## Install
```console
go get github.com/sinngae/gland
``` 

## What is Gland
Gland is a solution for Go RPC & Web Service common packages. It has:
+ gland
    + error handling
    + error tracing
    + return code handling

Gland want to be thought of as rich packages for Go Service utility.

## Binding error With Values
```go
    err := fmt.Errorf("this is cause")

    err = gland.WithMessage(err, "this is msg")
    // err.Error() -> {"message":"this is msg", "err":"this is cause"}

    err = gland.RetCode(err, 404)
    // err.Error() -> {"retcode":404, "err":"this is cause"}

    rcMsg := NewRetCodeMsg(404, "this is msg")
    err = gland.WithRetCodeMessage(err, rcMsg)
    // err = gland.WithRetCodeMessagef(err, 404, "this is msg")
    // err.Error() -> {"retcode":404, "message":"this is msg", "err":"this is cause"}

    err = gland.WithStack(err) 
    /* err.Error() will be:
        {"stack":
        "/Users/ziqiangren/gitwork/gland/internal/stack.go:19 (0x10fb1fd)
        /Users/ziqiangren/gitwork/gland/stack.go:22 (0x10fb1f1)
        /Users/ziqiangren/gitwork/gland/stack.go:17 (0x10fb133)
        /Users/ziqiangren/gitwork/gland/stack_test.go:9 (0x10fc2d1)
        /usr/local/Cellar/go@1.13/1.13.15/libexec/src/runtime/proc.go:5228 (0x103a189)
        /usr/local/Cellar/go@1.13/1.13.15/libexec/src/runtime/proc.go:5223 (0x103a156)
        /usr/local/Cellar/go@1.13/1.13.15/libexec/src/runtime/proc.go:190 (0x102df61)
        /usr/local/Cellar/go@1.13/1.13.15/libexec/src/runtime/asm_amd64.s:1357 (0x1059610)
        ", "sum":"d172cdbeb85c21276423d3ede5dfd046", "err":"this is cause"}
    */
```

## Roadmap
+ 0.1.5 add stack & hash sum

## License
MIT