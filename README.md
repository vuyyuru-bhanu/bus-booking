# Bus Booking Application

This is a bus booking application built with **Golang** (backend), **Angular** (frontend), **MySQL** (database), and **Nginx** (reverse proxy). The application supports booking bus tickets, searching for buses based on origin, destination, AC/non-AC, bus type (seater/sleeper), and company, as well as admin functionalities for managing routes, buses, bus images, and notifications. It runs in containers using Docker and can be deployed on Kubernetes.

Inspired by [Busbud](https://www.busbud.com/en-ca/country/ca), it includes five sample routes, five bus companies, and options for AC/non-AC and seater/sleeper buses.

## Features
- **User Features**:
  - Register and log in with secure JWT authentication.
  - Search buses by origin, destination, AC/non-AC, type, and company.
  - Book a bus with seat selection.
  - View service notifications.
- **Admin Features**:
  - Manage routes (create new routes).
  - Manage buses (add buses with details like company, type, capacity, and amenities).
  - Upload bus images (JPEG/PNG, max 5MB).
  - Post and view service notifications.
  - View all bookings.
- **Technical Features**:
  - RESTful API built with Golang and Gin.
  - Angular frontend with Material UI for a responsive interface.
  - MySQL database with GORM ORM.
  - Containerized with Docker and orchestrated with Kubernetes.
  - Nginx reverse proxy for serving the frontend and API.

## Prerequisites
- **Docker**: Install [Docker](https://docs.docker.com/get-docker/) for containerization.
- **Docker Compose**: Included with Docker Desktop, or install separately for Linux.
- **Kubernetes**: Install [kubectl](https://kubernetes.io/docs/tasks/tools/) and a Kubernetes cluster (e.g., [Minikube](https://minikube.sigs.k8s.io/docs/start/)) for Kubernetes deployment.
- **Git**: For version control.
- **Node.js**: Version 18 for building the Angular frontend.
- **Golang**: Version 1.21 for building the backend.
- **MySQL Client**: Optional, for manual database inspection.

## Project Structure
```
bus-booking-app/
├── backend/                # Golang backend (API)
├── frontend/               # Angular frontend
├── kubernetes/             # Kubernetes manifests
├── nginx/                  # Nginx configuration
├── docker-compose.yml      # Docker Compose configuration
├── init.sql               # MySQL initialization script
└── README.md              # This file
```

## Setup Instructions

### 1. Clone the Repository
```bash
git clone <your-repository-url>
cd bus-booking-app
```

### 2. Environment Variables
Create a `.env` file in the root directory with the following variables:
```bash
DB_HOST=mysql
DB_USER=bus_user
DB_PASSWORD=secure_password
DB_NAME=bus_booking
JWT_SECRET=your_jwt_secret_key
```
**Note**: Replace `your_jwt_secret_key` with a secure random string in production. The values above match the Docker Compose and Kubernetes configurations.

### 3. Database Setup
The `init.sql` script initializes the MySQL database with:
- Database schema for `users`, `routes`, `buses`, `bookings`, `bus_images`, and `notifications`.
- A default admin user:
  - **Email**: `admin@busbooking.com`
  - **Password**: `Admin123!`
  - **Role**: `admin`
- Sample routes (e.g., Toronto to Montreal, Vancouver to Calgary).
- Sample buses for each route with companies (Greyhound, Megabus, etc.), AC/non-AC, and seater/sleeper options.

The script is automatically applied when the MySQL container starts.

### 4. Running with Docker Compose
1. Build and start the services:
   ```bash
   docker-compose up --build
   ```
2. Access the application:
   - **Frontend**: `http://localhost`
   - **Backend API**: `http://localhost/api`
   - **MySQL**: `localhost:3306` (use a MySQL client to connect with `bus_user:secure_password`).
3. Log in with the admin credentials or register a new user.
4. To stop the services:
   ```bash
   docker-compose down
   ```

### 5. Running with Kubernetes
1. Ensure a Kubernetes cluster is running (e.g., Minikube):
   ```bash
   minikube start
   ```
2. Create a ConfigMap for the MySQL initialization script:
   ```bash
   kubectl create configmap mysql-initdb-config --from-file=init.sql
   ```
3. Apply the Kubernetes manifests:
   ```bash
   kubectl apply -f kubernetes/
   ```
4. If using Minikube, get the service URL:
   ```bash
   minikube service nginx-service --url
   ```
5. Access the application at the provided URL.
6. To clean up:
   ```bash
   kubectl delete -f kubernetes/
   minikube stop
   ```

### 6. Building and Running Locally (Without Docker)
#### Backend
1. Navigate to the backend directory:
   ```bash
   cd backend
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up a local MySQL database and update the `.env` file with the correct `DB_HOST`.
4. Run the backend:
   ```bash
   go run cmd/main.go
   ```

#### Frontend
1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Build and serve:
   ```bash
   npm start
   ```
4. Access the frontend at `http://localhost:4200`.

#### Nginx
Configure Nginx locally to proxy requests to `localhost:8080` (backend) and `localhost:4200` (frontend), using the provided `nginx/nginx.conf` as a guide.

## Default Admin Credentials
- **Admin URL**: `http://localhost/admin`
- **Email**: `admin@busbooking.com`
- **Password**: `Admin123!`
- **Note**: Change the password in production and store sensitive data in Kubernetes Secrets or a secrets manager.

## Security Considerations
- **JWT Secret**: Use a strong, unique `JWT_SECRET` in production.
- **Password Hashing**: Passwords are hashed using bcrypt.
- **Image Uploads**: Limited to JPEG/PNG, max 5MB, stored in `/uploads`.
- **Database**: Use secure passwords and consider Kubernetes Secrets for sensitive environment variables.
- **HTTPS**: Configure Nginx for HTTPS in production using SSL certificates.

## Sample Data
The `init.sql` seeds:
- **Routes**: 5 routes (e.g., Toronto-Montreal, Vancouver-Calgary).
- **Buses**: Multiple buses per route with companies (Greyhound, Megabus, Orleans Express, Rider Express, Ebus), AC/non-AC, seater/sleeper, and amenities (e.g., Wi-Fi, charging).
- **Admin User**: For admin access to manage routes, buses, and notifications.

## Troubleshooting
- **Database Connection Issues**: Verify `DB_HOST`, `DB_USER`, `DB_PASSWORD`, and `DB_NAME` in the `.env` file or Kubernetes manifests.
- **Frontend Not Loading**: Ensure the backend is running and Nginx is correctly proxying to `frontend-service:4200`.
- **Image Upload Fails**: Check the `/uploads` directory permissions and ensure it’s mounted correctly in Docker/Kubernetes.
- **Kubernetes Issues**: Use `kubectl logs` and `kubectl describe` to debug pod issues.

## Contributing
1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/your-feature`).
3. Commit changes (`git commit -m "Add your feature"`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

## License
MIT License. See [LICENSE](LICENSE) for details.

---

## Admin Credentials

- Username: admin
- Email: admin@busbooking.com
- Password: Hello123@

Note: Only the admin user can login with the username "admin". Regular users must login with their email and password.
