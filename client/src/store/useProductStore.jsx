import { create } from "zustand";
import { immer } from "zustand/middleware/immer";
import { toast } from "sonner";
import productApi from "@/services/product";

export const useProductStore = create(
  immer((set, get) => ({
    products: [],
    product: null,
    loading: false,

    // =====================
    // PUBLIC PRODUCT METHODS
    // =====================

    fetchAllProducts: async (params = {}) => {
      set((state) => {
        state.loading = true;
      });
      try {
        const data = await productApi.getAllProducts(params);
        set((state) => {
          state.products = data;
        });
      } catch (err) {
        toast.error("Gagal memuat produk");
      } finally {
        set((state) => {
          state.loading = false;
        });
      }
    },

    fetchProductBySlug: async (slug) => {
      set((state) => {
        state.loading = true;
      });
      try {
        const data = await productApi.getProductBySlug(slug);
        set((state) => {
          state.product = data;
        });
      } catch (err) {
        toast.error("Produk tidak ditemukan");
      } finally {
        set((state) => {
          state.loading = false;
        });
      }
    },

    searchProducts: async (queryParams) => {
      try {
        const data = await productApi.searchProducts(queryParams);
        set((state) => {
          state.products = data;
        });
      } catch (err) {
        toast.error("Gagal mencari produk");
      }
    },

    // =====================
    // ADMIN PRODUCT METHODS
    // =====================

    createProduct: async (formData) => {
      try {
        const { message } = await productApi.createProduct(formData);
        toast.success(message);
        await get().fetchAllProducts(); // refresh data
      } catch (err) {
        toast.error(err.message);
      }
    },

    updateProduct: async (id, formData) => {
      try {
        const { message } = await productApi.updateProduct(id, formData);
        toast.success(message);
        await get().fetchAllProducts();
      } catch (err) {
        toast.error(err.message);
      }
    },

    deleteProduct: async (id) => {
      try {
        const { message } = await productApi.deleteProduct(id);
        toast.success(message);
        await get().fetchAllProducts();
      } catch (err) {
        toast.error(err.message);
      }
    },

    uploadLocalImage: async (formData) => {
      try {
        const { message } = await productApi.uploadLocalImage(formData);
        toast.success(message);
      } catch (err) {
        toast.error(err.message);
      }
    },

    downloadImage: async (id) => {
      try {
        const blob = await productApi.downloadImage(id);
        const url = window.URL.createObjectURL(blob);
        const link = document.createElement("a");
        link.href = url;
        link.download = `product-image-${id}`;
        document.body.appendChild(link);
        link.click();
        link.remove();
        window.URL.revokeObjectURL(url);
      } catch (err) {
        toast.error("Gagal mengunduh gambar");
      }
    },

    deleteVariantProduct: async (variantId) => {
      try {
        const { message } = await productApi.deleteVariantProduct(variantId);
        toast.success(message);
        await get().fetchAllProducts();
      } catch (err) {
        toast.error(err.message);
      }
    },
  }))
);
