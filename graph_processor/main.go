package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "time"

    "github.com/go-redis/redis/v8"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "golang.org/x/net/context"

)

type Node struct {
    ID   int    `gorm:"primaryKey"`
    Name string
}

type Graph struct {
    nodes map[int]string
    edges map[int][]int
}

func NewGraph() *Graph {
    return &Graph{
        nodes: make(map[int]string),
        edges: make(map[int][]int),
    }
}

func (g *Graph) AddNode(id int, name string) {
    g.nodes[id] = name
}

func (g *Graph) AddEdge(from, to int) {
    g.edges[from] = append(g.edges[from], to)
    // For undirected graph, add the reverse edge
    g.edges[to] = append(g.edges[to], from)
}

func (g *Graph) Neighbors(node int) []int {
    return g.edges[node]
}

func (g *Graph) DFS(start int, visited map[int]bool, visitFunc func(int)) {
    if visited[start] {
        return
    }
    visited[start] = true
    visitFunc(start)
    for _, neighbor := range g.Neighbors(start) {
        g.DFS(neighbor, visited, visitFunc)
    }
}

func generateUniqueID(db *gorm.DB) int {
    rand.Seed(time.Now().UnixNano())
    for {
        id := rand.Intn(1000) // Generate a random ID between 0 and 999
        var count int64
        db.Model(&Node{}).Where("id = ?", id).Count(&count)
        if count == 0 {
            return id
        }
    }
}

func processGraph(db *gorm.DB, rdb *redis.Client, ctx context.Context) {
    // Create a new graph
    g := NewGraph()

    // Generate 10 unique nodes
    for i := 0; i < 10; i++ {
        id := generateUniqueID(db)
        name := fmt.Sprintf("Node%d", id)
        g.AddNode(id, name)
        db.Create(&Node{ID: id, Name: name})
    }

    // Generate random edges
    var nodeIDs []int
    db.Model(&Node{}).Pluck("id", &nodeIDs)

    for i := 0; i < len(nodeIDs); i++ {
        for j := i + 1; j < len(nodeIDs); j++ {
            if rand.Float32() < 0.5 { // 50% chance to add an edge
                g.AddEdge(nodeIDs[i], nodeIDs[j])
            }
        }
    }

    // Apply a simple graph algorithm (DFS)
    visited := make(map[int]bool)
    if len(nodeIDs) > 0 {
        g.DFS(nodeIDs[0], visited, func(node int) {
            fmt.Printf("Visited: %v\n", g.nodes[node])
        })
    }

    // Send visited nodes to Redis
    visitedNodes, err := json.Marshal(visited)
    if err != nil {
        log.Fatal(err)
    }
    err = rdb.Set(ctx, "visited_nodes", visitedNodes, 0).Err()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Graph processing complete. Data sent to Redis.")
}

func main() {
    // Connect to PostgreSQL
    dsn := "host=db user=postgres password=password dbname=graphdb port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    // Migrate the schema
    db.AutoMigrate(&Node{})

    // Connect to Redis
    rdb := redis.NewClient(&redis.Options{
        Addr: "redis:6379",
    })

    ctx := context.Background()

    for {
        processGraph(db, rdb, ctx)
        time.Sleep(10 * time.Second) // Wait for 10 seconds before repeating
    }
}
