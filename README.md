# cs-scoring

Build credit scoring microservices app with Go

Install required packages
```
go get -u github.com/gin-gonic/gin
#go get -u gopkg.in/resty.v1
```

Compile & run
```
go build
./cs-scoring
```

API Specs
```
curl -X POST -H "Accept: application/json" -H "Content-Type: application/json" -d @./json/ex-kupedes.json http://localhost:8000/las/kupedes
```
