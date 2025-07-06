# Lightsaber API

A modern Go REST API with Docker deployment.

## 🚀 Quick Start

### Development
```bash
# Start development environment
make docker/up

# Run tests
make test

# View logs
make docker/logs/api
```

### Production Deployment
```bash
# Deploy to production
make deploy/prod

# Or use the deployment script directly
./deploy.sh prod
```

## 📋 Available Commands

### Development
- `make run` - Run API locally
- `make test` - Run all tests
- `make docker/up` - Start dev services
- `make docker/down` - Stop dev services

### Production
- `make deploy/prod` - Deploy to production
- `make docker/prod/up` - Start production services
- `make docker/prod/down` - Stop production services

### Logs & Monitoring
- `make docker/logs` - View all logs
- `make docker/logs/api` - View API logs only

## 🔧 Configuration

### Environment Setup

1. **Copy the example environment file:**
   ```bash
   cp .env.example .env
   ```

2. **Edit `.env` with your settings:**
   ```env
   # Database
   DB_USER=greenlight
   DB_PASSWORD=your_secure_password

   # SMTP (for email functionality)
   HOST=sandbox.smtp.mailtrap.io
   PORT=2525
   USERNAME=your_smtp_username
   PASSWORD=your_smtp_password
   SENDER=your@email.com

   # CORS (for frontend integration)
   TRUSTED_ORIGINS=http://localhost:9000,http://localhost:9001

   # Database URL (for production)
   DATABASE_URL=postgres://username:password@host:port/database
   ```

⚠️ **Never commit your `.env` file to version control!** It contains sensitive information.
```

## 🌐 API Endpoints

- **Health Check**: `GET /v1/healthcheck`
- **Movies**: `GET|POST /v1/movies`
- **Users**: Authentication endpoints
- **Tokens**: Token management

## 🏗️ Architecture

- **Go 1.24** - Modern Go backend
- **PostgreSQL 16** - Database with migrations
- **Docker** - Containerized deployment
- **GitHub Actions** - CI/CD pipeline
- **Comprehensive Testing** - Unit and integration tests

## 📊 Production Features

- ✅ Health checks and monitoring
- ✅ Graceful shutdowns
- ✅ Database migrations
- ✅ Structured logging
- ✅ Security best practices
- ✅ Automatic restarts

## 🔒 Security

- Non-root container user
- Environment variable secrets
- CORS configuration
- Rate limiting
- Input validation

## 📈 Monitoring

Access your API:
- **Development**: http://localhost:4000
- **Production**: http://localhost (port 80)
- **Health Check**: `/v1/healthcheck`

---

Built with ❤️ using Go and Docker
