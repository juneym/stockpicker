Stockpicker


This is a sample project for

    1) evaluating the use of Glide as a dependency management tool.
    2) building a "library" focused, reference go-code layout that allows for multiple binaries under `cmd` directory.

Coming from PHP, this tool is very familiar to "composer".


Note: you need to install Glide (http://glide.sh/) by downloading the pre-build binaries from their website or
running

    go get github.com/Masterminds/glide


To download the dependencies, run

    glide install

To compile,

    go build cmd/stockpicker/main.go

The above command will create a binary file called "main"

    ./main --help

To run (during development)

    go run cmd/stockpicker/main.go


