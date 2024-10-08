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

    return (
        <ul className="space-y-2">
            {projects.map((project) => (
                <li
                    key={project.ID}
                    onClick={() => onProjectClick(project.ID)} // Call onProjectClick with the project ID
                    className={`p-3 rounded-md cursor-pointer transition-colors duration-150 ${
                        project.ID === activeProjectID 
                            ? 'bg-blue-500 text-white' 
                            : 'bg-white text-gray-800 hover:bg-blue-100'
                    }`}
                >
                    {project.Name}
                </li>
            ))}
        </ul>
    );
}

export default ProjectsList