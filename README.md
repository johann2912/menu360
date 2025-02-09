# menu360
Menú360 es una plataforma para la gestión de menús y pedidos en restaurantes, que permite a diferentes restaurantes gestionar sus menús, recibir pedidos y gestionar su funcionamiento de forma independiente. Esto es un proyecto didactico personal que tiene como reforzar conocimientos con Golang

📌 Menú360 - Restaurant Menu & Order Management API 🍽️
Menú360 is a REST API that allows restaurants to manage their menus and receive customer orders in an organized and efficient way. Each restaurant operates independently, ensuring secure and isolated data access.

✅ Multi-tenant: Each restaurant manages its own information without interference.
✅ Real-time menu and order management.
✅ Role-based access control (Admins, Employees, Customers).
✅ Order status workflow (pending, in preparation, ready, delivered).
✅ Secure authentication with JWT.

📌 Business Logic
🔹 User Roles
Role	                Description
Super Admin          	Full control over the platform. Can create and delete restaurants.
Restaurant            Admin	Manages the restaurant, creates menus, and handles orders.
Employee	            Processes and updates order statuses.
Customer	            Views restaurants, browses menus, and places orders.

Business Rules
  1️⃣ Anyone can register as a customer or restaurant owner.
  2️⃣ Customers can browse all restaurants and their menus.
  3️⃣ Each restaurant manages its own menu and orders independently.
  4️⃣ Customers can place orders at only one restaurant at a time.
  5️⃣ Orders follow these statuses:

pending → Order just placed.
in preparation → Order is being cooked.
ready → Order is ready for pickup/delivery.
delivered → Order has been completed.
canceled → Order canceled by the customer before preparation.
  6️⃣ A restaurant CANNOT be deleted if it has pending orders.
  7️⃣ A restaurant can be "closed" to stop accepting new orders.
  8️⃣ Customers can only see their own orders.
  9️⃣ Restaurant admins can add employees.
  🔟 A customer can cancel an order ONLY if it is still pending.

📌 Customer Flow
1️⃣ Registers on the platform.
2️⃣ Browses available restaurants.
3️⃣ Selects a restaurant and views its menu.
4️⃣ Adds items to the cart and confirms the order.
5️⃣ Tracks the order status in real time.
6️⃣ Can cancel the order only if it is still pending.

📌 Restaurant Admin Flow
1️⃣ Registers on the platform and creates a restaurant.
2️⃣ Adds employees and configures the menu.
3️⃣ Receives and manages customer orders.
4️⃣ Updates order statuses until they are delivered.
5️⃣ Can close the restaurant to stop receiving orders.
6️⃣ Can only delete the restaurant if there are no active orders.
