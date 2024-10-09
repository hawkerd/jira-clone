// Sidebar.js
"use client";

import React, { useState } from 'react';
import ProjectsList from './ProjectsList';

function Sidebar({ onProjectClick, activeProjectID }) {
    const [isFormVisible, setIsFormVisible] = useState(false); // Track form visibility
    const [projectName, setProjectName] = useState(""); // New project name

    const handleProjectSubmit = (e) => {
        e.preventDefault();
        // POST request to add the new project
        fetch('http://localhost:8080/projects', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: new URLSearchParams({ name: projectName })
        })
        .then((response) => response.json())
        .then((data) => {
            setProjectName(""); // clear input
            setIsFormVisible(false); // hide form
            window.location.reload(); // refresh the page
            setActiveProjectID(data.ID); // set active project to be this one
        })
        .catch((error) => console.error('Error creating project:', error));
    };

    return (
        <div className="w-64 h-full p-4 bg-[rgb(43,47,54)] text-white border-r border-gray-800 shadow-lg">
            <h2 className="text-xl font-semibold mb-6">Projects</h2>
            <ProjectsList onProjectClick={onProjectClick} activeProjectID={activeProjectID} />
            
            {/* New Project Button */}
            <button 
                className="mt-4 w-full bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded"
                onClick={() => setIsFormVisible(!isFormVisible)} // Toggle form visibility
            >
                + New Project
            </button>

            {/* New Project Form */}
            {isFormVisible && (
                <form onSubmit={handleProjectSubmit} className="mt-4">
                    <input
                        type="text"
                        value={projectName}
                        onChange={(e) => setProjectName(e.target.value)}
                        className="w-full p-2 bg-gray-700 text-white rounded mb-2"
                        placeholder="Project Name"
                        required
                    />
                    <button 
                        type="submit"
                        className="w-full bg-green-500 hover:bg-green-600 text-white py-2 px-4 rounded"
                    >
                        Create Project
                    </button>
                </form>
            )}
        </div>
    );
}

export default Sidebar;
