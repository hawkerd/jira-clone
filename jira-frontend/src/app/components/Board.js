// Board.js
import React, { useState, useEffect } from 'react';

function Board({ activeProjectID }) {
    const [tasks, setTasks] = useState([]); // List of tasks for the active project

    useEffect(() => {
        if (activeProjectID) {
            fetch(`http://localhost:8080/tasks?projectID=${activeProjectID}`)
                .then((response) => response.json())
                .then((data) => setTasks(data))
                .catch((error) => console.error('Error fetching tasks:', error));
        }
    }, [activeProjectID]);

    const tasksByStatus = {
        "To Do": tasks.filter((task) => task.Status === "To Do"),
        "In Progress": tasks.filter((task) => task.Status === "In Progress"),
        "Done": tasks.filter((task) => task.Status === "Done"),
    };

    return (
        <div className="flex space-x-4 p-4 bg-gray-500">
            {Object.keys(tasksByStatus).map((status) => (
                <div key={status} className="flex-1 bg-gray-700
                 p-4 rounded-lg shadow-md">
                    <h3 className="text-lg font-semibold mb-4">{status}</h3>
                    <ul className="space-y-2">
                        {tasksByStatus[status].map((task) => (
                            <li 
                                key={task.ID} 
                                className="bg-gray-500 p-3 rounded-md shadow-sm border border-gray-200"
                            >
                                <h4 className="font-bold">{task.Title}</h4>
                                <p className="text-gray-300">{task.Description}</p>
                            </li>
                        ))}
                    </ul>
                </div>
            ))}
        </div>
    );
}

export default Board;
