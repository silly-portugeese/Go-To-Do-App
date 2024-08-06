# To Do Application

This project features a series of Go applications, with a primary focus on developing a comprehensive [WebApp](#web-app). The development is structured in stages, starting with some exercises, followed by a command-line interface (CLI) app and culminating in a fully functional web application. The run.sh script helps you execute different parts of the project.

## Usage
To use the run.sh script, execute:

`./run.sh [command]`

### Available Commands
Run the exercises

`./run.sh exercises`


Run the concurrency exercises

`./run.sh concurrency_exercises`


Start the backend service for the WebApp.

`./run.sh webapp-backend`


Start the frontend service for the WebApp.

`./run.sh webapp-frontend`


Start both the backend and frontend services for the WebApp:

`./run.sh webapp`




## Exercises
1. Create a program using a variadic function to print a list of 10 things To Do. [Variadic Functions][Structures]
2. Create a program to output a list of 10 things To Do in JSON format. [Variadic Functions][Structures][JSON]
3. Create a program using a variadic function to output a list of 10 things To Do to a JSON format file. [Variadic Functions][Structures][JSON]
4. Create a console program to read a list of 10 things To Do from a JSON format file and display. [Variadic Functions][Structures][JSON]
5. Create a program that prints a list of things To Do and the current status of the To Do item using two goroutines which alternate between To Do Items and To Do statuses.[Concurrency][Waitgroups][Workerpools][Mutexes]
    
    Note: Start with the concurrency exercises to get a better understanding of how to use goroutines for this task.


## Concurrency exercises
1. Write a program to simulate a race condition occurring when one goroutine updates a data variable with odd numbers, while another updates the same data variable with even numbers. After each update , attempt to display the data contained in the data variable to screen. [Goroutines][Concurrency][Race Conditions]
2. Refactor the program created in the previous exercise to use channels, mutexes to synchronise all actions. [Concurrency][Waitgroups][Workerpools][Mutexes]


## Command line app
This extends the topics covered in "Exercises" to complete the build of a To Do list application.
1. Create a command line app to manage a To Do list stored in memory. This should enable a user to perform Create, Read, Update, and Delete operations on a list of To Do items. The list should contain a To Do item and a To Do status. [Structures][Arrays][Variadic Functions]


## Web app
1. Convert the command line app into web page app to manage To Do list stored in memory. [Structures][Arrays][Variadic functions][Keyboard Input][Interfaces]
2. Remote Commands - Create a server that can concurrently receive a list of pre-defined commands, The status of the server and the status of each task should be available via specific commands. [Concurrency][Goroutines] [Channels] - Not implemented
3. Extend to a web API to receive web page actions [Remote Commands] that are applied to To Do list stored in memory.[File Server][Web API][Interfaces][Http]


**The webapp should incorporate a REST API that does the following:**
- Read To Do task
- Create To Do task
- Read / list all To Do tasks
- Delete To Do task 
- Complete To Do task

(Data is in memory only and will be lost on switch off)

## Stretch Goals - Not implemented
The stretch goals integrate a traditional database and focus on concurrent user interactions.

1. Extend web API to receive web page actions [Remote Commands] that are applied to To Do list which is stored to a database table. [Concurrency] [Goroutines] [Channels] [Mutexes][Interfaces][Http][Databases]
2. Extend the Web API to receives actions [Remote Commands] to be applied to the To Do application list from multiple users. All actions to be applied to the database tables for each user. [Concurrency][Goroutines][Channels] [Mutexes][Interfaces][Databases]
