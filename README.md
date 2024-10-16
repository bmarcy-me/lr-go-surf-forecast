# LR GO SURF FORECAST

Find the best surf spots for next 7 days around La Rochelle, France.\
Backend implementation in Go.\
This project is a work in progress.

## Prerequisites
- Go 1.23 installed

## Start
```
go run cmd/main.go
```

## API endpoints

### /spots
/spots return the forecast for surf spots around La Rochelle

Available query parameters :
- start=2024-10-12T08:00:00Z (iso dateTime between 11/10/2024 and 20/10/2024 because we use static data for now)
- duration=2 (from 1 to 7)

```sh
curl -X GET http://localhost:8080/api/spots/start=2024-10-12T08:00:00Z&duration=2
```

The response contains each surf spots and the rating by hour with a score from 0 to 5

```json
{
    "spots": [
        {
            "id": "1",
            "name": "Plage de Gros Joncs - Ile de Ré",
            "ratings": [
                {
                    "rating": 2.221791666666667,
                    "time": "2024-10-12T09:00:00Z"
                },
                {
                    "rating": 2.3784027777777776,
                    "time": "2024-10-12T10:00:00Z"
                },
                ...
            ]
        },
        {
            "id": "2",
            "name": "Pointe du Lizay - Ile de Ré",
            "ratings": [
                {
                    "rating": 0.6341527777777778,
                    "time": "2024-10-12T09:00:00Z"
                },
                {
                    "rating": 0.9013472222222221,
                    "time": "2024-10-12T10:00:00Z"
                },
                ...
            ]
        },
        ...     
    ]
}
```

### /spots/best
/spots/best return the best surf spot around La Rochelle and the best time to go there in the next X days from a start date

Available query parameters :
- start=2024-10-17T08:00:00Z (iso dateTime between 11/10/2024 and 20/10/2024 because we use static data for now)
- duration=4 (from 1 to 7)

```sh
curl -X GET http://localhost:8080/api/spots/start=2024-10-17T08:00:00Z&duration=4
```

The response contains only one surf spot. The one with the best rating and the best time to go there.

```json
{
    "id": "1",
    "name": "Plage de Gros Joncs - Ile de Ré",
    "ratings": [
        {
            "rating": 4.609416666666666,
            "time": "2024-10-20T19:00:00Z"
        }
    ]
}
```

## Available surf spots
3 spots available for now 

|  Id   | Name                              |
| ----- | --------------------------------- |
| 1     | Plage de Gros Joncs - Ile de Ré   |
| 2     | Pointe du Lizay - Ile de Ré       |
| 3     | Plage de Vert Bois - Ile d'Oléron |



## To do list
- [ x ] api endpoint returning the best surf spot and the best time to go there
- [ ] querying stormglass at startup and store weather data in a db
- [ ] using the db instead of static json files
- [ ] docker for api server an db
- [ ] add surf spots around La Rochelle
- [ ] add tests