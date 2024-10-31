package main

import "gerenciador_contatos/contact"

func main() {
	nC1 := contact.Contact{Name: "Test1", Age: 31, Phone: "123456", Email: "test1@email.com"}
	nC2 := contact.Contact{Name: "Test2", Age: 32, Phone: "654321", Email: "test2@email.com"}
	nC3 := contact.Contact{Name: "Test3", Age: 33, Phone: "456123", Email: "test3@email.com"}
	nC4 := contact.Contact{Name: "Test4", Age: 34, Phone: "321654", Email: "test4@email.com"}

	contact.AddContact(nC1)
	contact.AddContact(nC2)
	contact.AddContact(nC3)

	contact.ListContacts()

	contact.EditContact(2, nC4)

	contact.ListContacts()

	contact.RemoveContact(2)

	contact.ListContacts()

	contact.RemoveContactByObject(nC1)

	contact.ListContacts()
}