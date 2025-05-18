# 📊 Query Compensation Data

A **read-only GraphQL API** that returns one or more records from a static set of compensation data.

---

## 🚀 Setup Instructions

### 1. Clone the repository
```bash
git clone <your-repo-url>
cd query-compensation-data
```

### 2. Prerequisites
- **Go version** `>= 1.23.9`

### 3. Start PostgreSQL and pgAdmin via Docker
```bash
make compose-up
```

### 4. Install Go dependencies
```bash
go mod download
```

### 5. Ingest data from CSV
The application uses dataset `salary_survey-3.csv`.  
To read and populate records in the database:
```bash
go run cmd/ingest/main.go
```

### 6. Start the GraphQL server
```bash
go run cmd/compensation/server.go
```

Once the server is running, access the endpoint at:

```
http://localhost:8080/comp_data/query
```

---

## 🔎 Example GraphQL Queries

You can use **Postman** or `curl` to send GraphQL requests to the API.

### 🧠 Fetch a Single Compensation by ID

```bash
curl --location 'http://localhost:8080/comp_data/query' --header 'Content-Type: application/json' --data '{
  "query": "query {\n  compensation(id: \"1\") {\n    compensation {\n      id\n      company\n    }\n    error {\n      message\n    }\n  }\n}",
  "variables": {}
}'
```

---

### 📄 List All Compensations (Sorted by Total Compensation Descending)

```bash
curl --location 'http://localhost:8080/comp_data/query' --header 'Content-Type: application/json' --data '{
  "query": "query {\n  compensations(sortBy: TOTALCOMP_DESC) {\n    compensations {\n      id\n      timestamp\n      company\n      title\n      city\n      state\n      total_comp\n      signing_bonus\n      base_salary\n      annual_bonus\n      annual_stock_value\n      years_exp\n      additional_comments\n      gender\n      years_at_company\n    }\n    count\n    error {\n      message\n    }\n  }\n}",
  "variables": {}
}'
```

---

## 📂 Schema Location

GraphQL schema definitions are available at:

```
internal/compensation/handler/graph/compensation/schema/compensation
```
