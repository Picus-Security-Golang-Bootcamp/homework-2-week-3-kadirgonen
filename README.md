## Homework | Week 3

In this application, we provide a few status checks of the books in the application by making a book application.

## Purpose of application

We have a list of books.
The book areas are as follows;
```
- Book ID
- Book Name
- Paper Number
- Stock Number
- Cost
- Stock Code
- ISBN
- Author (ID and Name)
```

1. List all books (list)
2. List the books in which the given entry is in the titles of the books (search)
3. Print book by ID(get)
4. Delete the book with the given ID. (The deleted book should be coming by ID.) (delete)
5. Buy the book with the ID given as many as you want and print the last information of the book on the screen. (buy)
6. If the wrong command is entered, it will print the usage on the screen. 
 
## Used Commands

### list command
```
go run main.go list
```
This command allows us to see the objects we define in the application as a list.

### search command 
```
go run main.go search <bookName>
go run main.go search Harry Potter
```
This command allows us to scan the objects we have defined in the application to check if they exist inside the application.

### get command
```
go run main.go get <bookID>
go run main.go get 6
```
This command allows to bring us the objects used in the application and their contents.

### delete command
```
go run main.go delete <bookID>
go run main.go delete 4
```
This command allows us to delete objects used in the application.

### buy command
```
go run main.go buy <bookID> <quantity>
go run main.go buy 3 4
```
This command checks the stock count of objects used in the application. It checks the stock of the items we will purchase later and shows us the number of stocks after the purchase accordingly.

## Requirements

* Go Language
* Git
* Go Module

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc

### Writing About Concurrency: https://medium.com/@kadirgnen/golang-concurrency-goroutine-channel-157f5d2ba60f ( I will continue this writing)
