package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

var contacts = []Contact{
	{ID: 1, Name: "Alice", Age: 30, Phone: "123-456-7890", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Age: 25, Phone: "098-765-4321", Email: "bob@example.com"},
}

var nextID = 0

// Função que verifica se um ID já existe na lista de contatos
func idExists(id int) bool {
	for _, contact := range contacts {
		if contact.ID == id {
			return true
		}
	}
	return false
}

func ListContacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, contacts)
}

func AddContact(c *gin.Context) {
	var newContact Contact

	if err := c.BindJSON(&newContact); err != nil {
		return
	}

	// Verifica se nextID já existe e incrementa até encontrar um ID único
	for idExists(newContact.ID) {
		nextID++
		newContact.ID = nextID
	}

	contacts = append(contacts, newContact)
	c.IndentedJSON(http.StatusCreated, newContact)
}

func EditContact(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	var updatedContact Contact
	if err := c.BindJSON(&updatedContact); err != nil {
		return
	}

	for i, contact := range contacts {
		if contact.ID == id {
			updatedContact.ID = id // Mantém o ID original
			contacts[i] = updatedContact
			c.IndentedJSON(http.StatusOK, updatedContact)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
}

func RemoveContact(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	for i, contact := range contacts {
		if contact.ID == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
}
