# Scheduler

#### Populate with sample data
``` go
	db.Create(&m.Task{
		Title:  "Task 1",
		Status: "TODO",
	})

	db.Create(&m.Task{
		Title:  "Task 2",
		Status: "In Progress",
	})
```
