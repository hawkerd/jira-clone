// Board.js
import React, { useState, useEffect } from 'react';
import TaskModal from "./TaskModal"

const apiUrl = process.env.NEXT_PUBLIC_API_URL; // For local testing, you might have it set in your .env file


function Board({ activeProjectID }) {
    const [tasks, setTasks] = useState([]); // List of tasks for the active project
    const [isTaskFormVisible, setIsTaskFormVisible] = useState(false);
    const [taskTitle, setTaskTitle] = useState("");
    const [selectedTask, setSelectedTask] = useState(null);


    useEffect(() => {
        if (!activeProjectID) return;
        
        const fetchTasks = () => {
            fetch(`${apiUrl}/tasks?projectID=${activeProjectID}`)
                .then((response) => response.json())
                .then((data) => setTasks(data))
                .catch((error) => console.error('Error fetching tasks:', error));
        };
    
        fetchTasks(); // Fetch tasks immediately
    
        const intervalId = setInterval(fetchTasks, 5000); // Fetch every 5 seconds
        
        return () => clearInterval(intervalId); // Clear interval on component unmount or project change
    }, [activeProjectID]);
    

    const handleNewTaskSubmit = (e) => {
        e.preventDefault();
        fetch(`${apiUrl}/tasks`, {
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
        .then(() => {
            setTaskTitle(""); // Clear title input
            setIsTaskFormVisible(false); // Hide form
            // Fetch the updated task list to ensure it's accurate
            fetch(`${apiUrl}/tasks?projectID=${activeProjectID}`)
                .then((response) => response.json())
                .then((updatedTasks) => setTasks(updatedTasks))
                .catch((error) => console.error('Error fetching updated tasks:', error));
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
