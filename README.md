# Stupid simple in memory URL shortener 

How to try:

```
go run main.go
```


to shorten the url:
```
curl -i --get 'http://localhost:8080/shorten/' \
  --data-urlencode 'url=https://example.com/articles?id=123&sort=newest'
```
to recover the url you shortened:
```
curl -i 'http://localhost:8080/sr/<token you are given>'
```



