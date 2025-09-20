# Mesh Framework

A lightweight Go framework for composing and orchestrating independent _features_ that communicate through an event bus. The system is designed to make modular services easy to initialize, run, update, and interconnect, without hard-wiring dependencies.

---

## Overview

The **Mesh** framework lets you build applications out of pluggable _features_. Each feature implements a common interface, can be dynamically added/removed, and hooks into a shared event bus.  

Key ideas:
- **Features** are modular units of behavior (services, workers, plugins).
- **Events** provide loose coupling between features.
- **Lifecycle hooks** (init, ready, update) allow features to coordinate their startup and execution phases.
- **Concurrency** is handled internally so features can run and update independently.

---

## Core Concepts

### Mesh
The central manager that:
- Holds all registered features.
- Provides an event bus (`EventEmitter`) for inter-feature communication.
- Orchestrates feature initialization and update loops:contentReference[oaicite:0]{index=0}.

### Features
A `Feature` is any component that implements:
```go
type Feature interface {
    Name() string
    Init(mesh Mesh) error
}
```

Extensions to the base interface provide richer lifecycle behaviors:contentReference[oaicite:1]{index=1}:
- **InitReadyFeature**: signals readiness via a channel before initialization.
- **UpdatingFeature**: receives periodic `Update(elapse time.Duration)` calls.

### Events
Mesh emits lifecycle events when features are added, removed, initialized, or encounter initialization errors:contentReference[oaicite:2]{index=2}:
- `feature added`
- `feature removed`
- `feature init`
- `feature init error`

Features may also subscribe to these events by implementing handler interfaces (e.g., `EventHandlerMeshFeatureAdded`).

---

## Design Rationale

This framework was designed to solve common problems in modular applications:
1. **Loose coupling** – features don’t depend directly on each other; they discover and react through events.
2. **Composability** – features can be added/removed dynamically at runtime.
3. **Concurrency management** – initialization and update loops run safely in goroutines, coordinated by sync primitives.
4. **Lifecycle awareness** – features can defer readiness, signal initialization success/failure, and update regularly.

This makes it useful for:
- Plugin systems
- Game engines or simulations
- Service-oriented backends where modules evolve independently

---

## Example Usage

```go
package main

import (
    "fmt"
    "time"
    "myproject/mesh"
)

type Logger struct{}

func (l *Logger) Name() string { return "logger" }
func (l *Logger) Init(m mesh.Mesh) error {
    fmt.Println("Logger initialized")
    return nil
}

func (l *Logger) Update(elapse time.Duration) {
    fmt.Println("Logger tick:", elapse)
}

func main() {
    m := mesh.New()
    m.AddFeature(&Logger{})

    go func() {
        if err := m.Run(); err != nil {
            panic(err)
        }
    }()

    time.Sleep(5 * time.Second)
    m.Stop()
}
```

This simple mesh runs a `Logger` feature that prints a message every update cycle.

---

## Project Layout

- **`mesh.go`** – core `Mesh` implementation, lifecycle management, event bus orchestration:contentReference[oaicite:3]{index=3}.
- **`feature.go`** – defines the `Feature` interfaces and extensions:contentReference[oaicite:4]{index=4}.
- **`events.go`** – event constants and handler interfaces for lifecycle hooks:contentReference[oaicite:5]{index=5}.

---

## Getting Started

### Installation
```bash
go get github.com/yourname/mesh
```

### Requirements
- Go 1.20+
- [gravestench/eventemitter](https://github.com/gravestench/eventemitter)

---

## Development

### Run Tests
```bash
go test ./...
```

### Contributing
Contributions are welcome! Please open an issue or PR for new features, bugfixes, or improvements.

---

## License
MIT License – feel free to use in commercial or personal projects.
