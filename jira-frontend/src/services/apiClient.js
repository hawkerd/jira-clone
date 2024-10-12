// src/services/apiClient.js

import axios from "axios";

const apiUrl = process.env.NEXT_PUBLIC_API_URL; // For local testing, you might have it set in your .env file

const apiClient = axios.create({
  baseURL: `${apiUrl}`,
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
