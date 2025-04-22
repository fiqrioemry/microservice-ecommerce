// src/location.ts
import { publicInstance } from ".";

// Get all provinces
const getProvinces = async () => {
  const res = await publicInstance.get("/provinces");
  return res.data;
};

// Get cities by province ID
const getCitiesByProvince = async (provinceId) => {
  const res = await publicInstance.get(`/provinces/${provinceId}/cities`);
  return res.data;
};

export default {
  getProvinces,
  getCitiesByProvince,
};
