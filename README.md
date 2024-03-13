# Cards API

#### This service can validate cards.

## Start the application:

1. Clone this repository (https://github.com/Taras-Rm/cards-api).
2. Install **Docker** on computer.
3. Enter into the project folder.
4. Create image `docker build -t cards-api .`
5. Run application `docker run -p 8080:8080 cards-api`

## There is 1 endpoint in the application:
- **POST** - localhost:8080/api/cards/validate (_validate card_)
#### Response (error field is optional):
  ```sh
{
    "valid": "false",
    "error": {
        "code": 400,
        "message": "invalid card number length"
    }
}
```