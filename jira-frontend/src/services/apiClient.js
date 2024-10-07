// src/services/apiClient.js

import axios from "axios";

const apiClient = axios.create({
  baseURL: "http://localhost:8080",
  headers: {
    "Content-Type": "application/json",
  },
});

export const fetchTasks = async () => {
  try {
    const response = await apiClient.get("/tasks");
    return response.data;
  } catch (error) {
    console.error("Error fetching tasks:", error);
    return [];
  }
};

export default apiClient;
