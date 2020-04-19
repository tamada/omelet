PWD := ${CURDIR}
CONTAINER_NAME := "wwwomelet"
CONTAINER_REMOVE_FLAG := "--rm"
BASE_URL := "https://tamada.github.io/omelet"


build:
	docker run ${CONTAINER_REMOVE_FLAG} --name ${CONTAINER_NAME}_build -v "${PWD}":/src -v ${PWD}/public:/output -e HUGO_THEME=cayman-hugo-theme -e HUGO_BASEURL=${BASE_URL} jojomi/hugo:latest

start:
	docker run ${CONTAINER_REMOVE_FLAG} -d --name ${CONTAINER_NAME} -p 1313:1313 -v "${PWD}":/src -v "$(PWD)"/public:/output -e HUGO_THEME=cayman-hugo-theme -e HUGO_WATCH="true" -e HUGO_BASEURL=${BASE_URL} jojomi/hugo:latest

stop:
	docker stop ${CONTAINER_NAME}
