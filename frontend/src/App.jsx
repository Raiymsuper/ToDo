import { useState, useEffect } from "react";
import { AddTask, GetTasks } from "../wailsjs/go/main/App";

function App() {
  const [task, setTask] = useState("");
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    fetchTasks();
  }, []);

  const fetchTasks = async () => {
    try {
      const response = await GetTasks();
      setTasks(response);
    } catch (error) {
      console.error("Ошибка загрузки задач:", error);
    }
  };

  const handleAddTask = async () => {
    if (!task.trim()) return;
    try {
      await AddTask(task);
      setTask("");
      fetchTasks();
    } catch (error) {
      console.error("Ошибка добавления задачи:", error);
    }
  };

  return (
    <div className="app">
      <h1>ToDo List</h1>
      <div>
        <input
          type="text"
          placeholder="Введите задачу..."
          value={task}
          onChange={(e) => setTask(e.target.value)}
        />
        <button onClick={handleAddTask}>Добавить</button>
      </div>
      <ul>
        {tasks.map((t) => (
          <li key={t.id}>{t.title}</li>
        ))}
      </ul>
    </div>
  );
}

export default App;