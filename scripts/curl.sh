id=$(uuidgen)

echo "Transaction id: $id"

curl http://localhost:8080/v1/transactions --silent \
  -H 'Content-Type: application/json' \
  -H "X-Idempotency-Key: 7adedb40-539a-45b7-976f-b22f3a2705c1" \
  -d '{
    "amount": 100,
    "payer": "01127017-ac87-436f-926b-67e49f9930dd",
    "payee": "02fc0614-29ce-4c74-9908-6ca6be44027b"
  }' | jq

# while true;
# do
# done

