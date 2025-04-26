// src/location.js
import { userInstance } from ".";

// Get all provinces
const getProvinces = async () => {
  const res = await userInstance.get("/locations/provinces");
  return res.data;
};

// Search provinces by name
const searchProvinces = async (query) => {
  const res = await userInstance.get(`/locations/provinces/search?q=${query}`);
  return res.data;
};

// Get cities by province ID
const getCitiesByProvince = async (provinceId) => {
  const res = await userInstance.get(
    `/locations/provinces/${provinceId}/cities`
  );
  return res.data;
};

// Search cities by name
const searchCities = async (query) => {
  const res = await userInstance.get(`/locations/cities/search?q=${query}`);
  return res.data;
};

// Get districts by city ID
const getDistrictsByCity = async (cityId) => {
  const res = await userInstance.get(`/locations/cities/${cityId}/districts`);
  return res.data;
};

// Get subdistricts by district ID
const getSubdistrictsByDistrict = async (districtId) => {
  const res = await userInstance.get(
    `/locations/districts/${districtId}/subdistricts`
  );
  return res.data;
};

// Get postal codes by subdistrict ID
const getPostalCodesBySubdistrict = async (subdistrictId) => {
  const res = await userInstance.get(
    `/locations/subdistricts/${subdistrictId}/postalcodes`
  );
  return res.data;
};

export default {
  getProvinces,
  searchProvinces,
  getCitiesByProvince,
  searchCities,
  getDistrictsByCity,
  getSubdistrictsByDistrict,
  getPostalCodesBySubdistrict,
};
