 

# Task Management

  

## Project Description

Fullstack Task Management is a system designed to help users efficiently manage tasks. 
It includes features to add, delete, and edit tasks, making it easier to organize and track daily activities.

  

---

  

## Technologies Used

### Front-end

- Framework: Vue 3

- UI Framework: Tailwind CSS

  

### Back-end

- Go language

- Framework: Fiber

  

### Database

- PostgreSQL

- GORM (ORM)

  

### Authentication

- JSON Web Tokens (JWT)

  

### Containerization

- Docker

  

---

  

## Installation

1. Clone the project from GitHub:

```bash

git clone https://github.com/ilham2000r/vue-go-taskmanager.git

```

2. Ensure [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) are installed.

```bash
docker --version 
docker-compose --version

```

3. Build and run the Docker containers:

```bash

docker-compose up --build

```

4. Open your browser and access the application at `http://localhost:5173`.

  

---

  

## Usage

  

### 1. Log in or Register

- Users can log in using their **username** and **password**.

- If you don’t have an account, select **Register** to create a new one.

  

### 2. Add New Task

- Provide task details such as name, description, due date, and priority.

- Click **Submit Button** to save the task.

  

### 3. Edit or Delete Tasks

- Navigate to **your task **.

- Edit the details of the task as needed.

- To delete a task, click the **trash icon** 🗑️ after viewing the task details.

  

### 4. Change Task Status

- Click on a task and change its status using the status button to: **"In Progress"**, **"Completed"**.

  

### 5. Search and Filter Tasks

- Use the **search bar** to find tasks by name.

- Use the **filter options** in the top-right corner to view tasks based on their priority.

  

> **Note:**

> - Notifications will appear after successful actions, such as saving or deleting tasks.

> - To log out, click the **"Logout"** button in the menu.

  

---

  

## Project Structure

  

```plaintext

VUE-GO-TASKMANAGEMENT/

├── client/
│ ├── public/
│ ├── src/
│ │ ├── api/
│ │ ├── components/
│ │ ├── layout/
│ │ ├── router/
│ │ ├── stores/
│ │ ├── views/
│ │ ├── App.vue
│ │ ├── main.js
│ └── Dockerfile
│
├── server/
│ ├── config/
│ ├── controllers/
│ ├── middleware/
│ ├── models/
│ ├── routes/
│ ├── main.go
│ └── Dockerfile
│
├── docker-compose.yml
├── README.md

```

  

---

  

## Screenshots

![Login page](https://img2.pic.in.th/pic/3b7f0452-06ca-4aac-8484-1ce01cc84abd.jpg)
![Index page - form input](https://img5.pic.in.th/file/secure-sv1/0ccc23b4-77b0-4c72-ae54-b82402cc917b.jpg)


![index page - task list](https://img5.pic.in.th/file/secure-sv1/06989ac9-8dde-4f49-9fb5-e7931149d03c.jpg)

![task detail page](https://img2.pic.in.th/pic/a3d18199-7b15-48bf-a572-aa41c91015c7.jpg)


## License
MIT License

## Contributing
Contributions, issues, and feature requests are welcome!


