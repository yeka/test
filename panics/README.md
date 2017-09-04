# Panics 
Simple package to catch & notify your panic or exceptions via slack or save into files.

```go
import "github.com/mataharimall/rebecca/panics"
```

## Configuration
```go
panics.SetOptions(&panics.Options{
	Env:             "TEST",
	SlackWebhookURL: "https://hooks.slack.com/services/blablabla/blablabla/blabla",
	Filepath:        "/var/log/myapplication", // it'll generate panics.log
	Channel:         "slackchannel",

	Tags: panics.Tags{"host": "127.0.0.1", "datacenter":"jkt"},
})
```

## Capture Custom Error
```go
panics.Capture(
    "Discount Anomaly",
    `{"user_id":123, "discount_amount" : 100000000}`,
)
```

## Capture Panic on HTTP Handler
```go
http.HandleFunc("/", panics.CaptureHandler(func(w http.ResponseWriter, r *http.Request) {
	panic("whatsapp bro, I'm panic now")
}))
```

## Capture Panic on httprouter handler
```go
router := httprouter.New()
router.POST("/", panics.CaptureHTTPRouterHandler(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    panic("panic is my middle name")
}))
```

## Capture Panic on negroni custom middleware
```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	panic("panic again, what's wrong with this life. may be I need to find my true love to handle this panic. ")
})
negro := negroni.New()
negro.Use(negroni.HandlerFunc(CaptureNegroniHandler))
```

## Example
### Slack Notification
![Notification Example](https://monosnap.com/file/Pjkw1uxjV8p0GnjevDwhHesUnTC2Ru.png)
