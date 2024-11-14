docker-build-and-run:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/dimastephen/chatserver:v0.1 .
	docker login -u token -p CRgAAAAA17-2TAyanR1-isuZuMXThkOlueMSXtiE cr.selcloud.ru/dimastephen
	docker push cr.selcloud.ru/dimastephen/chatserver:v0.1