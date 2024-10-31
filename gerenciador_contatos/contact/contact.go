package contact

import "fmt"

type Contact struct {
	Name  string
	Age   int
	Phone string
	Email string
}

var Contacts []Contact

func AddContact(c Contact) {
	fmt.Println("[Added]")
	fmt.Println("Name: ", c.Name, "Age: ", c.Age, "Phone: ", c.Phone, "E-mail: ", c.Email)
	Contacts = append(Contacts, c)
}

func ListContacts() {
	fmt.Println("[Listing]")
	for i := 0; i < len(Contacts); i++ {
		fmt.Println("Name: ", Contacts[i].Name, "Age: ", Contacts[i].Age, "Phone: ", Contacts[i].Phone, "E-mail: ", Contacts[i].Email)
	}
}

func EditContact(index int, c Contact) {
	if index >= 0 && index < len(Contacts) {
		fmt.Println("[Edited]")
		fmt.Println("Name: ", Contacts[index].Name, "Age: ", Contacts[index].Age, "Phone: ", Contacts[index].Phone, "E-mail: ", Contacts[index].Email)
		Contacts[index] = c
		fmt.Println("[Para]")
		fmt.Println("Name: ", c.Name, "Age: ", c.Age, "Phone: ", c.Phone, "E-mail: ", c.Email)
	}
}

func RemoveContact(index int) {
	if index >= 0 && index < len(Contacts) {
		fmt.Println("[Removed]")
		fmt.Println("Name: ", Contacts[index].Name, "Age: ", Contacts[index].Age, "Phone: ", Contacts[index].Phone, "E-mail: ", Contacts[index].Email)
		Contacts = append(Contacts[:index], Contacts[index+1:]...)
	}
}

func RemoveContactByObject(c Contact) {
    for i, contact := range Contacts {
        if contact == c {
            fmt.Println("[Removed]")
            fmt.Println("Name: ", contact.Name, "Age: ", contact.Age, "Phone: ", contact.Phone, "E-mail: ", contact.Email)
            Contacts = append(Contacts[:i], Contacts[i+1:]...)
            break
        }
    }
}
