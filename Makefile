NAME=gopher-translator
VERSION=0.1.0

build:
	docker build -t $(NAME):$(VERSION) .

restart: stop start

start:
	docker run \
		--name $(NAME) \
		-v $(CURDIR)\:/go/src/github.com/fujiwaram/gopher-translator \
		-itd $(NAME):$(VERSION)

contener=`docker ps -a -q`
image=`docker images | awk '/^<none>/ { print $$3 }'`

clean:
	@if [ "$(image)" != "" ] ; then \
		docker rmi $(image); \
	fi
	@if [ "$(contener)" != "" ] ; then \
		docker rm $(contener); \
	fi

stop:
	docker rm -f $(NAME)

attach:
	docker exec -it $(NAME) /bin/sh

test:
	docker exec -it $(NAME) go test

logs:
	docker logs $(NAME)