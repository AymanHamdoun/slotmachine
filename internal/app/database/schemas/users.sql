CREATE TABLE user_registration_methods (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE users (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    mobile_number VARCHAR(255),

    registration_method_id INTEGER NOT NULL,

    status TINYINT(1) DEFAULT 1,
  
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
  
    INDEX idx_email (email),
    INDEX idx_mobile_number (mobile_number)
);