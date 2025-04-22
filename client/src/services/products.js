// src/api/product.js
import { instance } from ".";

// =====================
// PUBLIC PRODUCT ROUTES
// =====================

const getAllProducts = async (params = {}) => {
  const res = await instance.get("/products", { params });
  return res.data;
};

const getProductBySlug = async (slug) => {
  const res = await instance.get(`/products/${slug}`);
  return res.data;
};

const searchProducts = async (queryParams) => {
  const res = await instance.get("/products/search", { params: queryParams });
  return res.data;
};

// =====================
// ADMIN PRODUCT ROUTES
// =====================

const createProduct = async (formData) => {
  const res = await instance.post("/products/admin", formData);
  return res.data;
};

const updateProduct = async (id, formData) => {
  const res = await instance.put(`/products/admin/${id}`, formData);
  return res.data;
};

const deleteProduct = async (id) => {
  const res = await instance.delete(`/products/admin/${id}`);
  return res.data;
};

const uploadLocalImage = async (formData) => {
  const res = await instance.post("/products/admin/upload-local", formData);
  return res.data;
};

const downloadImage = async (id) => {
  const res = await instance.get(`/products/admin/${id}/download`, {
    responseType: "blob",
  });
  return res.data;
};

const deleteVariantProduct = async (variantId) => {
  const res = await instance.delete(`/products/admin/variant/${variantId}`);
  return res.data;
};

export default {
  getAllProducts,
  getProductBySlug,
  searchProducts,
  createProduct,
  updateProduct,
  deleteProduct,
  uploadLocalImage,
  downloadImage,
  deleteVariantProduct,
};
