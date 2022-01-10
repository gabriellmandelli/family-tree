# `family-tree` - Family tree

## 🚀 About this project
The API must also provide an endpoint that returns the family tree of a certain individual containing all possible ascendants up to that level.

## 📋 Dependencies
These are some dependencies used in this repository:
- [Echo](https://echo.labstack.com/guide)
- [MongoDB](https://github.com/mongodb/mongo-go-driver)
- [Errorx](https://github.com/joomcode/errorx)
- [Viper](https://github.com/spf13/viper)

## ⚒ Requirements
- [golang 1.15+](https://golang.org/dl/)
- [VS Code](https://code.visualstudio.com/)
- [Insomnia](https://insomnia.rest/)

## ☕ Running

### Running with docker:

Start image:
```
make start-env
```
Rebuild image and run
```
make rebuild-env
```

## 🤔 How it works

### Request to persist person:
```
curl --request POST \
  --url http://localhost:8080/v1/person \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Eu"
}'
```

### Request to get person:
```
curl --request GET \
  --url http://localhost:8080/v1/person \
  --header 'Content-Type: application/json'
```

### Request to persist relationship:
```
curl --request POST \
  --url http://localhost:8080/v1/relationship \
  --header 'Content-Type: application/json' \
  --data '{
	"parent_id": "61886a0a29cb6d591f138658",
	"children_id": "61886a0429cb6d591f138654"
}'
```

### Request to get relationship:
```
curl --request GET \
  --url http://localhost:8080/v1/relationship \
  --header 'Content-Type: application/json'

```

### Request to get family tree:
```
curl --request GET \
  --url http://localhost:8080/v1/familytree/61886a0229cb6d591f138653 \
  --header 'Content-Type: application/json'
```
