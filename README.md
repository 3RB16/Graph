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
