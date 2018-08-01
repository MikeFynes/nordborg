#Nordborg

A small town in Sønderjylland, Denmark


## What to run
Run the various tests

    cd src
    go test ./... -v

## Running from intellij

sliding_window_test.go requires only the sliding_window.go

csv_util_test.go requires both csv_util.go and sliding_window.go

## No main file
I have not implemented a main file as all functionality is visible through the tests
