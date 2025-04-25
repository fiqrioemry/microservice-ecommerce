// =====================
// PROFILE (Auth Required)
// =====================

import { userInstance } from ".";

const getProfile = async () => {
  const res = await userInstance.get("/user/profile");
  return res.data;
};

const updateProfile = async (data) => {
  const res = await userInstance.put("/user/profile", data);
  return res.data;
};

// =====================
// ADDRESS (Auth Required)
// =====================

const getAddresses = async () => {
  const res = await userInstance.get("/user/addresses");
  return res.data;
};

const addAddress = async (data) => {
  const res = await userInstance.post("/user/addresses", data);
  return res.data;
};

const updateAddress = async (id, data) => {
  const res = await userInstance.put(`/user/addresses/${id}`, data);
  return res.data;
};

const deleteAddress = async (id) => {
  const res = await userInstance.delete(`/user/addresses/${id}`);
  return res.data;
};

const setMainAddress = async (id) => {
  const res = await userInstance.put(`/user/addresses/${id}/set-main`);
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
