# setup

make environment file

`cp env{.tmp,}.json`

# deploy

apex -r ap-northeast-1 --env-file ./env.json deploy
