# ****** ******* *********
# Golang Matrix API

Golang Matrix API is an application that exposes a simple API to perform operations in a given Matrix file.

## How to Run :fire:
The application can be executed using the command: `go run .`

## API

The app consists of 5 endpoints listed below:

***All requests are considering the follow input file matrix.csv:***

```
1,2,3
4,5,6
7,8,9
```

1. Echo -> `/echo`

    - Sending a request
        ```
        curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
        ```
    
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert -> `/invert`
    - Sending a request
        ```
        curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"
        ```

    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten -> `/flatten`
    - Sending a request
        ```
        curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"
        ```

    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Sending a request
        ```
        curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
        ```

    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Sending a request
        ```
        curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"
        ```

    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

The input file to these endpoints is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.  
