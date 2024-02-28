run: 
	rm -rf cmd/lab/lab && go build -o cmd/lab/lab cmd/lab/main.go && cmd/lab/lab
