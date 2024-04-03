

rm -f go.sum
go get -u github.com/Patrick-ring-motive/utils
go get -u github.com/Patrick-ring-motive/traigo@0.0.2
go get -u github.com/Patrick-ring-motive/ione@0.0.2
go build -ldflags "-g" -gcflags="-B -v -std"  -o httpne httpne.go