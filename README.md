## Clean Architecture minimum-asset

```
.
├── domain
├── driver
├── gateway
│   └── user
│       └── convert
├── port
│   ├── response
│   └── user
├── presenter
│   └── user
│       └── convert
├── rest
│   └── user
└── usecase
    └── user
```

## Layer description

### Frameworks & Drviers

- `rest`
  - REST web api
- `driver`
  - Database persistent

### Interface Adapter

- `gateway`
  - Implementing Port
  - `convert` is domain-model <=> data-model
- `presenter`
  - Implementing Output Port

### Usecase

- `usecase`
  - calling domain-logic
  - calling port

### Domain

- `domain`
  - Domain knowledge of defined by domain expert