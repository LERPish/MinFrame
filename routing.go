package minframe

import (
    "fmt"
    "net/http"
    "strings"
)

type Route struct {
    Name string
    Path string
    Method string
    Handler http.HandlerFunc
    Middleware []func(http.HandlerFunc) http.HandlerFunc
}
