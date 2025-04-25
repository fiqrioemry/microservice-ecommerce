import { toast } from "sonner";
import { create } from "zustand";
import carts from "@/services/carts";
import { immer } from "zustand/middleware/immer";

// src/store/useCartStore.jsx

export const useCartStore = create(
  immer((set, get) => ({
    carts: {},
    totalItems: 0,
    totalPrice: 0,
    loading: false,

    getTotalItems: () => {
      const totalItems =
        get().carts.items?.reduce((total, item) => total + item.quantity, 0) ||
        0;
      set({ totalItems });
    },

    getTotalPrice: () => {
      const totalPrice =
        get().carts.items?.reduce(
          (total, item) => total + item.price * item.quantity,
          0
        ) || 0;
      set({ totalPrice });
    },

    fetchCart: async () => {
      set({
        loading: true,
      });
      try {
        const { cart } = await carts.getCart();
        set({ carts: cart });
        get().getTotalItems();
        get().getTotalPrice();
      } catch (err) {
        toast.error("Gagal memuat keranjang");
      } finally {
        set({
          loading: false,
        });
      }
    },

    addItem: async (formData) => {
      try {
        const { message } = await carts.addToCart(formData);
        toast.success(message);
        await get().fetchCart();
      } catch (err) {
        toast.error(err.message);
      }
    },

    updateItem: async (itemId, data) => {
      try {
        console.log(itemId, data);
        const { message } = await carts.updateCartItem(itemId, data);
        toast.success(message);
        await get().fetchCart();
      } catch (err) {
        toast.error(err.message);
      }
    },

    removeItem: async (itemId) => {
      try {
        const { message } = await carts.removeCartItem(itemId);
        toast.success(message);
        set((state) => {
          state.carts.items =
            state.carts.items?.filter((i) => i.id !== itemId) || [];
        });
      } catch (err) {
        toast.error(err.message);
      }
    },

    clearCart: async () => {
      try {
        const { message } = await carts.clearCart();
        toast.success(message);
        set((state) => {
          state.carts = { items: [] };
        });
      } catch (err) {
        toast.error(err.message);
      }
    },
  }))
);
