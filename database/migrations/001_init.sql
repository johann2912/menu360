-- Create users table
CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password TEXT NOT NULL,
        role VARCHAR(20) CHECK (role IN ('admin', 'employee', 'customer')) DEFAULT 'customer',
        created_at TIMESTAMP DEFAULT NOW ()
    );

-- Create restaurants table
CREATE TABLE
    IF NOT EXISTS restaurants (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        owner_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        status VARCHAR(20) CHECK (status IN ('active', 'closed')) DEFAULT 'active',
        created_at TIMESTAMP DEFAULT NOW ()
    );

-- Create menus table
CREATE TABLE
    IF NOT EXISTS menus (
        id SERIAL PRIMARY KEY,
        restaurant_id INT NOT NULL REFERENCES restaurants (id) ON DELETE CASCADE,
        name VARCHAR(100) NOT NULL,
        created_at TIMESTAMP DEFAULT NOW ()
    );

-- Create dishes table
CREATE TABLE
    IF NOT EXISTS dishes (
        id SERIAL PRIMARY KEY,
        menu_id INT NOT NULL REFERENCES menus (id) ON DELETE CASCADE,
        name VARCHAR(100) NOT NULL,
        price DECIMAL(10, 2) NOT NULL,
        created_at TIMESTAMP DEFAULT NOW ()
    );

-- Create orders table
CREATE TABLE
    IF NOT EXISTS orders (
        id SERIAL PRIMARY KEY,
        customer_id INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        restaurant_id INT NOT NULL REFERENCES restaurants (id) ON DELETE CASCADE,
        status VARCHAR(20) CHECK (
            status IN (
                'pending',
                'preparing',
                'ready',
                'delivered',
                'canceled'
            )
        ) DEFAULT 'pending',
        created_at TIMESTAMP DEFAULT NOW ()
    );

-- Create order details table
CREATE TABLE
    IF NOT EXISTS order_details (
        id SERIAL PRIMARY KEY,
        order_id INT NOT NULL REFERENCES orders (id) ON DELETE CASCADE,
        dish_id INT NOT NULL REFERENCES dishes (id) ON DELETE CASCADE,
        quantity INT NOT NULL CHECK (quantity > 0),
        created_at TIMESTAMP DEFAULT NOW ()
    );