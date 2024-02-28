# Implementation of key-value storage like Redis.

The storage is implemented based on the map structure, which allows for the highest speed of operation. 
In order to implement multi-threaded access to the storage, an mutex is also added to the structure. 
Automatic deletion from storage is implemented using a timer. 
The code is covered with unit tests, you can also run and view it using `go run .`