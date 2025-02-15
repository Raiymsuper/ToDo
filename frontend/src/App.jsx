import { useState, useEffect } from "react";
import { ListTodos, Insert, ToggleStatus, DeleteTask } from "../wailsjs/go/main/App"; // Wails backend functions

function App() {
    const [tasks, setTasks] = useState([]);
    const [newTask, setNewTask] = useState("");

    // Load tasks from backend on mount
    useEffect(() => {
        loadTasks();
    }, []);

    const loadTasks = async () => {
        try {
            const data = await ListTodos();
            setTasks(data);
        } catch (error) {
            console.error("Error loading tasks:", error);
        }
    };

    const addTask = async () => {
        if (!newTask.trim()) return; // Prevent empty tasks
        try {
            await Insert({ task: newTask, status: false });
            setNewTask(""); // Clear input
            loadTasks(); // Reload tasks
        } catch (error) {
            console.error("Error adding task:", error);
        }
    };

    const toggleStatus = async (id) => {
        try {
            await window.go.main.App.ToggleStatus(id);
            setTasks(tasks.map(task =>
                task.Id === id ? { ...task, Status: !task.Status } : task
            ));
        } catch (error) {
            console.error("Error toggling status:", error);
        }
    };
    

    const deleteTask = async (id) => {
        try {
            await window.go.main.App.DeleteTask(id);
            setTasks(tasks.filter(task => task.Id !== id)); // Remove from state
        } catch (error) {
            console.error("Error deleting task:", error);
        }
    };

    return (
        <div className="container">
            <h1>To-Do List</h1>

            {/* Task Input Field */}
            <div className="task-input">
                <input 
                    type="text" 
                    value={newTask} 
                    onChange={(e) => setNewTask(e.target.value)}
                    placeholder="Enter a new task..."
                />
                <button onClick={addTask}>Add Task</button>
            </div>

            {/* Task List */}
            {/* <ul>
                {tasks.map((task) => (
                    <li key={task.Id} style={{ display: "flex", alignItems: "center", justifyContent: "space-between", color: "white" }}>
                        <span 
                            onClick={() => toggleStatus(task.Id)} 
                            style={{ textDecoration: task.Status ? "line-through" : "none", cursor: "pointer" }}
                        >
                            {task.Task} {task.Status ? "✅" : "❌"}
                        </span>
                        <button 
                            onClick={() => deleteTask(task.Id)} 
                            style={{ background: "red", color: "white", border: "none", cursor: "pointer", padding: "5px 10px", marginLeft: "10px" }}
                        >
                            Delete
                        </button>
                    </li>
                ))}
            </ul> */}
        </div>
    );
}

export default App;
