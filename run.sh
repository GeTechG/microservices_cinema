fuser -k 8080/tcp
fuser -k 5000/tcp
fuser -k 5001/tcp
cd users_service
sh run.sh &
cd ../movies_service
sh run.sh &
cd ../recommendation_service
sh run.sh &
