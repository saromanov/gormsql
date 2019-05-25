# gormsql
[![Go Report Card](https://goreportcard.com/badge/github.com/saromanov/gormsql)](https://goreportcard.com/report/github.com/saromanov/gormsql)


Generating of the models for gorm from sql

## Usage

Create sql expression for make new table

```
CREATE TABLE IF NOT EXISTS app_user (

  id int NOT NULL,
  username varchar(45) NOT NULL,  
  password varchar(450) NOT NULL,  
  enabled integer NOT NULL DEFAULT '1',
  created_at timestamp, 
  points float,
  small bigint,
  PRIMARY KEY (id)  
);
```

Apply generation

```
go run gormsql.go --dir out app.sql
```

After generation it'll create new file at the out dir

```go

package out
type App_user struct { Enabled int "`gorm:`NOT NULL;DEFAULT:`1``"; Created_at time.Time; Points float64; Small int64; Id int "`gorm:`NOT NULL;PRIMARY_KEY;`"; Username string "`gorm:`NOT NULL;`"; Password string "`gorm:`NOT NULL;`" }
```
