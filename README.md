# GoSort

GoSort is a collection of sorting algorithms implemented in Go.

## Prerequisites

Ensure you have Go installed on your machine. If not, you can download and install it from https://golang.org/dl/.

## How to Run

1. Navigate to the project folder

2. Compile the program:

``` bash
go build .
```

This will generate an executable file named `gosort.exe` in the directory.

3. Run the program using the generated executable. You can provide various flags to customize the array generation and sorting:

``` bash
gosort.exe -size=50 -max=1000 -algo=quicksort
```

OR

``` bash
gosort.exe -algo=quicksort 34 67 23 90 12 45 78 56 19 89
```

### Flags

- `-size`: This flag determines the size of the array to be generated. Default value is `10`. (Ignored if specific integer values are provided)

  Example:

  ``` bash
  gosort.exe -size=30
  ```

  This will generate an array of size 30.

- `-max`: This flag determines the maximum value that can be present in the array. Default value is `100`. (Ignored if specific integer values are provided)

  Example:

  ``` bash
  gosort.exe -max=500
  ```

  This will generate an array with numbers ranging from 0 to 499.

- `-algo`: This flag lets you choose the sorting algorithm. You can select either `quicksort` or `mergesort`. Default value is `quicksort`.

  Examples:

  ``` bash
  gosort.exe -algo=quicksort
  gosort.exe -algo=mergesort
  ```

4. The program will print the unsorted array followed by the sorted array.

## Example Output

Given an execution like:

``` bash
gosort.exe -size=10 -max=100 -algo=mergesort
```

OR

``` bash
gosort.exe -algo=mergesort 34 67 23 90 12 45 78 56 19 89
```

You might get an output like:

```
Unsorted array: [34 67 23 90 12 45 78 56 19 89]
Sorted array: [12 19 23 34 45 56 67 78 89 90]
```

(Note: The actual numbers in the unsorted array will differ with each execution since they are randomly generated.)

## Testing, Coverage, and Benchmarks

### Running Tests

To execute tests for the GoSort project, navigate to the project directory and run:

``` bash
go test ./...
```

This will run tests for all packages in the project.

### Test Coverage

To evaluate the code coverage of the tests:

``` bash
go test -cover ./...
```

This will provide a percentage representation of the code covered by the tests. For a more detailed view, you can generate a coverage profile and view it using a web browser:

``` bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

This will open up a browser displaying the covered and uncovered sections of your code.

### Benchmarks

To evaluate the performance of the sorting algorithms:

``` bash
go test -bench=.
```

This will run all benchmarks for the project and display the results, including the operations per second for each tested function.