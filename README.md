# Super Heroes API
API that supports a complete super heroes collection.

## Requirements
- Go >= 1.16.

---

## Dependencies
- gorilla/mux

---

## Build and Run
```shell
go run ./main
```

---

## Endpoints
- Get all heroes
```shell
GET /heroes
```

- Filter heroes
    - Available: name, gender, race, hairColor, publisher
```shell
GET /heroes?name=wolverine&gender=male
```

- Get Marvel Heroes
```shell
GET /heroes/marvel
```

- Get DC Heroes
```shell
GET /heroes/dc
```
---

## Payload
```shell
{
    "Name":"Batman",
    "Gender":"Male",
    "EyeColor":"blue",
    "Race":"Human",
    "HairColor":"black",
    "Height":188,
    "Publisher":"DC Comics",
    "SkinColor":"-",
    "Alignment":"good",
    "Width":95
}

```
