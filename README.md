# Graph Processing Application

## Overview

This project is a distributed system for graph data processing using a combination of Go, Python, PostgreSQL, Redis, and ClickHouse. The Go application generates graph data, performs graph algorithms, and stores the results in PostgreSQL and ClickHouse. The Python application retrieves and visualizes the graph data for analysis.

## Components

1. **Go Application**: Generates and processes graph data, stores results in PostgreSQL and ClickHouse.
2. **Python Application**: Fetches graph data from PostgreSQL and displays it for analysis.
3. **PostgreSQL**: Relational database to store graph data and metadata.
4. **Redis**: Acts as a message queue for passing processed graph data between applications.
5. **ClickHouse**: Columnar database management system for fast analytical queries.
6. **PgAdmin**: A web-based interface to manage PostgreSQL databases.

## Architecture

```plaintext
+-----------------+
| Go Application  |
+-----------------+
         |
   [Processes Data]
         |
+-----------------+     +----------------+
| PostgreSQL      |<--->| ClickHouse     |
+-----------------+     +----------------+
         |
   [Message Queue]
         |
+-----------------+
| Redis           |
+-----------------+
         |
+-----------------+
| Python App      |
+-----------------+

## Prerequisites
Docker: Ensure Docker and Docker Compose are installed on your system.
Go: (if running outside Docker) Go v1.17 or higher.
Python: (if running outside Docker) Python 3.9 or higher.
Setup Instructions
1. Clone the Repository
bash
Copy code
git clone https://github.com/your-username/graph-processing-app.git
cd graph-processing-app
2. Configure Environment Variables
Create a .env file in the root directory with the following content:

env
Copy code
POSTGRES_DB=graphdb
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password

CLICKHOUSE_DB=graphdb
CLICKHOUSE_USER=default
CLICKHOUSE_PASSWORD=password
CLICKHOUSE_HOST=clickhouse
CLICKHOUSE_PORT=9000
3. Build and Run the Containers
bash
Copy code
docker-compose up --build
4. Access Services
PgAdmin: Open a web browser and go to http://localhost:5050. Login with your credentials (configured in the .env file or docker-compose.yml).
Redis CLI: Use the following command to access the Redis CLI within the container:
bash
Copy code
docker exec -it <redis-container-name> redis-cli
ClickHouse: Access ClickHouse with a compatible SQL client or directly from the command line using the ClickHouse client.
5. Checking Logs
To check logs for troubleshooting, use:

bash
Copy code
docker-compose logs -f
Usage
Go Application
The Go application runs a loop that generates and processes graph data every 10 seconds:

Graph Generation: Randomly generates a graph with nodes and edges.
Graph Processing: Applies a Depth-First Search (DFS) algorithm to the graph.
Data Storage: Stores graph nodes and edges in PostgreSQL and ClickHouse.
Message Queue: Sends processed graph data to Redis for the Python app.
Python Application
The Python application continuously listens for new graph data from Redis, fetches additional details from PostgreSQL, and displays the data for analysis.

PgAdmin and ClickHouse
PgAdmin: Manage and visualize your PostgreSQL data.
ClickHouse: Perform analytical queries on the graph data.
Troubleshooting
PgAdmin Shows Empty Database: Ensure PgAdmin is connected to the correct PostgreSQL database (graphdb).
Docker Compose Errors: Check your Docker Compose setup for missing dependencies or environment variables.
Redis Connection Issues: Ensure that the Redis service is running and accessible from both Go and Python applications.
ClickHouse Connectivity: Ensure ClickHouse is configured correctly and that the Go application can connect to it.
Extending the Project
You can add more tables and functionalities by modifying the Go and Python applications:

Go Application: Update main.go to include new data structures and algorithms.
Python Application: Update main.py to fetch and visualize additional data.
Docker Configuration: Update docker-compose.yml to include any new services or dependencies.

