# Unit Converter
Sample solution for the [unit converter](https://roadmap.sh/projects/unit-converter) challenge from [roadmap.sh](https://roadmap.sh)

## Prerequisites
- Go should be installed on your system. If you haven't installed it yet, follow the instructions [here](https://golang.org/doc/install)

## How to Run
**1. Clone the repository** and run following command:
```bash
git clone https://github.com/David452/roadmapsh-backend.git
cd roadmapsh-backend/beginner-projects/unit-converter
```

**2. Run the application** with the following command:
```bash
go run main.go    # Run the project
```
Server will start listening on **port 8080** by default

**3. Access the Unit Converter**
- Open your web browser and navigate to `http://localhost:8080`.
- You should see a user interface for the unit converter

**4. Using the Converter**
- Select if you want to convert weight or length
- Select the units you want to convert (e.g., meters to feet) and enter the value to be converted 

## Notes
- Ensure you are in the correct directory (`roadmapsh-backend/beginner-projects/unit-converter`) when running the commands

## Troubleshooting
- If you encounter errors, ensure all dependencies are installed and your Go version is up to date.
- If you see a **"port already in use"** error, change the port in the `main.go` file:
```go
const (
    PORT = ":8080" // Change this to a different port, e.g., ":8081"
)
```
After changing the port number, save the file and run the application again with `go run main.go`