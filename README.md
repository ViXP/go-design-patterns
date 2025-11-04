# Go Design Patterns

This repository is my own implementation of the classic Design Patterns in Go language, introduced in the famous book
'Design Patterns: Elements of Reusable Object-Oriented Software' written by Erich Gamma, Richard Helm, Ralph Johnson and
John Vlissides.

Each example is heavily commented and I've tried to come up with the interesting and unique application area for every
pattern to make it interesting for the reader.

Each package demonstrates one pattern and includes a `Run()` function to execute a sample.

## Patterns Implemented

### Creational

- Builder
- Factory Method & Abstract Factory
- Prototype
- Singleton

### Structural

- Adapter
- Bridge
- Composite
- Decorator
- Facade
- Flyweight
- Proxy

### Behavioral

- Chain of Responsibility
- Command
- Interpreter
- Iterator
- Mediator
- Memento
- Observer
- State
- Strategy
- Template Method
- Visitor

### Tactical

- Specification

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
