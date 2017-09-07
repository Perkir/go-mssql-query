# Go MSSQL Query

Simple MSSQL query tool that queries an MSSQL database and outputs the query delimited with '|'

## Usage

```
gomssql <user> <password> <server> /path/to/mssql
```

#### Example

```
./gomssql sa F00bar123 172.31.54.239 ./inventory.sql
``` 
Where ```inventory.sql``` has something similar to: 

```
Select * From Inventory.dbo.Flavors
```

## Building
