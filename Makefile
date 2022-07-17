init:
	if [ ! -e _devenv/local/dev.env ]; then cp _devenv/local/dev.env.sample _devenv/local/dev.env; fi
	docker-compose up -d
	docker-compose exec -T local /bin/sh -c 'until (curl http://firebase:8081) do sleep 2; done'
	echo "Firestore -> OK"
	docker-compose exec -T local /bin/sh -c 'go run ./cmd/init-static-feed-settings/*.go'
test:
	docker-compose exec -T local /bin/sh -c 'go test ./...'