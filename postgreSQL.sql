
CREATE DATABASE  "GoTicket";



-- Membuat tabel users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL, 
    is_admin INTEGER DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Membuat tabel movies
CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    genre VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Membuat tabel showtimes
CREATE TABLE IF NOT EXISTS showtimes (
    id SERIAL PRIMARY KEY,
    movie_id INT NOT NULL,
    date_time TIMESTAMP NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (movie_id) REFERENCES movies(id)
);

-- Membuat tabel seats
CREATE TABLE IF NOT EXISTS seats (
    id SERIAL PRIMARY KEY,
    showtime_id INT NOT NULL,
    seat_number VARCHAR(10) NOT NULL,
    status INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (showtime_id) REFERENCES showtimes(id)
);

-- Membuat tabel checkouts
CREATE TABLE IF NOT EXISTS checkouts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    showtime_id INT NOT NULL,
    seat_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (showtime_id) REFERENCES showtimes(id),
    FOREIGN KEY (seat_id) REFERENCES seats(id)
);

-- Menambahkan data ke dalam tabel movies
INSERT INTO movies (title, genre) VALUES
('The Shawshank Redemption', 'Drama'),
('The Godfather', 'Crime, Drama'),
('The Dark Knight', 'Action, Crime, Drama');

-- Menambahkan data ke dalam tabel showtimes, perhatikan perbaikan pada sintaks INSERT untuk price
INSERT INTO showtimes (movie_id, date_time, price) VALUES
(1, '2024-04-25 13:00:00', 50000),
(1, '2024-04-25 16:00:00', 60000),
(2, '2024-04-26 13:00:00', 80000),
(3, '2024-04-26 20:00:00', 80000);


INSERT INTO seats (showtime_id, seat_number, status) VALUES
(1, 'A1', 1),
(1, 'A2', 1),
(1, 'A3', 1),
(1, 'A4', 1),
(1, 'B1', 1),
(1, 'B2',1),
(1, 'B3', 1),
(1, 'B4', 1),
(2, 'A1', 1),
(2, 'A2', 1),
(2, 'A3', 1),
(2, 'A4', 1),
(2, 'B1', 1),
(2, 'B2',1),
(2, 'B3', 1),
(2, 'B4', 1),
(3, 'A1', 1),
(3, 'A2', 1),
(3, 'A3', 1),
(3, 'A4', 1),
(3, 'B1', 1),
(3, 'B2',1),
(3, 'B3', 1),
(3, 'B4', 1),
(4, 'A1', 1),
(4, 'A2', 1),
(4, 'A3', 1),
(4, 'A4', 1),
(4, 'B1', 1),
(4, 'B2',1),
(4, 'B3', 1),
(4, 'B4', 1);



