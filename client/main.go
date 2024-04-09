package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "stu-info-mgr/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStudentServiceClient(conn)

	ctx := context.Background()
	// Prompt the user for input
	fmt.Println("Welcome to the Student Info Manager CLI!")
	for {
		fmt.Println("Please select an option:")
		fmt.Println("1. Add a student")
		fmt.Println("2. Query student by ID")
		fmt.Println("3. Query student by name")
		fmt.Println("4. Delete a student")
		fmt.Println("5. Exit")

		// Read user input
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			// Add a student
			var student pb.Student
			fmt.Println("Enter student ID:")
			fmt.Scanln(&student.Id)
			fmt.Println("Enter student name:")
			fmt.Scanln(&student.Name)

			// Call the AddStudent RPC
			_, err := c.Add(ctx, &pb.AddRequest{Student: &student})
			if err != nil {
				log.Fatalf("failed to add student: %v", err)
			}
			fmt.Println("Student added successfully!")

		case 2:
			// Query student by ID
			var studentID int32
			fmt.Println("Enter student ID:")
			fmt.Scanln(&studentID)

			// Call the GetStudentByID RPC
			student, err := c.Query(ctx, &pb.QueryRequest{Id: studentID})
			if err != nil {
				log.Fatalf("failed to get student: %v", err)
			}
			fmt.Printf("Student found:\nID: %d\nName: %s\n", student.Id, student.Name)

		case 3:
			// Query student by name
			var studentName string
			fmt.Println("Enter student name:")
			fmt.Scanln(&studentName)

			// Call the GetStudentByName RPC
			student, err := c.QueryByName(ctx, &pb.QueryByNameRequest{Name: studentName})
			if err != nil {
				log.Fatalf("failed to get students: %v", err)
			}
			fmt.Printf("Students found:\n")
			fmt.Printf("ID: %d\nName: %s\n\n", student.Id, student.Name)

		case 4:
			var studentID int32
			fmt.Println("Enter student ID:")
			fmt.Scanln(&studentID)

			_, err := c.Delete(ctx, &pb.DeleteRequest{Id: studentID})
			if err != nil {
				log.Fatalf("failed to delete student: %v", err)
			}
			fmt.Println("Student deleted successfully!")

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
