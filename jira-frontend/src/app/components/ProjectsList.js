// src/components/ProjectsList.js
"use client";

import React, { useEffect, useState } from 'react';

function ProjectsList({onProjectClick, activeProjectID}) {
    const [projects, setProjects] = useState([]) // list of projects (intitially empty) updated with setProjects function

    useEffect(() => {
        fetch('http://localhost:8080/projects') // fetch list of projects
            .then((response) => response.json()) // pass response body as json
            .then((data) => setProjects(data)) // call setProjects using json
            .catch((error) => console.error('Error fetching projects:', error)); // handle errors
    }, []) // runs only when component mounts

    const handleDelete = (projectID) => {
        fetch(`http://localhost:8080/projects/${projectID}`, {
            method: 'DELETE'
        })
        .then((response) => response.json())
        .then(() => {
            // Refresh the page to reload the updated projects list
            window.location.reload();
        })
        .catch((error) => console.error('Error deleting project:', error));
    };


    return (
        <ul>
            {projects.map((project) => (
                <li 
                    key={project.ID} 
                    onClick={() => onProjectClick(project.ID)}
                    className={`flex justify-between items-center p-2 cursor-pointer rounded-lg 
                        ${project.ID === activeProjectID ? 'bg-blue-500 text-white' : 'hover:bg-gray-700'}`}
                >
                    <span>{project.Name}</span>
                    <button 
                        onClick={(e) => {
                            e.stopPropagation(); // Prevent triggering the project selection
                            handleDelete(project.ID);
                        }}
                        className="ml-2 text-red-500 hover:text-red-700"
                    >
                        🗑️
                    </button>
                </li>
            ))}
        </ul>
    );
}

export default ProjectsList