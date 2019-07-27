# sa_store
store api for mysql db, xorm api cover


# usage 
```
go get github.com/theroadoffreedom/sa_store
import github.com/theroadoffreedom/sa_store
```

### sample
```
import (
	store github.com/theroadoffreedom/sa_store
)

err := store.InitStore(ip,port,username,password,dbname)
```

# test
```
go test -v -test.run TestUpdateReportId
go test -count=1 -v t_astock_operator_run_statis_test.go t_astock_operator_run_statis.go  mysql.go
```
