# First Go Api Mux

A tiny RestFul API is written in Go with Mux HTTP router.

## Router

```go
router := mux.NewRouter().StrictSlash(true)
router.HandleFunc("/", homePage).Methods("GET")
router.HandleFunc("/investories", getInvestories).Methods("GET")
router.HandleFunc("/investory/{uid}", getInvestory).Methods("GET")
router.HandleFunc("/investory", createInvestory).Methods("POST")
router.HandleFunc("/investory/{uid}", deleteInvestory).Methods("DELETE")
router.HandleFunc("/investory/{uid}", updateInvestory).Methods("PUT")
log.Fatal(http.ListenAndServe(":8000", router))
```

## Getting start

```
go get -u github.com/gorilla/mux
go mod init
go mod init github.com/basemax/FirstGoApiMux
go mod tidy
go run main.go
go build
```

### References

- https://www.youtube.com/watch?v=kYrdpapzcoc
- https://www.youtube.com/watch?v=wM9t9M8ohyk
- https://www.youtube.com/watch?v=eZe9jKZRvu0
- https://www.youtube.com/watch?v=V7QS42xCynY
- https://www.youtube.com/watch?v=_c1b6VFuSTk

Copyright 2021, Max Base
