// src/app/page.tsx
"use client";

import { useState } from "react";
import Sidebar from "./components/Sidebar";
import Board from "./components/Board";

export default function Home() {
  const [activeProjectID, setActiveProjectID] = useState(null); // Manage active project ID

  return (
    <main className="flex h-screen bg-[rgb(43,47,54)]">
      <Sidebar onProjectClick={setActiveProjectID} activeProjectID={activeProjectID} />
      <div className="flex-1 p-4">
        {activeProjectID ? (
          <Board activeProjectID={activeProjectID} />
        ) : (
          <p className="text-gray-600 bg-[rgb(43,47,54)]">Select a project to view its tasks</p>
        )}
      </div>
    </main>
  );
}


