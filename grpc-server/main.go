package main

import (
	"context"
	"database/sql"
	"golang-grpc-app/app"
	"log"
	"net"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	app.UnimplementedStudentsServer
	db *sql.DB
}

// GetStudents implements app.Server
func (s *server) GetStudents(ctx context.Context, in *app.StudentRequest) (*app.StudentResp, error) {
	log.Printf("Received request from ID: %v", in.GetId())
	resp, err := getStudentsFromDB(s.db, in.GetId())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func main() {

	db, err := createAndConnectToSQLiteDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	insertDataIntoDB(db)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	app.RegisterStudentsServer(s, &server{db: db})
	log.Println("Server started successfully...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func createAndConnectToSQLiteDB() (*sql.DB, error) {
	// remove db file if already excist.
	os.Remove("sqlite-database.db")

	log.Println("Creating sqlite-database.db...")
	// Create SQLite file
	file, err := os.Create("sqlite-database.db")
	if err != nil {
		return nil, err
	}
	file.Close()
	log.Println("sqlite-database.db file created")

	// Open the created SQLite File
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return nil, err
	}
	// Create Database Tables
	err = createTable(sqliteDatabase)
	if err != nil {
		return nil, err
	}

	return sqliteDatabase, nil
}

func insertDataIntoDB(sqliteDatabase *sql.DB) {
	// INSERT RECORDS
	insertStudent(sqliteDatabase, "0001", "Liana Kim", "Bachelor")
	insertStudent(sqliteDatabase, "0002", "Glen Rangel", "Bachelor")
	insertStudent(sqliteDatabase, "0003", "Martin Martins", "Master")
	insertStudent(sqliteDatabase, "0004", "Alayna Armitage", "PHD")
	insertStudent(sqliteDatabase, "0005", "Marni Benson", "Bachelor")
	insertStudent(sqliteDatabase, "0006", "Derrick Griffiths", "Master")
	insertStudent(sqliteDatabase, "0007", "Leigh Daly", "Bachelor")
	insertStudent(sqliteDatabase, "0008", "Marni Benson", "PHD")
	insertStudent(sqliteDatabase, "0009", "Klay Correa", "Bachelor")
}

func createTable(db *sql.DB) error {
	// SQL Statement for Create Table
	createStudentTableSQL := `
	CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
	  );`

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		return err
	}

	// Execute SQL Statements
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	log.Println("student table created")
	return nil
}

// We are passing db reference connection from main to our method with other parameters
func insertStudent(db *sql.DB, code string, name string, program string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, program)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getStudentsFromDB(db *sql.DB, id int32) (*app.StudentResp, error) {
	row, err := db.Query("SELECT * FROM student WHERE idStudent = $1;", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	resp := new(app.StudentResp)
	for row.Next() {
		// Iterate and fetch the records from result cursor
		err := row.Scan(&resp.Id, &resp.Code, &resp.Name, &resp.Program)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}
