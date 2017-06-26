#!/bin/bash
if [ "$BRANCH" == "master" ]; then
    docker build -t quay.io/genesor/cochonou .;
    docker tag quay.io/genesor/cochonou:latest quay.io/genesor/cochonou:$COMMIT;
    docker push quay.io/genesor/cochonou;
    export SSHPASS=$DEPLOY_PWD;
    sshpass -e ssh $DEPLOY_USER@$DEPLOY_HOST $DEPLOY_SCRIPT;
else
    docker build -t quay.io/genesor/cochonou:$BRANCH .;
    docker quay.io/genesor/cochonou:$BRANCH;
fi