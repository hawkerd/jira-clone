// src/components/Sidebar.js
"use client";

import React from 'react';
import ProjectsList from './ProjectsList';

function Sidebar({ onProjectClick, activeProjectID }) { // Receive props from Home
    return (
        <div className="w-64 h-full p-4 bg-gray-700 border-r border-gray-700 shadow-lg">
            <h2 className="text-xl font-semibold mb-6 text-white-400">Projects</h2>
            <ProjectsList 
                onProjectClick={onProjectClick} 
                activeProjectID={activeProjectID} 
            />
        </div>
    );
}

export default Sidebar;
