# query-compensation-data
A read-only API (REST or GraphAPI) that returns one or more records from a static set of compensation data.

Instructions for setup
1. Clone the repository and change the directory to root of the cloned repository.
2. Make sure golang version >= 1.23.9
3. To setup postgres db and pgadmin, you need to run the command below in your terminal. 
`make compose-up`
4. Run the command below to download all the dependencies required by the go module.
`go mod download`
6. I have used dataset csv3. To read the records from the csv file and to write to sql table: issue the following command in your terminal ->    
    `go run cmd/ingest/main.go`
7. To start the server:
`go run cmd/compensation/server.go`
    The server would have started at "http://localhost:8080/comp_data/query". You can use cURL command on terminal or postman to make query requests. Eg.
compensation get query

`curl --location 'http://localhost:8080/comp_data/query' \
--header 'Content-Type: application/json' \
--data '{"query":"query {\n  compensation(id: \"1\") {\n    compensation {\n      id\n      company\n    }\n    error {\n      message\n    }\n  }\n}","variables":{}}'`

compensation list query

`curl --location 'http://localhost:8080/comp_data/query' \
--header 'Content-Type: application/json' \
--data '{"query":"query {\n  compensations(sortBy: TOTALCOMP_DESC) {\n    compensations {\n      id\n        timestamp\n        company\n        title\n        city\n        state\n        total_comp\n        signing_bonus\n        base_salary\n        annual_bonus\n        annual_stock_value\n        years_exp\n        additional_comments\n        gender\n        years_at_company\n    }\n    count\n    error {\n      message\n    }\n  }\n}","variables":{}}'`



The graphQL schema files can be found at the location: internal/compensation/handler/graph/compensation/schema/compensation
 
