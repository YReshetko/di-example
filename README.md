### Installation

```
go get -u github.com/YReshetko/di-example
```

### Run first plugin

```
di-example int1
```

### Run second plugin

```
di-example int2
```

###How to extend

1. Create new integration folder
2. Create go file for assertion
3. Create new implementation for registry/Assertion
4. Use init function to inject new assertion
```go
func init() {
	registry.AddIntegrations("int_new", &newAssertion{})
}
```
5. Add new package to `imports.go`