package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	BookID      int     `json:"book_id"`
	Names       string  `json:"names"`
	PaperNumber int     `json:"paper_number"`
	StockNumber int     `json:"stock_number"`
	Cost        float64 `json:"cost"`
	StockCode   int     `json:"stock_code"`
	Isbn        int     `json:"ISBN"`
	Author      string  `json:"author"`
	Flag        bool    `json:"flag"`
}

func main() {

	bookJson := `[{"book_id":1,"names":"Lord of The Rings", "paper_number":423, "stock_number":5, "cost":4.99, "stock_code":4231245, "ISBN":946785423,"author":"John Tolkien", "flag":""},{"book_id":2, "names":"Harry Potter", "paper_number":432, "stock_number":4, "cost":7.12, "stock_code":4231325, "ISBN":946785432,"author":"Joanne Rowling", "flag":""},{"book_id":3, "names":"Hunger Games", "paper_number":339, "stock_number":7, "cost":9.23, "stock_code":4231345, "ISBN":946785465,"author":"Suzanne Collins", "flag":""},{"book_id":4, "names":"Treasure Island", "paper_number":312, "stock_number":8, "cost":8.26, "stock_code":4231254, "ISBN":946785476,"author":"Robert Stevenson", "flag":""},{"book_id":5, "names":"Moonlight", "paper_number":326, "stock_number":6, "cost":9.28, "stock_code":4231268, "ISBN":946785452,"author":"Harold Pinter", "flag":""},{"book_id":6, "names":"Oliver Twist", "paper_number":228, "stock_number":3, "cost":3.55, "stock_code":4231269, "ISBN":946785473,"author":"Charles Dickens", "flag":""}]`

	var books []Book
	json.Unmarshal([]byte(bookJson), &books)
	if os.Args[1] == "list" {

		ListBookNames(books)

	} else if os.Args[1] == "search" {

		var searchArg []string = os.Args[2:]
		fmt.Printf("Names : %+v", searchArg)

		args := os.Args[2:]
		for i, arg := range args {

			fmt.Printf("index: %d, value: %s\n", i, arg)
		}
		SearchBookName(books, strings.Join(searchArg, " "))

	} else if os.Args[1] == "get" {

		/*var getArg []string = os.Args[2:]
		fmt.Printf("book_id : %+v", getArg)*/

		args := os.Args[2:]
		for i, arg := range args {

			fmt.Printf("index: %d, value: %s\n", i, arg)
		}
		i1, err := strconv.Atoi(os.Args[2])
		if err == nil {
			fmt.Println(i1)
		}

		GetBookID(books, i1)
	} else if os.Args[1] == "buy" {

		/*var getArg []string = os.Args[2:]
		fmt.Printf("book_stock : %+v", getArg)*/

		args := os.Args[2:]
		for i, arg := range args {

			fmt.Printf("index: %d, value: %s\n", i, arg)
		}
		i1, err := strconv.Atoi(os.Args[2])
		if err == nil {
			fmt.Println(i1)
		}
		i2, err := strconv.Atoi(os.Args[3])
		if err == nil {
			fmt.Println(i2)
		}
		GetBookStock(books, i1, i2)
	} else if os.Args[1] == "delete" {

		/*var searchArg []string = os.Args[2:]
		fmt.Printf("Names : %+v", searchArg)*/

		args := os.Args[2:]
		for i, arg := range args {

			fmt.Printf("index: %d, value: %s\n", i, arg)
		}
		i1, err := strconv.Atoi(os.Args[2])
		if err == nil {
			fmt.Println(i1)
		}

		DeleteBookID(books, i1)
	} else {
		fmt.Println("Entered command is not valid!")
		usage()
	}
}
func DeleteBookID(books []Book, bookID int) {
	var bVal bool // default is false
	for i := range books {
		//fmt.Printf("Names : %+v", "bookID")
		if books[i].BookID == bookID {
			if !books[i].Flag {
				bVal = true
				books[i].Flag = true
				fmt.Printf("book : %+v", books[i])
				break
			}
		}
	}
	if bVal {
		fmt.Printf("book : %+v", strconv.Itoa(bookID)+" is Deleted ")
	} else {
		fmt.Printf("book : %+v", strconv.Itoa(bookID)+" Not Found")
	}

}
func GetBookStock(books []Book, bookID int, stockNumber int) {
	var bVal bool // default is false
	for i := range books {
		//fmt.Printf("Names : %+v", "stockNumber")
		if books[i].BookID == bookID {
			if books[i].StockNumber >= stockNumber {
				bVal = true
				fmt.Printf("book : %+v", books[i].StockNumber-stockNumber)
				break
			}
		}
	}
	if !bVal {
		fmt.Printf("book : %+v", strconv.Itoa(stockNumber)+" Stock Not Found")
	}

}

func GetBookID(books []Book, bookID int) {
	var bVal bool // default is false
	for i := range books {
		//fmt.Printf("Names : %+v", "bookID")
		if books[i].BookID == bookID {
			bVal = true
			fmt.Printf("book : %+v", books[i])
		}
	}
	if !bVal {
		fmt.Printf("book : %+v", strconv.Itoa(bookID)+" Not Found")
	}

}

func SearchBookName(books []Book, bookName string) {
	var bVal bool // default is false
	var SearchResult []string
	for i := range books {
		//fmt.Printf("Names : %+v", "bookName")
		if strings.Contains(strings.ToLower(books[i].Names), strings.ToLower(bookName)) {
			bVal = true
			SearchResult = append(SearchResult, books[i].Names)
			//fmt.Printf("Names : %+v", bookName)
		}
	}
	if bVal {
		fmt.Printf("Names : %+v", SearchResult)

	} else {
		fmt.Printf("Names : %+v", strings.ToLower(bookName)+" Not Found")
	}

}

func ListBookNames(books []Book) {

	for _, s := range books {
		fmt.Println(s)
	}

}

func usage() {
	fmt.Println("list: list books")
	fmt.Println("search: search a book")
	fmt.Println("get: get a book")
	fmt.Println("delete: delete a book")
	fmt.Println("buy: make a stock control")

}
