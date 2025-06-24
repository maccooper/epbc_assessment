#### Deployment Instructions

clone the repository

```
git clone https://github.com/maccooper/epbc_assessment
cd epbc_assessment (chdir on windows)

go run ./src/server.go ./src/handler.go ./src/digitOccurrence.go
```

open your browser, or use curl to visit http://localhost:8080/digit-occurrence?seriesStart=1&seriesEnd=11&seriesIncrement=1&specifiedDigit=1&seriesType=1

modify <1> value's within url to change api request to the endpoint

