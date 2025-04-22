// src/auth.ts

import { authInstance } from ".";

// Auth - Public
const login = async (data) => {
  const res = await authInstance.post("/auth/login", data);
  return res.data;
};

const logout = async () => {
  const res = await authInstance.post("/auth/logout");
  return res.data;
};

const register = async (data) => {
  const res = await authInstance.post("/auth/register", data);
  return res.data;
};

const forgotPassword = async (data) => {
  const res = await authInstance.post("/auth/forgot-password", data);
  return res.data;
};

const resetPassword = async (data) => {
  const res = await authInstance.post("/auth/reset-password", data);
  return res.data;
};

// Auth - User (Protected)
const me = async () => {
  const res = await authInstance.get("/auth/me");
  return res.data;
};

const changePassword = async (data) => {
  const res = await authInstance.put("/auth/change-password", data);
  return res.data;
};

// Admin - User Management (Protected)
const getAllUsers = async () => {
  const res = await authInstance.get("/auth/admin/user");
  return res.data;
};

const getUserByIdAdmin = async (id) => {
  const res = await authInstance.get(`/auth/admin/user/${id}`);
  return res.data;
};

export default {
  login,
  logout,
  register,
  forgotPassword,
  resetPassword,
  me,
  changePassword,
  getAllUsers,
  getUserByIdAdmin,
};
