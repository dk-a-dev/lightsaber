#!/bin/bash

echo "🚀 Running comprehensive API tests to generate metrics..."

# Test health check endpoint multiple times
echo "Testing health check endpoint..."
for i in {1..10}; do
    curl -s http://localhost:4000/v1/healthcheck > /dev/null
    sleep 0.5
done

# Test movies endpoints (will get 401 but still generates metrics)
echo "Testing movies endpoints..."
for i in {1..5}; do
    curl -s http://localhost:4000/v1/movies > /dev/null
    curl -s -X POST http://localhost:4000/v1/movies -H "Content-Type: application/json" -d '{"title":"Test"}' > /dev/null
    curl -s http://localhost:4000/v1/movies/1 > /dev/null
    curl -s -X PATCH http://localhost:4000/v1/movies/1 -H "Content-Type: application/json" -d '{"title":"Updated"}' > /dev/null
    curl -s -X DELETE http://localhost:4000/v1/movies/1 > /dev/null
    sleep 1
done

# Test user registration (will fail but generates metrics)
echo "Testing user endpoints..."
for i in {1..3}; do
    curl -s -X POST http://localhost:4000/v1/users -H "Content-Type: application/json" -d '{"name":"test","email":"test@example.com","password":"password"}' > /dev/null
    curl -s -X POST http://localhost:4000/v1/tokens/authentication -H "Content-Type: application/json" -d '{"email":"test@example.com","password":"password"}' > /dev/null
    sleep 1
done

# Test debug endpoint
echo "Testing debug endpoint..."
for i in {1..5}; do
    curl -s http://localhost:4000/debug/vars > /dev/null
    sleep 0.5
done

echo "✅ Test completed! Check your Grafana dashboard at http://localhost:3000"
echo "📊 You should now see metrics for all HTTP methods: GET, POST, PATCH, DELETE"
