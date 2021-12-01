## How to generate mock for interface

* [install mockery](https://github.com/vektra/mockery#installation)
* ***template***
```bash
mockery -dir {path were interface exists} -name {name of the interface} -output app/mocks -output mocks -case underscore
```
* ***example***
   
```bash
workshop-testing % mockery --dir 3-mocks/repository --name CommentsRepository --output 3-mocks/mocks --case underscore
```
* [mock](https://github.com/vektra/mockery#installation)