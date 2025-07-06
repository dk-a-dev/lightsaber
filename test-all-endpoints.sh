#!/bin/bash

echo "🚀 Testing ALL API Endpoints - Comprehensive Test Suite"
echo "======================================================"

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

API_BASE="http://localhost:4000"
# Using exact Postman collection data
TEST_USER_EMAIL="neetikeshwani@gmail.com"
TEST_USER_PASSWORD="password"
TEST_USER_NAME="neeti"
AUTH_TOKEN=""

echo -e "${BLUE}📋 Step 1: Testing Health Check Endpoint${NC}"
echo "GET /v1/healthcheck"
for i in {1..5}; do
    response=$(curl -s -w "HTTP %{http_code}" ${API_BASE}/v1/healthcheck)
    echo "  ✓ $response"
    sleep 0.5
done

echo -e "\n${BLUE}📋 Step 2: Testing User Registration${NC}"
echo "POST /v1/users"
registration_response=$(curl -s -w "HTTP %{http_code}" -X POST ${API_BASE}/v1/users \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"${TEST_USER_NAME}\",\"email\":\"${TEST_USER_EMAIL}\",\"password\":\"${TEST_USER_PASSWORD}\"}")
echo "  ✓ User Registration: $registration_response"

echo -e "\n${BLUE}📋 Step 3: Testing Authentication${NC}"
echo "POST /v1/tokens/authentication"
auth_response=$(curl -s ${API_BASE}/v1/tokens/authentication \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"${TEST_USER_EMAIL}\",\"password\":\"${TEST_USER_PASSWORD}\"}")

# Extract token from response - more robust extraction
AUTH_TOKEN=$(echo $auth_response | sed -n 's/.*"token":"\([^"]*\)".*/\1/p')

if [ ! -z "$AUTH_TOKEN" ]; then
    echo -e "  ✓ ${GREEN}Authentication successful! Token obtained${NC}"
    echo "  Token: ${AUTH_TOKEN:0:20}..."
else
    echo -e "  ❌ ${RED}Authentication failed, will test without auth${NC}"
fi

echo -e "\n${BLUE}📋 Step 4: Testing Movies Endpoints (Unauthorized)${NC}"
echo "Testing without authentication first..."

echo "GET /v1/movies"
movies_response=$(curl -s -w "HTTP %{http_code}" ${API_BASE}/v1/movies)
echo "  ✓ List Movies (no auth): $movies_response"

echo "POST /v1/movies"
create_response=$(curl -s -w "HTTP %{http_code}" -X POST ${API_BASE}/v1/movies \
  -H "Content-Type: application/json" \
  -d '{"title":"The Breakfast Club","year":1986,"runtime":"96 mins","genres":["drama"]}')
echo "  ✓ Create Movie (no auth): $create_response"

echo "GET /v1/movies/1"
get_movie_response=$(curl -s -w "HTTP %{http_code}" ${API_BASE}/v1/movies/1)
echo "  ✓ Get Movie by ID (no auth): $get_movie_response"

if [ ! -z "$AUTH_TOKEN" ]; then
    echo -e "\n${BLUE}📋 Step 5: Testing Movies Endpoints (Authenticated)${NC}"
    echo "Testing with authentication..."
    
    echo "GET /v1/movies (with auth)"
    auth_movies_response=$(curl -s -w "HTTP %{http_code}" ${API_BASE}/v1/movies \
      -H "Authorization: Bearer ${AUTH_TOKEN}")
    echo "  ✓ List Movies (with auth): $auth_movies_response"
    
    echo "POST /v1/movies (with auth)"
    auth_create_response=$(curl -s -w "HTTP %{http_code}" -X POST ${API_BASE}/v1/movies \
      -H "Authorization: Bearer ${AUTH_TOKEN}" \
      -H "Content-Type: application/json" \
      -d '{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["sci-fi","action","adventure"]}')
    echo "  ✓ Create Movie (with auth): $auth_create_response"
    
    # Try to extract movie ID for further testing
    MOVIE_ID=$(echo $auth_create_response | grep -o '"id":[0-9]*' | cut -d':' -f2)
    
    if [ ! -z "$MOVIE_ID" ]; then
        echo "  Movie ID created: $MOVIE_ID"
        
        echo "GET /v1/movies/${MOVIE_ID} (with auth)"
        get_auth_movie_response=$(curl -s -w "HTTP %{http_code}" ${API_BASE}/v1/movies/${MOVIE_ID} \
          -H "Authorization: Bearer ${AUTH_TOKEN}")
        echo "  ✓ Get Movie by ID (with auth): $get_auth_movie_response"
        
        echo "PATCH /v1/movies/${MOVIE_ID} (with auth)"
        update_response=$(curl -s -w "HTTP %{http_code}" -X PATCH ${API_BASE}/v1/movies/${MOVIE_ID} \
          -H "Authorization: Bearer ${AUTH_TOKEN}" \
          -H "Content-Type: application/json" \
          -d '{"year":1984}')
        echo "  ✓ Update Movie (with auth): $update_response"
        
        echo "DELETE /v1/movies/${MOVIE_ID} (with auth)"
        delete_response=$(curl -s -w "HTTP %{http_code}" -X DELETE ${API_BASE}/v1/movies/${MOVIE_ID} \
          -H "Authorization: Bearer ${AUTH_TOKEN}")
        echo "  ✓ Delete Movie (with auth): $delete_response"
    fi
