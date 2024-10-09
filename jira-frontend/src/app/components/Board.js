// Board.js
import React, { useState, useEffect } from 'react';
import TaskModal from "./TaskModal"

function Board({ activeProjectID }) {
    const [tasks, setTasks] = useState([]); // List of tasks for the active project
    const [isTaskFormVisible, setIsTaskFormVisible] = useState(false);
    const [taskTitle, setTaskTitle] = useState("");
    const [selectedTask, setSelectedTask] = useState(null);


    useEffect(() => {
        if (activeProjectID) {
            fetch(`http://localhost:8080/tasks?projectID=${activeProjectID}`)
                .then((response) => response.json())
                .then((data) => setTasks(data))
                .catch((error) => console.error('Error fetching tasks:', error));
        }
    }, [activeProjectID]);

    const handleNewTaskSubmit = (e) => {
        e.preventDefault();
        fetch('http://localhost:8080/tasks', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: new URLSearchParams({ 
                title: taskTitle, 
                status: "To Do", 
                projectID: activeProjectID 
            })
        })
        .then((response) => response.json())
        .then((data) => {
            setTasks((prevTasks) => [...prevTasks, data]); // add new task to list
            setTaskTitle(""); // clear title input
            setIsTaskFormVisible(false); // hide form
        })
        .catch((error) => console.error('Error creating task:', error));
    };

    const tasksByStatus = {
        "To Do": tasks.filter((task) => task.Status === "To Do"),
        "In Progress": tasks.filter((task) => task.Status === "In Progress"),
        "Done": tasks.filter((task) => task.Status === "Done"),
    };

    return (
        <div className="flex space-x-4 p-4 bg-[rgb(43,47,54)]">
            {Object.keys(tasksByStatus).map((status) => (
                <div key={status} className="flex-1 bg-gray-700
                 p-4 rounded-lg shadow-md">
                    <h3 className="text-lg font-semibold mb-4">{status}</h3>
                    <ul className="space-y-2">
                        {tasksByStatus[status].map((task) => (
                            <li 
                                key={task.ID} 
                                className="bg-gray-500 p-3 rounded-md shadow-sm border border-gray-200"
                                onClick={() => setSelectedTask(task)}
                            >
                                <h4 className="font-bold">{task.Title}</h4>
                                <p className="text-gray-300">{task.Description}</p>
                            </li>
                        ))}
                    </ul>
                    {status === "To Do" && (
                        <>
                            <button 
                                onClick={() => setIsTaskFormVisible(true)} 
                                className="mt-2 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
                            >
                                + New Task
                            </button>
                            {isTaskFormVisible && (
                                <form onSubmit={handleNewTaskSubmit} className="mt-4 space-y-2">
                                    <input 
                                        type="text" 
                                        value={taskTitle}
                                        onChange={(e) => setTaskTitle(e.target.value)}
                                        placeholder="Task Title"
                                        className="w-full p-2 border border-gray-300 rounded"
                                        required
                                    />
                                    <button 
                                        type="submit"
                                        className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600"
                                    >
                                        Create Task
                                    </button>
                                </form>
                            )}
                        </>
                    )}
                </div>
            ))}
            {selectedTask && (
                <TaskModal 
                    task={selectedTask} 
                    onClose={() => setSelectedTask(null)} // Close modal
                />
            )}
        </div>
    );
}

export default Board;
