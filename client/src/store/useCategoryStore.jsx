// 2. stores/useCategoryStore.jsx
import { create } from "zustand";
import toast from "react-hot-toast";
import categoryApi from "@/api/categories";
import { devtools } from "zustand/middleware";
import { queryClient } from "@/lib/react-query";

export const useCategoryStore = create(
  devtools((set) => ({
    loading: false,

    createCategory: async (data) => {
      set({ loading: true });
      try {
        await categoryApi.createCategory(data);
        toast.success("Category created successfully");
        queryClient.invalidateQueries({ queryKey: ["categories"] });
      } catch (error) {
        toast.error(error.message);
      } finally {
        set({ loading: false });
      }
    },

    updateCategory: async (id, data) => {
      set({ loading: true });
      try {
        await categoryApi.updateCategory(id, data);
        toast.success("Category updated successfully");
        queryClient.invalidateQueries({ queryKey: ["categories"] });
      } catch (error) {
        toast.error(error.message);
      } finally {
        set({ loading: false });
      }
    },

    deleteCategory: async (id) => {
      set({ loading: true });
      try {
        await categoryApi.deleteCategory(id);
        toast.success("Category deleted successfully");
        queryClient.invalidateQueries({ queryKey: ["categories"] });
      } catch (error) {
        toast.error(error.message);
      } finally {
        set({ loading: false });
      }
    },

    createSubcategory: async (categoryId, data) => {
      set({ loading: true });
      try {
        await categoryApi.createSubcategory(categoryId, data);
        toast.success("Subcategory created successfully");
        queryClient.invalidateQueries({ queryKey: ["categories"] });
      } catch (error) {
        toast.error(error.message);
      } finally {
        set({ loading: false });
      }
    },

    updateSubcategory: async (subId, data) => {
      set({ loading: true });
      try {
        await categoryApi.updateSubcategory(subId, data);
        toast.success("Subcategory updated successfully");
        queryClient.invalidateQueries({ queryKey: ["categories"] });
      } catch (error) {
        toast.error(error.message);
      } finally {
        set({ loading: false });
      }
    },

    deleteSubcategory: async (subId) => {
      set({ loading: true });
      try {
        await categoryApi.deleteSubcategory(subId);
        toast.success("Subcategory deleted successfully");
        queryClient.invalidateQueries({ queryKey: ["categories"] });
      } catch (error) {
        toast.error(error.message);
      } finally {
        set({ loading: false });
      }
    },
  }))
);
