# DDD project in Go

This web application called Finder returns a match person according to a given location (x, y).

## Util requests

* Add a person

```
$ curl -i -X POST http://localhost:8080/persons -H "Accept: application/json" -H "Content-Type: application/json" -d '{"name":"MAURICIO"}'

```

* Find a match

```
$ curl -i -X GET http://localhost:8080/persons-match -H "Accept: application/json" -H "Content-Type: application/json" -d '{"x":10, "y":20}'

```

## Project hierarchy

application/  
| finder.go <--- Singleton  
domain/  
| entity/  
| | person.go  
| repository/  
| | person_repository.go <--- Singleton  
| service/  
| | matching_service.go <--- Singleton  
| value/  
| | location.go  
infrastructure/  
| main.go  
| persistence/  
| | person_repository.go  
interface/  
| web/  
| | controller.go  
| | router.go  
