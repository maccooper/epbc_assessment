
# API Endpoints

API endpoints built for transactional database provided within EPBC Homework assignment, tables for database can be found within the associated pdf in the repository.

Goal: Built CRUD Operations for each table. [Order, LineItem, Customer, CustomerAddress, ReturnItem]


## Order
| Operation  | Method | Endpoint| Description|
| ------------- |:-------------:|:-------------:|:-------------:|
| Create        | POST     | /orders     | Create a new order     |
| Read All      | GET     | /orders     | Get all orders     |
| Read One      | GET     | /orders/:id     | Get order by ID    |
| Update        | PUT     | /orders/:id     | Update an existing order     |
| Delete        | DELETE     | /orders/:id     | Remove an existing order     |

## LineItem
| Operation  | Method | Endpoint| Description|
| ------------- |:-------------:|:-------------:|:-------------:|
| Create        | POST     | /orders/:orderid/lineitems     |Create a new line item at specified order|
| Read All      | GET     | /lineitems     |Get all line items |
| Read          | GET     | /lineitems/:id     | Get line item by ID    |
| Update        | PUT     | /lineitems/:id     | Update an existing line item    |
| Delete        | DELETE     | /lineitems/:id     | Remove an existing line item     |

## Customer
| Operation  | Method | Endpoint| Description|
| ------------- |:-------------:|:-------------:|:-------------:|
| Create        | POST     | /customers     | Create a new customer     |
| Read All      | GET     | /customers     | Get all customers     |
| Read One      | GET     | /customers/:id     | Get customer by ID     |
| Update        | PUT     | /customers/:id     | Update an existing customer    |
| Delete        | DELETE     | /customers/:id     | remove an existing customer     |

## Customer Address
| Operation  | Method | Endpoint| Description|
| ------------- |:-------------:|:-------------:|:-------------:|
| Create        | POST     | /customers/:customerid/customeraddresses     | Create an address for a customer     |
| Read All      | GET     | /customeraddresses     | Get all customer addresses    |
| Read One      | GET     | /customeraddresses/:id     | Get customer address by ID      |
| Update        | PUT     | /customeraddresses/:id     | Update an existing address     |
| Delete        | DELETE     | /customeraddresses/:id     | remove an existing address     |

## Return Item
| Operation  | Method | Endpoint| Description|
| ------------- |:-------------:|:-------------:|:-------------:|
| Create        | POST     | /returns     | Create a new return     |
| Read All      | GET     | /returns     | Get all returns     |
| Read One      | GET     | /returns/:id     | Get a return by id     |
| Update        | PUT     | /returns/:id     | Update an existing return     |
| Delete        | DELETE     | /returns/:id     | Remove an existing return     |
