# Basic Full Stack Application

## Frontend

## Backend

Golang is the coding language used for the Server Side of the application

Handler Functions needed for backend to communicate with frontend

```
<div class="header">
      <div class="taskbar">
        <a href="index.html">Home</a>
        <a href="insert.html">Insert</a>
        <a href="read.html">Read</a>
        <a href="update.html">Update</a>
        <a href="delete.html">Delete</a>
      </div>
</div>
```

Test code to query an active db (testschdb)

```
rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var klass string
		var grade string

		if err := rows.Scan(&name, &klass, &grade); err != nil {
			panic(err)
		}
		fmt.Println("Name: ", name, "Class: ", klass, "Grade: ", grade)
	}
```

## Database

Postgresql is the database storage container for all of the data

### Tables

this example uses a simple database with one table and three fields.
