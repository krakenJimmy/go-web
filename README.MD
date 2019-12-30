# Introduce

This repository is the exercise of my road for learning go

## Features

+ signin and signup with jwt auth
+ file upload
+ etc...

## go resources I learned from

[https://www.sohamkamani.com/blog/golang](https://www.sohamkamani.com/blog/golang)

[https://github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)

[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

[https://github.com/coderminer/restful](https://github.com/coderminer/restful)

[https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body](https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body)

[https://gowebexamples.com/routes-using-gorilla-mux/](https://gowebexamples.com/routes-using-gorilla-mux/)


## website that helps me slove problem 

[set zap log](https://blog.sandipb.net/2018/05/04/using-zap-working-with-global-loggers/)

### how to log with logrus

```go
log.SetReportCaller(true)
log.WithFields(log.Fields{
        "user": fmt.Sprintf("%+v", user),
}).Info()
```