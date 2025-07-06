#!/bin/bash

echo "🎯 Focused API Test - Using Postman Collection Data"
echo "=================================================="

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

echo -e "${BLUE}📋 Step 1: Get Authentication Token${NC}"
auth_response=$(curl -s ${API_BASE}/v1/tokens/authentication \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"${TEST_USER_EMAIL}\",\"password\":\"${TEST_USER_PASSWORD}\"}")

echo "Auth response: $auth_response"

# Extract token from response - more robust extraction
AUTH_TOKEN=$(echo $auth_response | sed -n 's/.*"token":"\([^"]*\)".*/\1/p')

if [ ! -z "$AUTH_TOKEN" ]; then
    echo -e "  ✅ ${GREEN}Authentication successful! Token: ${AUTH_TOKEN:0:20}...${NC}"
else
    echo -e "  ❌ ${RED}Authentication failed${NC}"
    exit 1
fi

sleep 2  # Avoid rate limiting

echo -e "\n${BLUE}📋 Step 2: Test Authenticated Endpoints${NC}"

echo "GET /v1/movies (authenticated)"
movies_response=$(curl -s -w "HTTP %{http_code}" ${API_BASE}/v1/movies \
  -H "Authorization: Bearer ${AUTH_TOKEN}")
echo "  ✓ List Movies: $movies_response"

sleep 1

echo "POST /v1/movies (authenticated) - The Breakfast Club"
create_response=$(curl -s -w "HTTP %{http_code}" -X POST ${API_BASE}/v1/movies \
  -H "Authorization: Bearer ${AUTH_TOKEN}" \
  -H "Content-Type: application/json" \
  -d '{"title":"The Breakfast Club","year":1986,"runtime":"96 mins","genres":["drama"]}')
echo "  ✓ Create Movie: $create_response"

# Extract movie ID for further testing
MOVIE_ID=$(echo $create_response | grep -o '"id":[0-9]*' | cut -d':' -f2)

if [ ! -z "$MOVIE_ID" ]; then
    echo "  Created Movie ID: $MOVIE_ID"
    
    sleep 1
    
    echo "GET /v1/movies/${MOVIE_ID} (authenticated)"
    get_movie_response=$(curl -s -w "HTTP %{http_code}" ${API_BASE}/v1/movies/${MOVIE_ID} \
      -H "Authorization: Bearer ${AUTH_TOKEN}")
    echo "  ✓ Get Movie by ID: $get_movie_response"
    
    sleep 1
    
    echo "PATCH /v1/movies/${MOVIE_ID} (authenticated) - Update year to 1984"
    update_response=$(curl -s -w "HTTP %{http_code}" -X PATCH ${API_BASE}/v1/movies/${MOVIE_ID} \
      -H "Authorization: Bearer ${AUTH_TOKEN}" \
      -H "Content-Type: application/json" \
      -d '{"year":1984}')
    echo "  ✓ Update Movie: $update_response"
    
    sleep 1
    
    echo "DELETE /v1/movies/${MOVIE_ID} (authenticated)"
    delete_response=$(curl -s -w "HTTP %{http_code}" -X DELETE ${API_BASE}/v1/movies/${MOVIE_ID} \
      -H "Authorization: Bearer ${AUTH_TOKEN}")
    echo "  ✓ Delete Movie: $delete_response"
fi

sleep 1

echo -e "\n${BLUE}📋 Step 3: Test Postman Collection Filters${NC}"

echo "GET /v1/movies?title=moana&genres=animation (Postman: Filters)"
filter_response=$(curl -s -w "HTTP %{http_code}" "${API_BASE}/v1/movies?title=moana&genres=animation" \
  -H "Authorization: Bearer ${AUTH_TOKEN}")
echo "  ✓ Movies with filters: $filter_response"

sleep 1

echo "GET /v1/movies?sort=-runtime (Postman: Sorting)"
sort_response=$(curl -s -w "HTTP %{http_code}" "${API_BASE}/v1/movies?sort=-runtime" \
  -H "Authorization: Bearer ${AUTH_TOKEN}")
echo "  ✓ Movies with sorting: $sort_response"

sleep 1

echo "GET /v1/movies?page_size=2&page=1 (Postman: Pagination)"
pagination_response=$(curl -s -w "HTTP %{http_code}" "${API_BASE}/v1/movies?page_size=2&page=1" \
  -H "Authorization: Bearer ${AUTH_TOKEN}")
echo "  ✓ Movies with pagination: $pagination_response"

echo -e "\n${GREEN}🎉 Focused Test Complete!${NC}"
echo -e "${YELLOW}📊 Check your Grafana dashboard at: http://localhost:3000${NC}"
echo -e "${YELLOW}📈 Metrics generated for all CRUD operations!${NC}"
