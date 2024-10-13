import React, { useState } from 'react';

function TaskModal({ task, onClose }) {
    const [description, setDescription] = useState(task.Description);
    const [isEditing, setIsEditing] = useState(false);

    const handleSave = () => {
        fetch(`${process.env.NEXT_PUBLIC_API_URL}/tasks/${task.ID}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: task.Title,           // Assuming task.Title exists
                description: description,
                status: task.Status,         // Assuming task.Status exists
                projectID: task.ProjectID    // Assuming task.ProjectID exists
            }),
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(() => {
            setIsEditing(false); // Exit edit mode
        })
        .catch((error) => {
            console.error("Error updating task:", error);
        });
    };
    
    

    return (
        <div className="fixed inset-0 bg-black bg-opacity-60 flex items-center justify-center z-50">
            <div className="bg-gray-800 text-white w-3/4 h-3/4 p-6 rounded-lg shadow-lg relative overflow-y-auto">
                {/* Close button */}
                <button 
                    onClick={onClose} 
                    className="absolute top-4 right-4 text-gray-300 hover:text-white text-xl font-bold"
                >
                    &times;
                </button>
                
                {/* Task Title */}
                <h2 className="text-3xl font-bold mb-4">{task.Title}</h2>

                {/* Editable Description */}
                <div className="mb-6">
                    <label className="text-lg font-semibold">Description:</label>
                    {isEditing ? (
                        <div>
                            <textarea
                                value={description}
                                onChange={(e) => setDescription(e.target.value)}
                                onBlur={handleSave}  // Save on blur or can be removed if Save button is used exclusively
                                autoFocus
                                className="w-full mt-2 p-2 rounded bg-gray-700 text-white border border-gray-600"
                                placeholder="Enter task description"
                            />
                            
                            {/* Save and Cancel Buttons */}
                            <div className="mt-2 flex space-x-2">
                                <button
                                    onClick={handleSave}
                                    className="px-4 py-1 bg-green-500 text-white rounded hover:bg-green-600"
                                >
                                    Save
                                </button>
                                <button
                                    onClick={() => {
                                        setDescription(task.Description); // Reset to original description on cancel
                                        setIsEditing(false);  // Exit edit mode
                                    }}
                                    className="px-4 py-1 bg-red-500 text-white rounded hover:bg-red-600"
                                >
                                    Cancel
                                </button>
                            </div>
                        </div>
                    ) : (
                        <p 
                            className="text-lg mt-2 cursor-pointer" 
                            onClick={() => setIsEditing(true)}
                        >
                            {description || "No description provided."}
                        </p>
                    )}
                </div>




                {/* Status */}
                <p className="text-sm mb-2"><strong>Status:</strong> {task.Status}</p>

                {/* Optional: More Task Details */}
                <div className="mt-8">
                    <h3 className="text-xl font-semibold mb-2">Additional Details</h3>
                    <p>Created at: {new Date(task.CreatedAt).toLocaleDateString()}</p>
                    <p>Updated at: {new Date(task.UpdatedAt).toLocaleDateString()}</p>
                </div>
            </div>
        </div>
    );
}

export default TaskModal;
