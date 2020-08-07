run: bin/Rocket_Elevator_Rest_API_GOLang
	@PATH="$(PWD)/bin:$(PATH)" heroku local

bin/Rocket_Elevator_Rest_API_GOLang: server.go
	go build -o bin/Rocket_Elevator_Rest_API_GOLang server.go

clean:
	rm -rf bin