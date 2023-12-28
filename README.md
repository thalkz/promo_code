# Promo Code

A API to create and validate promo codes, written in Go

## How to start server
- Go 1.21 required
- clone repository
- `cd promo_code`
- `go run main.go` to start serving on :8080

## How to Run tests
- `go test ./...` from root folder to run all tests
- `go test ./promocode` to run promocode unit tests
- `go test ./endpoints` to run endpoints integration tests

## Comments for the reviewers
Hi 👋, I'm going to stop now, even if the project is not fully completed yet. I've worked about 2 hours this morning, then did a 1h break and finally about 3 hours this afternoon. In total, that's about 5 hours.

Overall, I struggled a bit with handeling JSONs. The parsing and validating is a bit more involved than in Node.js, for example. I also took a long time before finding a good representation for `AgeRestriction`. I finally used pointers to ints, which works well, I think.

### What is done
- The "promocode" package parses & validates promocodes
- Both the parsing & validating are unit tested, but I didn't have time to cover all corner-cases
- Two routes are served at port 8080:
    - PUT /add to add new promocodes to the database
    - GET /verify to see if the promocode exists and check if the arguments are accepeted
- The database stores promocodes by name, and is a global variable

### What I would add with more time
- Some comments to explain the purpose of the main structs/functions
- Each endpoint has only 1 integration test. I would like to test different cases, including error responses
- There is a low of duplication in the tests, with in the testdata and the testing functions: this could be cleaned up
- Replace the placeholder for the weather API call
- Replace the database implementation & avoid using a global variable (error prone for the tests)
- Separate the promocode tests into their own package (promocode_package)
- Add authentification to the /verify endpoint (to restrict access to the Marketing team)
- Write documentation (README)

### Other ideas for later on
- Add a Dockerfile, and instructions on how to build it
- Add CI/CD (via Github actions, for example)
- Add test coverage badges
- Use an OpenAPI package package to generate the API docs from docstrings

## Tasks

### Milestone A - PromoCodes
- [x] Create PromoCode structs
- [x] Parse PromoCode from json
- [x] Create Argument struct
- [x] Parse Argument from json
- [x] Create validation logic
- [ ] Handle errors

### Milestone B - Call Weather API 
- [ ] Read openweathermap docs
- [ ] Implement request & response structs + parsing
- [ ] Handle errors

### Milestone C - Server
- [x] Import Gin framework + basic setup
- [x] Create router + empty routes
- [x] Add in memory database
- [x] Implement /add
- [x] Implement /verify

### Milestone D - Security
- [ ] Add auth method for /add route

### Milestone E - Cleanup & docs
- [ ] Write instructions in README
- [ ] Verify comments & docstrings
- [ ] Verify project architecture

### Bonus
- [ ] Fuzzing tests
- [ ] Contenairization
- [ ] Test coverage