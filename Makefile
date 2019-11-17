VERSION ?= 0.1.0

image:
	docker build -t michellenoorali/greeting-service:${VERSION} .
