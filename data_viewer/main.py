import json
import time
import redis
import psycopg2
from psycopg2.extras import RealDictCursor

def main():
    # Connect to Redis
    r = redis.Redis(host='redis', port=6379, db=0, decode_responses=True)

    # Connect to PostgreSQL
    conn = psycopg2.connect(
        dbname='graphdb',
        user='postgres',
        password='password',
        host='db',
        port='5432'
    )
    
    while True:
        # Check for new data in Redis
        visited_nodes = r.get('visited_nodes')
        
        if visited_nodes:
            # Fetch node details from PostgreSQL
            with conn.cursor(cursor_factory=RealDictCursor) as cursor:
                query = "SELECT * FROM nodes"
                cursor.execute(query)
                nodes = cursor.fetchall()
                print(nodes)

            # Clear the Redis key
            r.delete('visited_nodes')
        
        time.sleep(1)  # Wait for 1 second before checking again

if __name__ == "__main__":
    main()
