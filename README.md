# To Do Application

## Introduction
The assignments in the main content page and Part 1 of this application are aimed to provide sufficient knowledge to enable you to build a To Do Application for that can be accessed by multiple users. Part 1 will have you build multiple programs within the context of the To Do Application. Part 2 has incremental steps to progressively lead you to build an application that should be able to deal with multiple users performing Create, Read, Update and Deletions actions to their To Do lists at the same time.

The primary focus is on the back-end aspects of the application, the front-end is out of scope though good to add. 

## Project
**By the end of this project, you should be able to create a REST API that does the following:**
- Read To Do task
- Create To Do task
- Read / list all To Do tasks
- Delete To Do task 
- Complete To Do task

(Data is in memory only and will be lost on switch off)

**The application should:**
- Be tested (unit tests, integration tests etc.)
- Follow Go best practices

**Stretch Goal:**
- Add database - structural piece

## Part 1
10. Create a program using a variadic function to print a list of 10 things To Do. [Variadic Functions][Structures]
11. Create a program to output a list of 10 things To Do in JSON format. [Variadic Functions][Structures][JSON]
12. Create a program using a variadic function to output a list of 10 things To Do to a JSON format file. [Variadic Functions][Structures][JSON]
13. Create a console program to read a list of 10 things To Do from a JSON format file and display. [Variadic Functions][Structures][JSON]
14. Write a program to simulate a race condition occurring when one goroutine updates a data variable with odd numbers, while another updates the same data variable with even numbers. After each update , attempt to display the data contained in the data variable to screen. [Goroutines][Concurrency][Race Conditions]
15. Refactor the program created in exercise 14 to use channels, mutexes to synchronise all actions. [Concurrency][Waitgroups][Workerpools][Mutexes]
16. Create a program that prints a list of things To Do and the current status of the To Do item using two goroutines which alternate between To Do Items and To Do statuses [Concurrency][Waitgroups][Workerpools][Mutexes]

## Part 2
Part 2 extends the topics covered in Part 1 to complete the build of a To Do list application.

17. Create a command line app to manage a To Do list stored in memory. This should enable a user to perform Create, Read, Update, and Delete operations on a list of To Do items. The list should contain a To Do item and a To Do status. [Structures][Arrays][Variadic Functions]
18. Convert the command line app into web page app to manage To Do list stored in memory. [Structures][Arrays][Variadic functions][Keyboard Input][Interfaces]
19. Remote Commands - Create a server that can concurrently receive a list of pre-defined commands, The status of the server and the status of each task should be available via specific commands. [Concurrency][Goroutines] [Channels]
20. Extend to a web API to receive web page actions [Remote Commands] that are applied to To Do list stored in memory.[ File Server][Web API][Interfaces][Http]

## Stretch Goals
The stretch goals integrate a traditional database and focus on concurrent user interactions.

21. Extend web API to receive web page actions [Remote Commands] that are applied to To Do list which is stored to a database table. [Concurrency] [Goroutines] [Channels] [Mutexes][Interfaces][Http][Databases]
22. Extend the Web API to receives actions [Remote Commands] to be applied to the To Do application list from multiple users. All actions to be applied to the database tables for each user. [Concurrency][Goroutines][Channels] [Mutexes][Interfaces][Databases]
