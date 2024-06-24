# build.sh
#!/bin/bash


cp ../go.sum ../go.mod ./

cd .. && cd deployments
docker-compose -f docker-compose.development.yaml up --build


rm ../build/go.sum ../build/go.mod