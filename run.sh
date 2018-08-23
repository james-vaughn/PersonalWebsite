git pull
go build --tags prod
pkill PersonalWebsite
nohup ./PersonalWebsite & >> dev/null