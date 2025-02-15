import { useState, useEffect } from "react";
import { ListTodos, Insert, ToggleStatus, DeleteTask } from "../wailsjs/go/main/App";

function App() {
    const [tasks, setTasks] = useState([]);
    const [newTask, setNewTask] = useState("");
    const [deleteId, setDeleteId] = useState(null); 
    const [showModal, setShowModal] = useState(false);

    useEffect(() => {
        loadTasks();
    }, []);

    const loadTasks = async () => {
        try {
            const data = await ListTodos();
            setTasks(data || []);
        } catch (error) {
            console.error("Ошибка при загрузке задач:", error);
            setTasks([]);
        }
    };

    const addTask = async () => {
        if (!newTask.trim()) return;
        try {
            await Insert({ task: newTask, status: false });
            setNewTask("");
            loadTasks();
        } catch (error) {
            console.error("Ошибка при добавлении задачи:", error);
        }
    };

    const toggleStatus = async (id) => {
        try {
            await ToggleStatus(id);
            setTasks(tasks.map(task =>
                task.Id === id ? { ...task, Status: !task.Status } : task
            ));
        } catch (error) {
            console.error("Ошибка при изменении статуса:", error);
        }
    };

    const confirmDelete = (id) => {
        setDeleteId(id);
        setShowModal(true);
    };

    const deleteTask = async () => {
        if (!deleteId) return;
        try {
            await DeleteTask(deleteId);
            setTasks(tasks.filter(task => task.Id !== deleteId));
        } catch (error) {
            console.error("Ошибка при удалении задачи:", error);
        }
        setShowModal(false);
        setDeleteId(null);
    };

    const pendingTasks = tasks.filter(task => !task.Status);
    const completedTasks = tasks.filter(task => task.Status);

    return (
        <div className="container" style={{ color: "white", textAlign: "center", padding: "20px" }}>
            <h1>Список дел</h1>

            {/* Поле ввода задачи */}
            <div className="task-input">
                <input 
                    type="text" 
                    value={newTask} 
                    onChange={(e) => setNewTask(e.target.value)}
                    placeholder="Введите новую задачу..."
                />
                <button onClick={addTask}>Добавить</button>
            </div>

            <div style={{
                display: "flex",
                justifyContent: "space-around",
                alignItems: "start",
                marginTop: "20px"
            }}>
                {/* Невыполненные задачи */}
                <div style={{ flex: 1, padding: "20px" }}>
                    <h2>Невыполненные</h2>
                    <ul style={{ listStyle: "none", padding: 0 }}>
                        {pendingTasks.map((task) => (
                            <li key={task.Id} style={{ display: "flex", justifyContent: "space-between", alignItems: "center", padding: "10px 0" }}>
                                <span 
                                    onClick={() => toggleStatus(task.Id)} 
                                    style={{ cursor: "pointer" }}
                                >
                                    {task.Task} ❌
                                </span>
                                <button 
                                    onClick={() => confirmDelete(task.Id)} 
                                    style={{ background: "red", color: "white", border: "none", cursor: "pointer", padding: "5px 10px", marginLeft: "10px" }}
                                >
                                    Удалить
                                </button>
                            </li>
                        ))}
                    </ul>
                </div>

                {/* Выполненные задачи */}
                <div style={{ flex: 1, padding: "20px" }}>
                    <h2>Выполненные</h2>
                    <ul style={{ listStyle: "none", padding: 0 }}>
                        {completedTasks.map((task) => (
                            <li key={task.Id} style={{ display: "flex", justifyContent: "space-between", alignItems: "center", padding: "10px 0" }}>
                                <span 
                                    onClick={() => toggleStatus(task.Id)} 
                                    style={{ textDecoration: "line-through", cursor: "pointer" }}
                                >
                                    {task.Task} ✅
                                </span>
                                <button 
                                    onClick={() => confirmDelete(task.Id)} 
                                    style={{ background: "red", color: "white", border: "none", cursor: "pointer", padding: "5px 10px", marginLeft: "10px" }}
                                >
                                    Удалить
                                </button>
                            </li>
                        ))}
                    </ul>
                </div>
            </div>

            {/* Модальное окно подтверждения удаления */}
            {showModal && (
                <div style={{
                    position: "fixed", top: 0, left: 0, width: "100%", height: "100%",
                    background: "rgba(0,0,0,0.5)", display: "flex", alignItems: "center",
                    justifyContent: "center"
                }}>
                    <div style={{
                        background: "white", padding: "20px", borderRadius: "8px", textAlign: "center", color: "black"
                    }}>
                        <h3>Вы уверены, что хотите удалить задачу?</h3>
                        <button onClick={deleteTask} style={{ marginRight: "10px", padding: "5px 10px", background: "red", color: "white", border: "none", cursor: "pointer" }}>Удалить</button>
                        <button onClick={() => setShowModal(false)} style={{ padding: "5px 10px", background: "gray", color: "white", border: "none", cursor: "pointer" }}>Отмена</button>
                    </div>
                </div>
            )}
        </div>
    );
}

export default App;