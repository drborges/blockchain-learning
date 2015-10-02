update:
	go get -u github.com/jmcvetta/neoism

neo4j:
	docker run -i -t --rm --name neo4j -v $(HOME)/neo4j-data\:/data -p 8474\:7474 neo4j/neo4j

start: update neo4j
	go run graph/app.go

test:
	go test ./...