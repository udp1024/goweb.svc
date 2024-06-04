# goweb.svc
a simple REST web service written in go using gin router to support crud operations on json/data.json file

## Purpose
The purpose of this code repo is to provide updated Go language example for a web-service. 
As of **'2024/05'** many examples on the web are outdated because they use deprecated modules.

## Data Structure
The data resides in the json/ directory, in file named data.jspon

### Sample data
```
[
  {
    "id": "1",
    "icon": "images/whatsapp-transp.png",
    "name": "Whats App",
    "description": "messaging",
    "link": "https://web.whatsapp.com"
  }
]
```

## Exposed Methods
The web service exposes following methods;
```
1. GET "/cards" - returns all cards in the file
2. GET "/cards/id" - returns card with specified "id" - e.g. get /cards/2
3. POST "/cards" - adds a new card to the json/data.json file. Returns card data with id number set.
4. PUT "cards/id" - updates the card identified by the provided id with the provided data. returns updated data
5. DELETE "/cards/id" deletes the card specified with the id.
```

## Testing
Example curl commands to use some of the exposed interfaces are;
```
# return all cards
curl http://localhost:8080/cards --include --header "Content-Type: application/json" --request "GET"

# return card 5
curl http://localhost:8080/cards/5 --include --header "Content-Type: application/json" --request "GET"

# Create a new card
curl http://localhost:8080/cards --include --header "Content-Type: application/json" --request "POST" --data '{"id":"","icon":"images/webmin.png","name":"vmsvr1 webmin","description":"Vmsvr1 webmin again","link":"https://vmsvr1.udp1024.com:10000"}'

# update card 5
curl http://localhost:8080/cards/5 --include --header "Content-Type: application/json" --request "PUT" --data '{"id": "5","icon": "images/google_drive.png","name": "Google Drive","description": "G Drive Storage","link": "https://drive.google.com"}'

# update card 10
curl http://localhost:8080/cards/10 --include --header "Content-Type: application/json" --request "PUT" --data '{"id":"10","icon":"images/webmin.png","name":"vmsvr1 webmin","description":"Vmsvr1 webmin2","link":"https://vmsvr1.udp1024.com:10000"}'

# Delete a record
curl http://localhost:8080/cards/10 --request "DELETE"

```

