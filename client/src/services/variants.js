// src/api/variant.js

import { productInstance } from ".";
// =====================
// PUBLIC
// =====================

const getAllVariantTypes = async () => {
  const res = await productInstance.get("/variants");
  return res.data;
};

// =====================
// ADMIN
// =====================

const createVariantType = async (data) => {
  const res = await productInstance.post("/variants", data);
  return res.data;
};

const updateVariantType = async (id, data) => {
  const res = await productInstance.put(`/variants/${id}`, data);
  return res.data;
};

const deleteVariantType = async (id) => {
  const res = await productInstance.delete(`/variants/${id}`);
  return res.data;
};

const addVariantValue = async (typeId, data) => {
  const res = await productInstance.post(`/variants/${typeId}/values`, data);
  return res.data;
};

const updateVariantValue = async (valueId, data) => {
  const res = await productInstance.put(`/variants/values/${valueId}`, data);
  return res.data;
};

const deleteVariantValue = async (valueId) => {
  const res = await productInstance.delete(`/variants/values/${valueId}`);
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
};
