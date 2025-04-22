// src/category.ts
import { instance } from ".";

// Public: Get all categories
const getAllCategories = async () => {
  const res = await instance.get("/categories");
  return res.data;
};

// Admin: Create a new category
const createCategory = async (data) => {
  const res = await instance.post("/admin/categories", data);
  return res.data;
};

// Admin: Update category by ID
const updateCategory = async (id, data) => {
  const res = await instance.put(`/admin/categories/${id}`, data);
  return res.data;
};

// Admin: Delete category by ID
const deleteCategory = async (id) => {
  const res = await instance.delete(`/admin/categories/${id}`);
  return res.data;
};

// Admin: Create subcategory for a category
const createSubcategory = async (categoryId, data) => {
  const res = await instance.post(
    `/admin/categories/${categoryId}/subcategories`,
    data
  );
  return res.data;
};

// Admin: Update subcategory by ID
const updateSubcategory = async (subId, data) => {
  const res = await instance.put(
    `/admin/categories/subcategories/${subId}`,
    data
  );
  return res.data;
};

// Admin: Delete subcategory by ID
const deleteSubcategory = async (subId) => {
  const res = await instance.delete(`/admin/categories/subcategories/${subId}`);
  return res.data;
};

export default {
  getAllCategories,
  createCategory,
  updateCategory,
  deleteCategory,
  createSubcategory,
  updateSubcategory,
  deleteSubcategory,
};
