// src/location.ts
import { userInstance } from ".";

// Get all provinces
const getProvinces = async () => {
  const res = await userInstance.get("/provinces");
  return res.data;
};

// Get cities by province ID
const getCitiesByProvince = async (provinceId) => {
  const res = await userInstance.get(`/provinces/${provinceId}/cities`);
  return res.data;
};

export default {
  getProvinces,
  getCitiesByProvince,
};
