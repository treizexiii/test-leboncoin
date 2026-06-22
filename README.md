# Todo

- endpoint 1:
  - method GET
  - name 'fizzbuzz'
  - parameters: num1(int), num2 (int), limit (int), word1 (str), word2. (str)
  - return: str list
- endpoint 2:
  - method: GET
  - name: 'stats'
  - return: most reccurent requet

# Config

Rewrite your own config file or override them with env variable (use APP_ as a prefix)

# Run

Build and run with docker

```bash
docker compose up -d
````

build and run without docker
```bash
go build -o dist/main .
cp config.yml dist
./dist/main
```
# Endpoints

### `api/fizzbuzz` POST

Request:
```json
{
    "num1": 2,
    "num2": 5,
    "limit": 15,
    "str1": "Fizz",
    "str2": "Buzz"
}
```

Response:
```json
{
  "result": [
    "1",
    "Fizz",
    "3",
    "Fizz",
    "Buzz",
    "Fizz",
    "7",
    "Fizz",
    "9",
    "FizzBuzz",
    "11",
    "Fizz",
    "13",
    "Fizz",
    "Buzz"
  ]
}
```

### `api/stats` GET

Response

```json
{
  "count": 7,
  "top_requests": {
    "num1": 2,
    "num2": 5,
    "limit": 15,
    "str1": "Fizz",
    "str2": "Buzz"
  }
}
```