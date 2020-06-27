project = jobscheduler
.Phony : install
install:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${project} cmd/main.go


.Phony : run
run: delete
	 -docker network create xyz
	 docker pull consul
	 docker run -d --name=dev-consul --network=xyz  -p8500:8500 --rm consul
	 docker build -t ${project} .
	 docker run  --name ${project} --rm --network=xyz --env-file=envconfig ${project}

.Phony : delete
delete:
	-docker stop dev-consul
	-docker stop jobscheduler
