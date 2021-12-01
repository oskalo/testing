* ***simple benchmark***
```bash
go test -bench=.
```

* ***benchmark with memory***
```bash
go test -bench=. -benchmem
```

* ***save benchmark result***
```bash
 go test -bench=. -benchmem calculate_test.go > new.txt
```
