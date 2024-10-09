// src/models/Task.js

export default class Task {
    constructor(id, title, description, status, projectID, createdAt, updatedAt) {
      this.id = id;
      this.title = title;
      this.description = description;
      this.status = status;
      this.projectID = projectID;
      this.createdAt = createdAt;
      this.updatedAt = updatedAt;
    }
  }
  