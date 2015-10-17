
#!/bin/bash
set -e

SRC_DIR=/go/src/github.com/aleasoluciones/godockerminimal

# We need a temp dir to mount to the build process 
# because when using this script from docker, the 
# real dir mounted is from the host
WORKDIR=/tmp/$(hostname)
mkdir -p ${WORKDIR}
cp -a . ${WORKDIR}
docker run -v ${WORKDIR}:${SRC_DIR} -e CGO_ENABLED=0 -e GOOS=linux -ti golang bash -c "cd ${SRC_DIR};go build -a -installsuffix cgo ."
cp -v ${WORKDIR}/godockerminimal .
docker build  --no-cache -t aleasoluciones/godockerminimal .

rm -rf ${WORKDIR}
