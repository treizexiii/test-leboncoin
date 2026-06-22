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

# RUN

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
