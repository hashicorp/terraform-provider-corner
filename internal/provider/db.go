package provider

import (
	"fmt"

	"github.com/hashicorp/go-memdb"
)

// Person represents a person in the database.
type Person struct {
	// Email must be unique, and is treated as the Person's UUID.
	Email string
	Name  string
	Age   int
}

// A Client manages communication with the memdb.
type Client struct {
	db *memdb.MemDB
}

// NewClient returns a new memdb client.
func NewClient() (*Client, error) {
	c := &Client{}

	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"person": &memdb.TableSchema{
				Name: "person",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
					"age": &memdb.IndexSchema{
						Name:    "age",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Age"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	c.db = db

	return c, nil
}

func (c *Client) CreatePerson(person *Person) error {
	txn := c.db.Txn(true)
	defer txn.Abort()

	// uniqueness: error if email already exists in db
	existingPerson, err := c.ReadPerson(person.Email)
	if err != nil {
		return fmt.Errorf("Error determining if person with email %s already exists in db: %s", person.Email, err)
	}

	if existingPerson != nil {
		return fmt.Errorf("Cannot create person: person already exists with email %s", person.Email)
	}

	if err := txn.Insert("person", person); err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func (c *Client) ReadPerson(email string) (*Person, error) {
	// Create read-only transaction
	txn := c.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("person", "id", email)
	if err != nil {
		return nil, err
	}

	if raw != nil {
		p := raw.(*Person)
		return p, nil
	}

	return nil, nil
}

func (c *Client) UpdatePerson(person *Person) error {
	txn := c.db.Txn(true)
	defer txn.Abort()

	// email is unique (enforced on write)
	raw, err := txn.First("person", "id", person.Email)
	if err != nil {
		return err
	}

	p := raw.(*Person)

	p.Name = person.Name
	p.Age = person.Age

	err = txn.Insert("person", p)
	if err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func (c *Client) DeletePerson(person *Person) error {
	txn := c.db.Txn(true)
	defer txn.Abort()

	existingPerson, err := c.ReadPerson(person.Email)
	if err != nil {
		return fmt.Errorf("Error determining if person with email %s exists in db: %s", person.Email, err)
	}

	if existingPerson == nil {
		return fmt.Errorf("Cannot delete person with email %s: email not in db", person.Email)
	}

	err = txn.Delete("person", person)
	if err != nil {
		return err
	}

	txn.Commit()

	return nil
}
