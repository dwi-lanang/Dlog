# Dlog

Initial

```go
var (
    urlServer = "" 
    mode = "" //debug or monitor
    channel = "" //api
)
dl := dlog.Init(urlServer, dlog.Config{
    Mode:    mode,
    Channel: channel,
}, func(str string) {

})
```

Monitor Callback
```go
func main(){
    dl := dlog.Init(urlServer, dlog.Config{
        Mode:    mode,
        Channel: channel,
    }, monitorCallback)
}

func monitorCallback(data string){
    
}
```

Send Log

state string, parameter interface{}
```go
dl.Send("log1", map[string]interface{}{
    "name":"Tony",
    "hobby":"Skateboard"
})
```