Go E-Commerce API
A Go-based backend for an e-commerce platform. This API provides RESTful endpoints to manage products, users, orders, and integrates Stripe for payment processing. The application uses PostgreSQL for the database and Docker for both the application and database setup.

Prerequisites
Go (Programming Language)

PostgreSQL (Database)

Docker (For containerized setup)

Postman (For testing API endpoints)

Note: No ORM is used; the application interacts with PostgreSQL via plain SQL.

Setup
Environment Setup
Clone the repository:

bash
Copy
git clone https://github.com/AhmedMoniem96/Go-Ecomm.git
Build the Docker containers:

bash
Copy
docker-compose up --build
API Features
Admin Endpoints
Create Product

POST /products

Description: Create a new product.

Update Product

PUT /products/{id}

Description: Update an existing product by ID.

Delete Product

DELETE /products/{id}

Description: Delete an existing product by ID.

Get Product Sales

GET /products/admin/{username}

Description: Get sales data for a specific product based on the username.

User Endpoints
Get All Users

GET /users

Description: Retrieve all users.

Register User

POST /users/register

Description: Register a new user.

Login User

POST /users/login

Description: Login and receive a JWT token.

Add Product to Cart

POST /users/cart

Description: Add a product to the user's cart.

Get Cart Items

GET /users/cart

Description: Retrieve all items in the user's cart.

Add Order Product

POST /users/addOrder-product

Description: Add products to an order.

Add Order

POST /users/addOrder

Description: Complete an order by adding the products to the user's order history.

Add Credit Card

POST /users/add-credit

Description: Add a credit card to the user's account.

Delete Credit Card

DELETE /users/credit/{card_id}

Description: Delete a specific credit card by ID.

Get Order History

GET /users/history

Description: Retrieve the user's order history.

Product in Cart
Add Product to Cart

POST /addProduct-cart

Description: Add a product to the user's cart.

