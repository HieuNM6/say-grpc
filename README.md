# say-grpc
Example project using grpc, flite, golang  

### How to run project
1. Clone, go to cloned directory and run command `go get` to get all dependencies
2. Build docker image by go to `backend` directory and run `make build`
3. Start docker backend server `docker run -p 8080:8080 campoy/say`
4. Start client by go to `say` directory and run `go run main.go "content want to sent to flite"`,
using `go run main.go -h` to get more information about command
5. Check output generated
