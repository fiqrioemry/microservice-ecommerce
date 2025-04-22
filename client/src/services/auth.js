// src/auth.ts
import { instance } from ".";

// Auth - Public
const login = async (data) => {
  const res = await instance.post("/auth/login", data);
  return res.data;
};

const logout = async () => {
  const res = await instance.post("/auth/logout");
  return res.data;
};

const register = async (data) => {
  const res = await instance.post("/auth/register", data);
  return res.data;
};

const forgotPassword = async (data) => {
  const res = await instance.post("/auth/forgot-password", data);
  return res.data;
};

const resetPassword = async (data) => {
  const res = await instance.post("/auth/reset-password", data);
  return res.data;
};

// Auth - User (Protected)
const me = async () => {
  const res = await instance.get("/auth/me");
  return res.data;
};

const changePassword = async (data) => {
  const res = await instance.put("/auth/change-password", data);
  return res.data;
};

// Admin - User Management (Protected)
const getAllUsers = async () => {
  const res = await instance.get("/auth/admin/user");
  return res.data;
};

const getUserByIdAdmin = async (id) => {
  const res = await instance.get(`/auth/admin/user/${id}`);
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
