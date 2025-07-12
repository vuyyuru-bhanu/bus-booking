-- Create the database
CREATE DATABASE IF NOT EXISTS bus_booking;
USE bus_booking;

-- Users table
CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    role ENUM('user', 'admin') NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_email (email),
    INDEX idx_role (role)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Routes table
CREATE TABLE routes (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    origin VARCHAR(255) NOT NULL,
    destination VARCHAR(255) NOT NULL,
    departure_time VARCHAR(255) NOT NULL,
    duration INT NOT NULL CHECK (duration > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_origin_destination (origin, destination)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Buses table
CREATE TABLE buses (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    route_id BIGINT UNSIGNED NOT NULL,
    company VARCHAR(255) NOT NULL,
    ac BOOLEAN NOT NULL,
    type ENUM('seater', 'sleeper') NOT NULL,
    capacity INT NOT NULL CHECK (capacity > 0),
    available_seats INT NOT NULL CHECK (available_seats >= 0), -- ✅ fixed
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    amenities JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (route_id) REFERENCES routes(id) ON DELETE CASCADE,
    INDEX idx_route_id (route_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bookings table
CREATE TABLE bookings (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    bus_id BIGINT UNSIGNED NOT NULL,
    seat_number INT NOT NULL CHECK (seat_number > 0),
    status ENUM('confirmed', 'cancelled') DEFAULT 'confirmed',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (bus_id) REFERENCES buses(id) ON DELETE CASCADE,
    UNIQUE INDEX idx_bus_seat (bus_id, seat_number),
    INDEX idx_user_id (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Bus images table
CREATE TABLE bus_images (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    bus_id BIGINT UNSIGNED NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (bus_id) REFERENCES buses(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Notifications table
CREATE TABLE notifications (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Payments table
CREATE TABLE payments (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    booking_id BIGINT UNSIGNED NOT NULL,
    amount DECIMAL(10,2) NOT NULL CHECK (amount >= 0),
    payment_status ENUM('completed', 'pending', 'failed') NOT NULL DEFAULT 'pending',
    payment_method ENUM('credit_card', 'debit_card', 'upi', 'cash') NOT NULL,
    transaction_id VARCHAR(100) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (booking_id) REFERENCES bookings(id) ON DELETE CASCADE,
    INDEX idx_booking_id (booking_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Admin dashboard view
CREATE VIEW admin_booking_summary AS
SELECT 
    b.id AS booking_id,
    u.email AS user_email,
    r.origin,
    r.destination,
    r.departure_time,
    b.seat_number,
    b.status AS booking_status,
    p.amount,
    p.payment_status
FROM bookings b
JOIN users u ON b.user_id = u.id
JOIN buses bu ON b.bus_id = bu.id
JOIN routes r ON bu.route_id = r.id
LEFT JOIN payments p ON b.id = p.booking_id;

-- Seed data
INSERT INTO users (email, password, name, role)
VALUES 
    ('admin@busbooking.com', '$2a$10$0YzvY7z0Qz7fG9z7z7z7z7z7z7z7z7z7z7z7z7z7z7z7z7z7z7z7z', 'Admin User', 'admin'),
    ('passenger@busbooking.com', '$2a$10$examplehashedpassword0987654321', 'John Doe', 'user');

INSERT INTO routes (origin, destination, departure_time, duration)
VALUES 
    ('New York', 'Boston', '2025-07-15 08:00:00', 240),
    ('Chicago', 'Detroit', '2025-07-15 09:00:00', 300);

INSERT INTO buses (route_id, company, ac, type, capacity, available_seats, price, amenities)
VALUES 
    (1, 'City Travel Co', TRUE, 'sleeper', 40, 40, 50.00, '{"wifi": true, "charging": true}'),
    (2, 'Express Rides', FALSE, 'seater', 50, 50, 35.00, '{"wifi": false, "charging": true}');

INSERT INTO bookings (user_id, bus_id, seat_number, status)
VALUES 
    (2, 1, 1, 'confirmed'),
    (2, 2, 5, 'confirmed');

INSERT INTO payments (booking_id, amount, payment_status, payment_method, transaction_id)
VALUES 
    (1, 50.00, 'completed', 'credit_card', 'TXN123456789'),
    (2, 35.00, 'pending', 'upi', 'TXN987654321');

INSERT INTO bus_images (bus_id, image_url)
VALUES 
    (1, 'https://example.com/images/bus1.jpg'),
    (1, 'https://example.com/images/bus1_interior.jpg'),
    (2, 'https://example.com/images/bus2.jpg');

INSERT INTO notifications (message)
VALUES 
    ('Booking confirmed for New York to Boston on 2025-07-15'),
    ('Payment pending for booking ID 2');
