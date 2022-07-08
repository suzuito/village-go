init:
	if [ ! -e _devenv/local/dev.env1 ]; then cp _devenv/local/dev.env.sample _devenv/local/dev.env; fi
	docker-compose up -d
test:
	echo "FIXME!"