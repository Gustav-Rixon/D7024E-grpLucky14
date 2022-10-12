
CONT_NAME="d7024e-grplucky14-kademlia-" # Common part of container names

echo "Initilizing all running Kademlia containers..."

# Get all container IDs
cont_ids=$(docker ps -aq -f name=$CONT_NAME -f status=running) 

arrIp=()

# Init each node with their assigned IPs
for id in $cont_ids; do
  cont_ip="$(docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $id)"
  #echo $cont_ip
  arrIp+=($cont_ip)
  #echo $arrIp
done

# Insert a known node in the network to each routing table
known_node_cid="$(echo "$cont_ids" | head -n 1)"
i=0
for id in $cont_ids; do
    if [ "$id" != "$known_node_cid" ]; then
        docker exec -ti $id cli join ${arrIp[i]}
        #echo $id "joining on" ${arrIp[i]}
        ((i++))
    fi
done

echo "Done"