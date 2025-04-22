import { toast } from "sonner";
import { create } from "zustand";
import product from "../api/product";

export const useProductStore = create((set) => ({
  categories: null,
  provinces: null,
  cities: null,
  subCategories: null,
  checkingAuth: true,

  submitTest: async (formData) => {
    const content = Object.entries(formData)
      .map(([key, value]) => `${key}: ${value}`)
      .join("\n");

    toast("Result:", {
      description: content,
    });
  },

  getCategories: async () => {
    try {
      const { categories } = await product.getCategories();
      set({ categories });
    } catch {
      set({ categories: [] });
    }
  },

  getSubCategories: async () => {
    try {
      const { subCategories } = await product.getsubCategories();
      set({ subCategories });
    } catch {
      set({ SubCategories: [] });
    }
  },

  getProvinces: async () => {
    try {
      const { provinces } = await product.getProvinces();
      set({ provinces });
    } catch {
      set({ provinces: [] });
    }
  },

  getCities: async () => {
    try {
      const { cities } = await product.getProvinces();
      set({ cities });
    } catch {
      set({ cities: [] });
    }
  },

  loadOptions: async (key, dependentValue = false) => {
    switch (key) {
      case "province":
        const { provinces } = await product.getProvinces();
        return provinces.map((p) => ({ label: p.name, value: p.id }));
  
      case "city":
        if (dependentValue === true) return [];
        const { cities } = await product.getCities(dependentValue);
        return cities.map((c) => ({ label: c.name, value: c.id }));
  
      default:
        return [];
    }
}));
