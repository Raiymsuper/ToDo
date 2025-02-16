# To-Do List App (Wails + React + SQLite)

## 📌 Описание
Простое To-Do приложение, созданное с использованием **Wails**, **React**, **Go** и **SQLite**. Позволяет:
- Отмечать задачи выполненными.
- Удалять задачи.
- Сохранять данные между перезапусками.

## 🛠 Технологии
- **Frontend**: React (Vite)
- **Backend**: Go (Wails)
- **База данных**: SQLite

## 🚀 Установка и запуск

### 1️⃣ Клонирование репозитория
```bash
  git clone https://github.com/Raiymsuper/ToDo.git
  cd ToDo
```

### 2️⃣ Сборка готового приложения
```bash
  wails build
```
После сборки **готовый исполняемый файл** будет находиться в `./build/bin/`.
**Однако он не запускается так**

### Запуск только через
```bash
  wails dev
```
## 🎯 Функционал
### Добавление задачи:
1. Введите текст задачи.
2. Нажмите **Add Task**.

### Отметка задачи выполненной или невыполненой:
- Кликните на задачу для изменения ее состояния.
![image](https://github.com/user-attachments/assets/70dac05d-4a2f-4995-82bf-4a62cc895821)
![image](https://github.com/user-attachments/assets/9f8f812b-bce8-4af2-b0cb-3cf7aea3f319)


### Удалить задачу:
- Нажмите красную кнопку "удалить" рядом с задачей.
![image](https://github.com/user-attachments/assets/5da12fcf-ec5d-4c88-a5ce-2c419d51744b)
