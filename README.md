# Rocket_Elevator_Rest_API_GoLang
This is the url to have access to the REST API
```
https://polar-atoll-59739.herokuapp.com
```

## Building
```diff
! This is a GET method
https://polar-atoll-59739.herokuapp.com/api/building
```

## Lead
```diff
! This is a GET method
https://polar-atoll-59739.herokuapp.com/api/lead
```

## Elevator
### GET data
```diff
! This is a GET method
https://polar-atoll-59739.herokuapp.com/api/elevator
```
return a list of inactive elevators

```diff
! This is a GET method
https://polar-atoll-59739.herokuapp.com/api/elevator/id
```
you can change id by a number to get the elevator related to it
### Update data

```diff
! This is a PUT method
https://polar-atoll-59739.herokuapp.com/api/elevator/id/status
```
you can change id by a number to get the elevator related to it and status by one of these statuts (active,inactive,intervention)
Another solution you can use is:
```diff
! This is a PUT method
https://polar-atoll-59739.herokuapp.com/api/elevator/id/
BODY:
{"status": "active"}
```
you can change the active status by the status you want
## Column
### GET data

```diff
! This is a GET method
https://polar-atoll-59739.herokuapp.com/api/column/id
```
you can change id by a number to get the column related to it
### Update data

```diff
! This is a PUT method
https://polar-atoll-59739.herokuapp.com/api/column/id/status
```
you can change id by a number to get the column related to it and status by one of these statuts (active,inactive,intervention)
Another solution you can use is:
```diff
! This is a PUT method
https://polar-atoll-59739.herokuapp.com/api/column/id/
BODY:
{"status": "active"}
```
you can change the active status by the status you want
## Battery
### GET data

```diff
! This is a GET method
https://polar-atoll-59739.herokuapp.com/api/battery/id
```
you can change id by a number to get the battery related to it
### Update data

```diff
! This is a PUT method
https://polar-atoll-59739.herokuapp.com/api/column/id/status
```
you can change id by a number to get the battery related to it and status by one of these statuts (active,inactive,intervention)
Another solution you can use is:
```diff
! This is a PUT method
https://polar-atoll-59739.herokuapp.com/api/battery/id/
BODY:
{"status": "active"}
```
you can change the active status by the status you want


-------------------BONUS----------------
## Building details
### GET data

```diff
! This is a GET method
https://polar-atoll-59739.herokuapp.com/api/buildingdetails/id
```

## Tech Phone
### Update data

```diff
! This is a PUT method
https://polar-atoll-59739.herokuapp.com/api/building/id/phone
```
