# go-samp-query

SA:MP Query API for Go

See [GoDoc](https://godoc.org/github.com/Southclaws/go-samp-query) for full
documentation

## Examples

Most of the time, you'll only need one function:

```go
server, err := GetServerInfo(ctx, "192.168.1.1", true)
if err != nil {
    // handle
}
```

The `attemptDecode` parameter determines whether or not the library should
attempt to guess the encoding of the text fields such as hostname etc.

If you want to get specific data about a server, you can create a query and
selectively query for data:

```go
query, err := NewLegacyQuery(host)
if err != nil {
    // handle
}
defer query.Close()

rules, err := query.GetRules(ctx)
if err != nil {
    return
}
```
