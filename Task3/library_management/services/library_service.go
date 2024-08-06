package services

import (
    "errors"
    "library_management/models"
)

type LibraryManager interface {
    AddBook(book models.Book)
    RemoveBook(bookID int)
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []models.Book
    ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
    Books map[int]models.Book
    Members map[int]models.Member
}

func NewLibrary() *Library {
    return &Library{
        Books: make(map[int]models.Book),
        Members: make(map[int]models.Member),
    }
}

func (l *Library) AddBook(book models.Book){
    l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int){
    delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error{
    book, bookExists := l.Books[bookID]
    member, memberExists := l.Members[memberID]

    if !bookExists {
        return errors.New("book does not exist")
    }
    if !memberExists {
        return errors.New("member does not exist")
    }
    
    if book.Status == "Borrowed" {
        return errors.New("the book has already been borrowed")
    }
    book.Status = "Borrowed"
    l.Books[bookID] = book

    member.BorrowedBooks = append(member.BorrowedBooks, book)
    l.Members[memberID] = member

    return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
    book, bookExists := l.Books[bookID]
    member, memberExists := l.Members[memberID]

    if !bookExists {
        return errors.New("book does not exist")
    }
    if !memberExists {
        return errors.New("member does not exist")
    }
    
    if book.Status == "Available" {
        return errors.New("the book has never been borrowed")
    }

    book.Status = "Available"
    l.Books[bookID] = book

    var newBorrowedBooks []models.Book
    for _,borrowedBook := range member.BorrowedBooks {
        if borrowedBook != book {
            newBorrowedBooks = append(newBorrowedBooks, borrowedBook)
        }
    }

    member.BorrowedBooks = newBorrowedBooks
    l.Members[memberID] = member

    return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
    var availableBooks []models.Book
    for _, book := range l.Books {
        if book.Status == "Available" {
            availableBooks = append(availableBooks, book)
        }
    }

    return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
    member, memberExists := l.Members[memberID]
    if !memberExists {
        return nil
    }
    return member.BorrowedBooks
}
