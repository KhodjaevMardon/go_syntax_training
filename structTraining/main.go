package main

type Contact struct {
	Id        int
	FirstName string
	LastName  string
	Phone     uint64
	Email     string
	position  string
}

type ContactList struct {
	list []Contact
}

func (this *ContactList) CreateContact(contact Contact) {
	this.list = append(this.list, contact)
}

func (this *ContactList) DeleteById(id int) {
	index := -1
	for i, contact := range this.list {
		if contact.Id == id {
			index = i
			break
		}
	}
	this.list = append(this.list[:index], this.list[index+1:]...)
}

func (this *ContactList) DeleteByFullName(firstname, lastname string) {
	index := -1
	for i, contact := range this.list {
		if contact.FirstName == firstname && contact.LastName == lastname {
			index = i
			break
		}
	}
	this.list = append(this.list[:index], this.list[index+1:]...)
}

func (this *ContactList) DeleteByPhone(phonenum uint64) {
	index := -1
	for i, contact := range this.list {
		if contact.Phone == phonenum {
			index = i
			break
		}
	}
	this.list = append(this.list[:index], this.list[index+1:]...)
}

func (this *ContactList) DeleteByEmail(email string) {
	index := -1
	for i, contact := range this.list {
		if contact.Email == email {
			index = i
			break
		}
	}
	this.list = append(this.list[:index], this.list[index+1:]...)
}

func (this *ContactList) DeleteAllByFullName(firstname, lastname string) {

}

type Task struct {
	Id        int
	Name      string
	Status    string
	Priority  int
	CreatedAt string
	CreatedBy Contact
	DueDate   string
}

type TaskList struct {
	list []Task
}

func (this *TaskList) CreateTask(task Task) {
	this.list = append(this.list, task)
}

func main() {

}
