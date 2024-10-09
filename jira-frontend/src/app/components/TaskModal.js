// src/components/TaskModal.js
import React from 'react';

function TaskModal({ task, onClose }) {
    if (!task) return null;

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
                
                {/* Task details */}
                <h2 className="text-3xl font-bold mb-4">{task.Title}</h2>
                <p className="text-lg mb-6">{task.Description}</p>
                <p className="text-sm mb-2"><strong>Status:</strong> {task.Status}</p>

                {/* Optional: Add more task details here */}
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