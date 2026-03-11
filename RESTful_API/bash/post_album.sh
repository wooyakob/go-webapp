curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "3","title": "Recollections","artist": "Simply Red","price": 39.99, "currency": "USD", "release_year": 2025}'