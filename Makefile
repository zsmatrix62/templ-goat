ssh_target =
image_name = 

dep: build trans trans-img restart
dep-docs: build-docs trans-img-docs restart-docs

build: 
	go generate 
	docker buildx build --platform=linux/amd64 -t $(image_name):latest .

trans-img: 
	docker save $(image_name) | bzip2 | ssh $(ssh_target) docker load

trans: 

restart: trans

# dev
air: 
	@air
