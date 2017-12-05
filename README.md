# mocksrv

Simple and high performance mock server implementation.

## Install

```
go get github.com/mashiro/mocksrv
```

## Usage

### Routing

Root only
```
mocksrv --get "/=200"
```

Any parameter
```
mocksrv --get "/:name=200"
```

Wildcard
```
mocksrv --get "/*paths=200"
```

### Serving static files

FS
```
mocksrv --root "/assets:./assets"
```

File
```
mocksrv --file "/favicon.ico:./assets/favicon.ico"
```

