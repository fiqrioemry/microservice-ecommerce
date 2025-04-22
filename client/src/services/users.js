// src/api/user.js
import { instance } from ".";

// =====================
// PROFILE (Auth Required)
// =====================

const getProfile = async () => {
  const res = await instance.get("/user/profile");
  return res.data;
};

const updateProfile = async (data) => {
  const res = await instance.put("/user/profile", data);
  return res.data;
};

// =====================
// ADDRESS (Auth Required)
// =====================

const getAddresses = async () => {
  const res = await instance.get("/user/addresses");
  return res.data;
};

const addAddress = async (data) => {
  const res = await instance.post("/user/addresses", data);
  return res.data;
};

const updateAddress = async (id, data) => {
  const res = await instance.put(`/user/addresses/${id}`, data);
  return res.data;
};

const deleteAddress = async (id) => {
  const res = await instance.delete(`/user/addresses/${id}`);
  return res.data;
};

const setMainAddress = async (id) => {
  const res = await instance.put(`/user/addresses/${id}/set-main`);
  return res.data;
};

export default {
  getProfile,
  updateProfile,
  getAddresses,
  addAddress,
  updateAddress,
  deleteAddress,
  setMainAddress,
};
