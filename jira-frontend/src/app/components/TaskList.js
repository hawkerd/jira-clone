// src/app/components/TaskList.js

"use client";

import React, { useEffect, useState } from "react";
import Task from "../../models/Task";
import { fetchTasks } from "../../services/apiClient";
import styles from "./TaskList.module.css";

const TaskList = () => {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    const getTasks = async () => {
      const data = await fetchTasks();
      const taskObjects = data.map(
        taskData =>
          new Task(
            taskData.ID,
            taskData.Title,
            taskData.Description,
            taskData.Status,
            taskData.ProjectID,
            taskData.CreatedAt,
            taskData.UpdatedAt
          )
      );
      setTasks(taskObjects);
    };

    getTasks();
  }, []);

  return (
    <div className={styles.taskListContainer}>
      <h2 className={styles.taskListTitle}>Tasks</h2>
      <ul>
        {tasks.length === 0 ? (
          <p className={styles.noTasksMessage}>No tasks found.</p>
        ) : (
          tasks.map((task, index) => (
            <li key={task.id || index} className={styles.taskItem}>
              <h4>{task.title || "No Title"}</h4>
              <p>{task.description || "No Description"}</p>
              <p className={styles.taskItemStatus}>Status: {task.status || "No Status"}</p>
              <p>Project ID: {task.projectID || "No Project ID"}</p>
            </li>
          ))
        )}
      </ul>
    </div>
  );
};

export default TaskList;
