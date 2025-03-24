# fahrul4215 golang boilerplate

## Installation

### Frameworks and libraries:

```bash
go get github.com/gin-gonic/gin #HTTP framework
go get github.com/jackc/pgx/v5 #PostgreSQL driver
go get gorm.io/gorm #ORM
go get gorm.io/driver/postgres #GORM PostgreSQL driver
go get github.com/spf13/viper #Configuration management
go get github.com/go-playground/validator/v10 #Data validation
go get github.com/go-redis/redis/v8 #Redis client
go get github.com/sirupsen/logrus #Logging
go get github.com/golang-jwt/jwt/v4 #JWT
```

### Monitoring and observability:

```bash
go get github.com/prometheus/client_golang/prometheus #Prometheus
go get github.com/uber-go/zap #Alternative to logrus [not working]
```

### Testing and Code Quality:

```bash
go get github.com/stretchr/testify/assert #Testing [not working]
go get github.com/axw/gocov/gocov@latest #Code coverage
go get github.com/modocache/gover@latest #Coverage Merging
go get github.com/golangci/golangci-lint/cmd/golangci-lint@latest #Linter
go get github.com/fzipp/gocyclo/cmd/gocyclo@latest #Cyclomatic complexity
```

### Documentation:

```bash
go get github.com/swaggo/swag #Swagger
```

### Database:

#### Installing postgresql:

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

##### Optional

Installing postgis:
```bash
sudo apt install postgis
```

Enable postgis extension:
```sql
CREATE EXTENSION IF NOT EXISTS postgis;
```

##### Additional:

```bash
go get github.com/jinzhu/copier
```

## Usage

```bash
go-studi-kasus-kredit-plus
```