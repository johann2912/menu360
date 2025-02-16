---

## 📌 Menú360 - Restaurant Menu & Order Management API 🍽️  

Menú360 is a **REST API** that allows restaurants to manage their menus and receive customer orders in an organized and efficient way. Each restaurant operates independently, ensuring secure and isolated data access.  

### ✅ Features  
- **Multi-tenant:** Each restaurant manages its own information without interference.  
- **Real-time menu and order management.**  
- **Role-based access control (Admins, Employees, Customers).**  
- **Order status workflow (`pending`, `in preparation`, `ready`, `delivered`).**  
- **Secure authentication with JWT.**  



## 📌 Business Logic  

### 🔹 User Roles  

| **Role**               | **Description** |
|------------------------|----------------|
| **Super Admin**        | Full control over the platform. Can create and delete restaurants. |
| **Restaurant Admin**   | Manages the restaurant, creates menus, and handles orders. |
| **Employee**           | Processes and updates order statuses. |
| **Customer**           | Views restaurants, browses menus, and places orders. |


### 🔹 Business Rules  

1️⃣ **Anyone can register as a customer or restaurant owner.**  
2️⃣ **Customers can browse all restaurants and their menus.**  
3️⃣ **Each restaurant manages its own menu and orders independently.**  
4️⃣ **Customers can place orders at only one restaurant at a time.**  
5️⃣ **Orders follow these statuses:**  

   - **`pending`** → Order just placed.  
   - **`in preparation`** → Order is being cooked.  
   - **`ready`** → Order is ready for pickup/delivery.  
   - **`delivered`** → Order has been completed.  
   - **`canceled`** → Order canceled by the customer before preparation.  

6️⃣ **A restaurant CANNOT be deleted if it has pending orders.**  
7️⃣ **A restaurant can be "closed" to stop accepting new orders.**  
8️⃣ **Customers can only see their own orders.**  
9️⃣ **Restaurant admins can add employees.**  
🔟 **A customer can cancel an order ONLY if it is still `pending`.**  

## 📌 Customer Flow  

1️⃣ Registers on the platform.  
2️⃣ Browses available restaurants.  
3️⃣ Selects a restaurant and views its menu.  
4️⃣ Adds items to the cart and confirms the order.  
5️⃣ Tracks the order status in real time.  
6️⃣ Can cancel the order only if it is still `pending`.  

## 📌 Restaurant Admin Flow  

1️⃣ Registers on the platform and creates a restaurant.  
2️⃣ Adds employees and configures the menu.  
3️⃣ Receives and manages customer orders.  
4️⃣ Updates order statuses until they are delivered.  
5️⃣ Can close the restaurant to stop receiving orders.  
6️⃣ Can only delete the restaurant if there are no active orders.  

---

## 📌 Implemented Architecture

Menú360 follows a circular architecture (also known as clean or hexagonal architecture) to maintain a clear separation between the application layers and facilitate scalability and maintenance. Below is a diagram of the architecture:

![Circular Architecture](./documentation/images/CleanArchitecture.jpg)

---

## 📌 Folder Structure

```
menu360/
│── cmd/                                              # 📌 Application entry point
│   ├── main.go                                       # Starts the server and loads modules
│── config/                                           # 📌 Global configuration
│   ├── database.go                                   # Connection to PostgreSQL or MongoDB
│   ├── env.go                                        # Loads environment variables
│── database/                   
│   ├── migrations/                                   # 📌 SQL scripts for the database
│   │   ├── 001_init.sql                              # First migration with main tables
│   ├── models/                                       # 📌 Database models (Domain Entities)
│   │   ├── user.go                                   # User model
│   │   ├── restaurant.go                             # Restaurant model
│   │   ├── menu.go                                   # Menu model
│   │   ├── order.go                                  # Order model
│   ├── repository/                                   # 📌 Concrete repository implementations
│   │   ├── user_repository_postgres.go               # PostgreSQL implementation
│   │   ├── user_repository_mongo.go                  # MongoDB implementation
│   │   ├── restaurant_repository_postgres.go         # PostgreSQL implementation
│   │   ├── restaurant_repository_mongo.go            # MongoDB implementation
│   │   ├── menu_repository_postgres.go               # PostgreSQL implementation
│   │   ├── menu_repository_mongo.go                  # MongoDB implementation
│   │   ├── order_repository_postgres.go              # PostgreSQL implementation
│   │   ├── order_repository_mongo.go                 # MongoDB implementation
│── internal/                   
│   ├── user/                   
│   │   ├── repository.go                             # 📌 Repository interface (Contract)
│   │   ├── usecase/             
│   │   │   ├── create_user.go                        # Use case: Create user
│   │   │   ├── login_user.go                         # Use case: User login
│   │   │   ├── get_user.go                           # Use case: Get user
│   │   ├── module.go                                 # Module configuration
│   │   ├── controller.go                             # HTTP controller
│   │   ├── routes.go                                 # Route definitions
│   ├── restaurant/             
│   │   ├── repository.go                             # 📌 Repository interface (Contract)
│   │   ├── usecase/             
│   │   │   ├── create_restaurant.go 
│   │   │   ├── close_restaurant.go
│   │   ├── module.go           
│   │   ├── controller.go       
│   │   ├── routes.go           
│   ├── menu/                   
│   │   ├── repository.go                              # 📌 Repository interface (Contract)
│   │   ├── usecase/
│   │   │   ├── create_menu.go
│   │   │   ├── update_menu.go
│   │   ├── module.go           
│   │   ├── controller.go       
│   │   ├── routes.go           
│   ├── order/                  
│   │   ├── repository.go                             # 📌 Repository interface (Contract)
│   │   ├── usecase/
│   │   │   ├── create_order.go
│   │   │   ├── cancel_order.go
│   ├── module.go           
│   ├── controller.go       
│   ├── routes.go           
│   ├── middleware/             
│   │   ├── auth.go                                   # Authentication middleware (JWT)
│   │   ├── role.go                                   # Role middleware
│── router/                     
│   ├── router.go                                     # Global router
│   ├── user_routes.go                                # User module routes
│   ├── restaurant_routes.go                          # Restaurant module routes
│   ├── menu_routes.go                                # Menu module routes
│   ├── order_routes.go                               # Order module routes
│── pkg/                        
│   ├── logger.go                                     # Centralized logging
│── test/                                             # 📂 Unit tests
│── documentation/                                    # 📌 Documentation folder
│── go.mod                                            # 📌 Go module
│── Makefile                                          # 📌 Automation scripts
│── Dockerfile                                        # 📌 Container configuration
│── docker-compose.yml                                # 📌 Stack with PostgreSQL and migrations
│── .env                                              # 📌 Environment variables
│── README.md                                         # 📌 Project documentation
 ```
---
