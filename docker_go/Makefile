image-build:
	docker image build -f Dockerfile_build -t gow_api .

install-dev:
	docker container run -it -p 9000:9000 --rm -v $(PWD):/api/src -t gow_api

run-dev:
	docker container run -it -p 9000:9000 --rm -v $(PWD):/api/src -t gow_api run main.go

build-dev:
	#docker container run -it -p 9000:9000 --rm -v $(PWD):/api/src -t gow_api build -a -installsuffix cgo -o api 
	docker container run -it -p 9000:9000 --rm -v $(PWD):/api/src gow_api build -o api .

test-dev:
	docker container run -it -p 9000:9000 --rm -v $(PWD):/api/src -t gow_api test

image:
	docker image build -f Dockerfile -t api .
start:
	docker container run -d -p 9000:9000 --name api01 -t api
stop:
	docker container rm -f api01
	
clean:
	go clean
	echo "Y" | docker container prune
	echo "Y" | docker image prune
	echo "Y" | docker volume prune
