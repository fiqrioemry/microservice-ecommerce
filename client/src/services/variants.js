// src/api/variant.js
import { instance } from ".";

// =====================
// PUBLIC
// =====================

const getAllVariantTypes = async () => {
  const res = await instance.get("/variants");
  return res.data;
};

// =====================
// ADMIN
// =====================

const createVariantType = async (data) => {
  const res = await instance.post("/variants", data);
  return res.data;
};

const updateVariantType = async (id, data) => {
  const res = await instance.put(`/variants/${id}`, data);
  return res.data;
};

const deleteVariantType = async (id) => {
  const res = await instance.delete(`/variants/${id}`);
  return res.data;
};

const addVariantValue = async (typeId, data) => {
  const res = await instance.post(`/variants/${typeId}/values`, data);
  return res.data;
};

const updateVariantValue = async (valueId, data) => {
  const res = await instance.put(`/variants/values/${valueId}`, data);
  return res.data;
};

const deleteVariantValue = async (valueId) => {
  const res = await instance.delete(`/variants/values/${valueId}`);
  return res.data;
};

const mapToCategory = async (data) => {
  const res = await instance.post("/variants/map/category", data);
  return res.data;
};

const mapToSubcategory = async (data) => {
  const res = await instance.post("/variants/map/subcategory", data);
  return res.data;
};

export default {
  getAllVariantTypes,
  createVariantType,
  updateVariantType,
  deleteVariantType,
  addVariantValue,
  updateVariantValue,
  deleteVariantValue,
  mapToCategory,
  mapToSubcategory,
};
