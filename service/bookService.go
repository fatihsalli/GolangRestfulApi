package service

import (
	"RestfulWithEcho/dtos"
	"RestfulWithEcho/models"
	"RestfulWithEcho/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BookService struct {
	Repository repository.IBookRepository
}

type IBookService interface {
	Insert(bookDto dtos.BookCreateDto) (bool, error)
	GetAll() ([]dtos.BookDto, error)
	GetBookById(id string) (dtos.BookDto, error)
	Update(bookDto dtos.BookUpdateDto) (bool, error)
	Delete(id string) (bool, error)
}

// NewBookService => to create new BookService
func NewBookService(repository repository.IBookRepository) BookService {
	return BookService{Repository: repository}
}

func (b BookService) Insert(bookDto dtos.BookCreateDto) (bool, error) {
	var book models.Book

	// we can use automapper, but it will cause performance loss.
	book.Title = bookDto.Title
	book.Quantity = bookDto.Quantity
	book.Author = bookDto.Author
	// to create id and created date value
	book.ID = uuid.New().String()
	book.CreatedDate = primitive.NewDateTimeFromTime(time.Now())

	result, err := b.Repository.Insert(book)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func (b BookService) GetAll() ([]dtos.BookDto, error) {
	result, err := b.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	var bookDto dtos.BookDto
	var booksDto []dtos.BookDto

	for _, v := range result {
		bookDto.ID = v.ID
		bookDto.Title = v.Title
		bookDto.Quantity = v.Quantity
		bookDto.Author = v.Author

		booksDto = append(booksDto, bookDto)
	}

	return booksDto, nil
}

func (b BookService) GetBookById(id string) (dtos.BookDto, error) {
	result, err := b.Repository.GetBookById(id)

	var bookDto dtos.BookDto

	if err != nil {
		return bookDto, err
	}

	bookDto.ID = result.ID
	bookDto.Title = result.Title
	bookDto.Author = result.Author
	bookDto.Quantity = result.Quantity

	return bookDto, nil
}

func (b BookService) Update(bookDto dtos.BookUpdateDto) (bool, error) {
	var book models.Book

	// we can use automapper, but it will cause performance loss.
	book.ID = bookDto.ID
	book.Title = bookDto.Title
	book.Quantity = bookDto.Quantity
	book.Author = bookDto.Author
	// to create updated date value
	book.UpdatedDate = primitive.NewDateTimeFromTime(time.Now())

	result, err := b.Repository.Update(book)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func (b BookService) Delete(id string) (bool, error) {
	result, err := b.Repository.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}
