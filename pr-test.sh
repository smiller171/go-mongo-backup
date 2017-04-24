unzip OpenWhere_go-mongo-backup.zip

curl -H "Authorization: token $GITHUB_KEY" -XPOST "https://api.github.com/repos/OpenWhere/go-mongo-backup/statuses/5d1389a3d6f2ba541afb6382427d219f35cf3f8d" -d '{"state":"pending", "description":"Started the build!","context":"ci/codePipeline"}'
