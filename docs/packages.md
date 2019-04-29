# Packages

## Topics
- Exporting
- Package level Variables are Global variables
- Dependency Injection
- Circular dependencies


### Dependency Injection

Wrong way:
```
type Server struct {
  config *Config
}

func New() *Server {
  return &Server{
    config: buildMyConfigSomehow(),
  }
}
```

Right way:
```
// ---------

type Server struct {
  config        *Config
  personService *PersonService
}

func NewServer(config *Config, service *PersonService) *Server {
  return &Server{
    config:        config,
    personService: service,
  }
}

// ---------

type PersonService struct {
  config     *Config
  repository *PersonRepository
}

func NewPersonService(config *Config, repository *PersonRepository) *PersonService {
  return &PersonService{config: config, repository: repository}
}

// ---------

type PersonRepository struct {
  database *sql.DB
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
  return &PersonRepository{database: database}
}

// ---------

type Server struct {
  config *Config
}

func New(config *Config) *Server {
  return &Server{
    config: config,
  }
}

// ---------

func main() {
  config := NewConfig()

  db, err := ConnectDatabase(config)

  if err != nil {
    panic(err)
  }

  personRepository := NewPersonRepository(db)
  personService := NewPersonService(config, personRepository)
  server := NewServer(config, personService)
  server.Run()
}

```

#### DI Frameworks
- Dig
  - Runtime using Reflection
- Wire (https://blog.golang.org/wire) 
  - Compile time

### Package Structure

### Shared Libraries
```
github.com/wpengine/vault
|- README.md
|- LICENSE
|- cfg/
|- log/
|- web/
|- util/
```

### Applications

```
github.com/wpengine/ssh-gateway
README.go
|- cmd/
  |- ssh-gateway
     |- ssh-gateway.go
  |- podgen
     |- podgen.go
|- internal
  |- sessions/
  |- pods/
  |- shell/
|- vendor
```

