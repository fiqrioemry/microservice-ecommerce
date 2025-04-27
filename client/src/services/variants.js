// src/api/variant.js

import { product } from ".";

// =====================
// PUBLIC
// =====================

const getAllVariantTypes = async () => {
  const res = await product.get("/variants");
  return res.data;
};

// =====================
// ADMIN
// =====================

const createVariantType = async (data) => {
  const res = await product.post("/variants", data);
  return res.data;
};

const updateVariantType = async (id, data) => {
  const res = await product.put(`/variants/${id}`, data);
  return res.data;
};

const deleteVariantType = async (id) => {
  const res = await product.delete(`/variants/${id}`);
  return res.data;
};

const addVariantValue = async (typeId, data) => {
  const res = await product.post(`/variants/${typeId}/values`, data);
  return res.data;
};

const updateVariantValue = async (valueId, data) => {
  const res = await product.put(`/variants/values/${valueId}`, data);
  return res.data;
};

const deleteVariantValue = async (valueId) => {
  const res = await product.delete(`/variants/values/${valueId}`);
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