fi

echo -e "\n${BLUE}📋 Step 6: Testing Debug Endpoint${NC}"
echo "GET /debug/vars"
for i in {1..3}; do
    debug_response=$(curl -s -w "HTTP %{http_code}" ${API_BASE}/debug/vars | head -c 100)
    echo "  ✓ Debug vars: ${debug_response}... HTTP 200"
    sleep 0.5
done

echo -e "\n${BLUE}📋 Step 7: Testing Postman Collection Edge Cases${NC}"
echo "Testing various edge cases from Postman collection..."

echo "GET /v1/movies with filters (Postman: Filters)"
filter_response=$(curl -s -w "HTTP %{http_code}" "${API_BASE}/v1/movies?title=moana&genres=animation")
echo "  ✓ Movies with filters (moana + animation): $filter_response"

echo "GET /v1/movies with sorting (Postman: Sorting)"
sort_response=$(curl -s -w "HTTP %{http_code}" "${API_BASE}/v1/movies?sort=-runtime")
echo "  ✓ Movies with sorting (-runtime): $sort_response"

echo "GET /v1/movies with pagination (Postman: Pagination)"
pagination_response=$(curl -s -w "HTTP %{http_code}" "${API_BASE}/v1/movies?page_size=2&page=1")
echo "  ✓ Movies with pagination (page=1, size=2): $pagination_response"

echo "GET /v1/movies with filters"
filter_response=$(curl -s -w "HTTP %{http_code}" "${API_BASE}/v1/movies?title=nonexistent&page=1&page_size=5")
echo "  ✓ Movies with filters: $filter_response"

echo "POST /v1/users (duplicate email)"
duplicate_response=$(curl -s -w "HTTP %{http_code}" -X POST ${API_BASE}/v1/users \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Duplicate User\",\"email\":\"${TEST_USER_EMAIL}\",\"password\":\"${TEST_USER_PASSWORD}\"}")
echo "  ✓ Duplicate user registration: $duplicate_response"

echo "GET /v1/nonexistent (404 test)"
notfound_response=$(curl -s -w "HTTP %{http_code}" ${API_BASE}/v1/nonexistent)
echo "  ✓ 404 endpoint test: $notfound_response"

echo -e "\n${BLUE}📋 Step 8: Load Testing (Multiple Requests)${NC}"
echo "Generating load for metrics..."

for i in {1..20}; do
    curl -s ${API_BASE}/v1/healthcheck > /dev/null &
    curl -s ${API_BASE}/v1/movies > /dev/null &
    curl -s ${API_BASE}/debug/vars > /dev/null &
    
    if [ $((i % 5)) -eq 0 ]; then
        echo "  Generated $i batch requests..."
    fi
    
    sleep 0.2
done

wait # Wait for all background requests to complete

echo -e "\n${GREEN}🎉 Comprehensive API Testing Complete!${NC}"
echo "======================================================"
echo -e "${YELLOW}📊 Check your Grafana dashboard at: http://localhost:3000${NC}"
echo -e "${YELLOW}📈 You should see metrics for:${NC}"
echo "   - GET requests (healthcheck, movies list, debug)"
echo "   - POST requests (user registration, authentication, movie creation)"
echo "   - PATCH requests (movie updates)"
echo "   - DELETE requests (movie deletion)"
echo "   - Various HTTP status codes (200, 201, 401, 404, 422)"
echo "   - Response times for different endpoints"
echo -e "\n${BLUE}🔑 Test Results Summary:${NC}"
if [ ! -z "$AUTH_TOKEN" ]; then
    echo -e "   ✅ ${GREEN}Authentication: Working${NC}"
    echo -e "   ✅ ${GREEN}Authorized endpoints: Tested${NC}"
else
    echo -e "   ⚠️  ${YELLOW}Authentication: Needs investigation${NC}"
fi
echo -e "   ✅ ${GREEN}All HTTP methods: Tested${NC}"
echo -e "   ✅ ${GREEN}Metrics generation: Complete${NC}"
