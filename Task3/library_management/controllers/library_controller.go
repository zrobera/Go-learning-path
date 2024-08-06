package controllers

import (
    "fmt"
    "library_management/models"
    "library_management/services"
	"bufio"
	"strings"
	"strconv"
	"os"
)

func RunLibraryManagement() {
    var library services.LibraryManager = services.NewLibrary()
	reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("Library Management System")
        fmt.Println("1. Add Book")
        fmt.Println("2. Remove Book")
        fmt.Println("3. Borrow Book")
        fmt.Println("4. Return Book")
        fmt.Println("5. List Available Books")
        fmt.Println("6. List Borrowed Books by Member")
        fmt.Print("Enter your choice number: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            addBook(library,reader)
        case 2:
            removeBook(library)
        case 3:
            borrowBook(library)
        case 4:
            returnBook(library)
        case 5:
            listAvailableBooks(library)
        case 6:
            listBorrowedBooks(library)
        default:
            fmt.Println("Invalid choice, please try again.")
        }
    }
}

func addBook(library services.LibraryManager, reader *bufio.Reader) {
    fmt.Print("Enter book ID: ")
    idStr, _ := reader.ReadString('\n')
    idStr = strings.TrimSpace(idStr)
    id, _ := strconv.Atoi(idStr)

    fmt.Print("Enter book title: ")
    title, _ := reader.ReadString('\n')
    title = strings.TrimSpace(title)

    fmt.Print("Enter book author: ")
    author, _ := reader.ReadString('\n')
    author = strings.TrimSpace(author)

    book := models.Book{
        ID:     id,
        Title:  title,
        Author: author,
        Status: "Available",
    }
    library.AddBook(book)
    fmt.Println("Book added successfully!")
    fmt.Println("------------------------------------------------------------")
    fmt.Println("------------------------------------------------------------")
}

func removeBook(library services.LibraryManager) {
    var id int
    fmt.Print("Enter book ID to remove: ")
    fmt.Scan(&id)
    library.RemoveBook(id)
    fmt.Println("Book removed successfully!")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("------------------------------------------------------------")
}

func borrowBook(library services.LibraryManager) {
    var bookID, memberID int
    fmt.Print("Enter book ID to borrow: ")
    fmt.Scan(&bookID)
    fmt.Print("Enter member ID: ")
    fmt.Scan(&memberID)

    err := library.BorrowBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book borrowed successfully!")
    }
	fmt.Println("------------------------------------------------------------")
	fmt.Println("------------------------------------------------------------")
}

func returnBook(library services.LibraryManager) {
    var bookID, memberID int
    fmt.Print("Enter book ID to return: ")
    fmt.Scan(&bookID)
    fmt.Print("Enter member ID: ")
    fmt.Scan(&memberID)

    err := library.ReturnBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book returned successfully!")
    }
	fmt.Println("------------------------------------------------------------")
	fmt.Println("------------------------------------------------------------")
}

func listAvailableBooks(library services.LibraryManager) {
    books := library.ListAvailableBooks()
    if len(books) == 0 {
        fmt.Println("No available books.")
        return
    }
    fmt.Println("Available Books:")
    for _, book := range books {
        fmt.Printf("ID: %v, Title: %v, Author: %v\n", book.ID, book.Title, book.Author)
    }
	fmt.Println("------------------------------------------------------------")
	fmt.Println("------------------------------------------------------------")
}

func listBorrowedBooks(library services.LibraryManager) {
    var memberID int
    fmt.Print("Enter member ID: ")
    fmt.Scan(&memberID)

    books := library.ListBorrowedBooks(memberID)
    if len(books) == 0 {
        fmt.Println("No borrowed books.")
        return
    }
    fmt.Println("Borrowed Books:")
    for _, book := range books {
        fmt.Printf("ID: %v, Title: %v, Author: %v\n", book.ID, book.Title, book.Author)
    }
	fmt.Println("------------------------------------------------------------")
	fmt.Println("------------------------------------------------------------")
}
