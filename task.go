package main

type Item struct {
	ID        int
	Title     string
	Completed bool
}

type Tasks struct {
	Items          []Item
	Count          int
	CompletedCount int
}

func fetchTasks() ([]Item, error) {
	var items []Item

	rows, err := DB.Query("SELECT id, title, completed FROM tasks ORDER BY position;")
	if err != nil {
		return []Item{}, err
	}

	defer rows.Close()

	for rows.Next() {
		item := Item{}
		err := rows.Scan(&item.ID, &item.Title, &item.Completed)
		if err != nil {
			return []Item{}, err
		}
		items = append(items, item)
	}
	return items, nil
}

func fetchTask(id int) (Item, error) {
	var item Item
	err := DB.QueryRow("SELECT id, title, completed FROM tasks WHERE id = (?)", id).Scan(&item.ID, &item.Title, &item.Completed)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func updateTask(id int, title string) (Item, error) {
	var item Item
	err := DB.QueryRow("UPDATE tasks SET title = (?) WHERE id = (?) RETURNING id, title, completed", title, id).Scan(&item.ID, &item.Title, &item.Completed)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}
