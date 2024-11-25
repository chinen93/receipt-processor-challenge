curl -d @simple-receipt.json -X POST http://localhost:8080/receipts/process 
echo
curl -d @full-receipt.json -X POST http://localhost:8080/receipts/process
echo
curl -d @morning-receipt.json -X POST http://localhost:8080/receipts/process
echo
curl http://localhost:8080/receipts/0/points
echo
curl http://localhost:8080/receipts/1/points
echo
curl http://localhost:8080/receipts/2/points
echo