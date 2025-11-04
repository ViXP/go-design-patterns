# Go Design Patterns

This repository is my own implementation of the classic Design Patterns in Go language, introduced in the famous book
'Design Patterns: Elements of Reusable Object-Oriented Software' written by Erich Gamma, Richard Helm, Ralph Johnson and
John Vlissides.

Each example is heavily commented and I've tried to come up with the interesting and unique application area for every
pattern to make it interesting for the reader.

Each package demonstrates one pattern and includes a `Run()` function to execute a sample.

## Patterns Implemented

### Creational

- [Builder](https://pkg.go.dev/github.com/ViXP/go-design-patterns/builder)
- [Factory Method & Abstract Factory](https://pkg.go.dev/github.com/ViXP/go-design-patterns/factory)
- [Prototype](https://pkg.go.dev/github.com/ViXP/go-design-patterns/prototype)
- [Singleton](https://pkg.go.dev/github.com/ViXP/go-design-patterns/singleton)

### Structural

- [Adapter](https://pkg.go.dev/github.com/ViXP/go-design-patterns/adapter)
- [Bridge](https://pkg.go.dev/github.com/ViXP/go-design-patterns/bridge)
- [Composite](https://pkg.go.dev/github.com/ViXP/go-design-patterns/composite)
- [Decorator](https://pkg.go.dev/github.com/ViXP/go-design-patterns/decorator)
- [Facade](https://pkg.go.dev/github.com/ViXP/go-design-patterns/facade)
- [Flyweight](https://pkg.go.dev/github.com/ViXP/go-design-patterns/flyweight)
- [Proxy](https://pkg.go.dev/github.com/ViXP/go-design-patterns/proxy)

### Behavioral

- [Chain of Responsibility](https://pkg.go.dev/github.com/ViXP/go-design-patterns/chain_of_responsibility)
- [Command](https://pkg.go.dev/github.com/ViXP/go-design-patterns/command)
- [Interpreter](https://pkg.go.dev/github.com/ViXP/go-design-patterns/interpreter)
- [Iterator](https://pkg.go.dev/github.com/ViXP/go-design-patterns/iterator)
- [Mediator](https://pkg.go.dev/github.com/ViXP/go-design-patterns/mediator)
- [Memento](https://pkg.go.dev/github.com/ViXP/go-design-patterns/memento)
- [Observer](https://pkg.go.dev/github.com/ViXP/go-design-patterns/observer)
- [State](https://pkg.go.dev/github.com/ViXP/go-design-patterns/state)
- [Strategy](https://pkg.go.dev/github.com/ViXP/go-design-patterns/strategy)
- [Template Method](https://pkg.go.dev/github.com/ViXP/go-design-patterns/template_method)
- [Visitor](https://pkg.go.dev/github.com/ViXP/go-design-patterns/visitor)

### Tactical

- [Specification](https://pkg.go.dev/github.com/ViXP/go-design-patterns/specification)

## Running

```bash
git clone https://github.com/ViXP/go-design-patterns
cd go-design-patterns
go run .
```

Please, uncomment the lines in `main.go` file to execute the example of the specific design pattern you like.

## Documentation

Full GoDoc available at:
[https://pkg.go.dev/github.com/ViXP/go-design-patterns](https://pkg.go.dev/github.com/ViXP/go-design-patterns)

## License

MIT License © ViXP
