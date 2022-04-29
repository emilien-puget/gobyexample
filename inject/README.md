# Dependency injection

Some examples on how not do to and how to dependency injection

## Examples

Let's take an struct Handle that does something in database using another struct Dao

### Concrete struct

[source](./struct)

```go
type Handler struct {
    dao *Dao
}


func NewHandler(dao *Dao) *Handler {
    return &Handler{dao: dao}
}
```

| Pros                                     | Cons        |
|------------------------------------------|-------------|
| Clear identification of what is injected | Cannot mock |

The simple fact that a mock is impossible makes this solution not acceptable.

### Named interface

[source](./inter)

```go
type Handler struct {
    dao DaoI
}

type DaoI interface {
    GetSomething(ctx context.Context) (string, error)
}


func NewHandler(dao *Dao) *Handler {
    return &Handler{dao: dao}
}
```

| Pros                                                                  | Cons                            |
|-----------------------------------------------------------------------|---------------------------------|
| Good identification of injected parameters thanks to the New function | wonky named interface           |
| Can mock                                                              | needs to maintain a mock struct |

This is a way better solution than the previous one as we can mock our dependency

[daoMock.go](./inter/testdata/daoMock.go)

There is a slight issue with the named interface such as

- in which file do we put it
- What happens when multiple struct depends on it
- How do we name it

#### Mock

Here is the mock used in our case, it is very barebone and everything is up to the test to do anything with it.

[daoMock.go](./inter/testdata/daoMock.go)

```go
type DaoMock struct {
	GetSomethingFunc func(ctx context.Context) (string, error)
}

func (m DaoMock) GetSomething(ctx context.Context) (string, error) {
	return m.GetSomethingFunc(ctx)
}
```

[handler_test.go](./inter/handler_test.go)

```go
func TestHandler_handleFunc(t *testing.T) {
    h := Handler{
        dao: testdata.DaoMock{GetSomethingFunc: func(ctx context.Context) (string, error) {
            return "", nil
        }},
    }
    h.handleFunc(context.Background())
}
```

### Anonymous interface

[source](./inter_anon)

```go

type Handler struct {
    dao interface {
        GetSomething(ctx context.Context) (string, error)
    }
}

func NewHandler(dao *Dao) *Handler {
    return &Handler{dao: dao}
}
```

| Pros                                                                  | Cons                            |
|-----------------------------------------------------------------------|---------------------------------|
| Good identification of injected parameters thanks to the New function | needs to maintain a mock struct |
| Can mock                                                              |                                 |

We solve here the questions of the previous solution regarding the location of the interface

#### Mock

We use the same mock as above.

### Functions

[source](./fn)

```go
type Handler struct {
    daoGetSomething func (ctx context.Context) (string, error)
}

func NewHandler(dao *Dao) *Handler {
    return &Handler{daoGetSomething: dao.GetSomething}
}
```

| Pros                                                                  | Cons                        |
|-----------------------------------------------------------------------|-----------------------------|
| Good identification of injected parameters thanks to the New function | ownership of functions lost |
| Can mock                                                              |                             |
| no needs to maintain a mock struct                                    |                             |

#### Mock

[handler_test.go](./fn/handler_test.go)

```go
func TestHandler_handleFunc(t *testing.T) {
    h := Handler{
        daoGetSomething: func(ctx context.Context) (string, error) {
            return "", nil
        },
    }
    h.handleFunc(context.Background())
}
```

## Conclusion

The best way to handle dependency injection would be with [anonymous interface](./inter_anon) if you are willing to
maintain mocks, [functions](./fn) otherwise.
