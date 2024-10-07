// src/app/page.tsx

import TaskList from "./components/TaskList";

export default function Home() {
  return (
    <main>
      <h1>Welcome to the Jira Clone</h1>
      <TaskList />
    </main>
  );
}
