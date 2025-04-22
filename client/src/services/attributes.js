// src/api/attribute.js
import { instance } from ".";

// =====================
// PUBLIC
// =====================

const getAllAttributes = async () => {
  const res = await instance.get("/attributes");
  return res.data;
};

// =====================
// ADMIN
// =====================

const createAttribute = async (data) => {
  const res = await instance.post("/attributes", data);
  return res.data;
};

const updateAttribute = async (id, data) => {
  const res = await instance.put(`/attributes/${id}`, data);
  return res.data;
};

const deleteAttribute = async (id) => {
  const res = await instance.delete(`/attributes/${id}`);
  return res.data;
};

const addAttributeValue = async (attributeId, data) => {
  const res = await instance.post(`/attributes/${attributeId}/values`, data);
  return res.data;
};

const updateAttributeValue = async (valueId, data) => {
  const res = await instance.put(`/attributes/values/${valueId}`, data);
  return res.data;
};

const deleteAttributeValue = async (valueId) => {
  const res = await instance.delete(`/attributes/values/${valueId}`);
  return res.data;
};

export default {
  getAllAttributes,
  createAttribute,
  updateAttribute,
  deleteAttribute,
  addAttributeValue,
  updateAttributeValue,
  deleteAttributeValue,
};
