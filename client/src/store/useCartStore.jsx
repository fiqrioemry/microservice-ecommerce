import { toast } from "sonner";
import { create } from "zustand";
import carts from "@/services/carts";
import { immer } from "zustand/middleware/immer";

export const useCartStore = create(
  immer((set, get) => ({
    items: [],
    loading: false,

    // 🧠 Getter
    get totalItems() {
      return get().items.reduce((total, item) => total + item.quantity, 0);
    },

    get totalPrice() {
      return get().items.reduce(
        (total, item) => total + item.price * item.quantity,
        0
      );
    },

    // ✅ Fetch cart (initial load)
    fetchCart: async () => {
      set((state) => {
        state.loading = true;
      });
      try {
        const data = await carts.getCart();
        set((state) => {
          state.items = data;
        });
      } catch (err) {
        toast.error("Gagal memuat keranjang");
      } finally {
        set((state) => {
          state.loading = false;
        });
      }
    },

    // ✅ Add new item
    addItem: async (formData) => {
      try {
        const { message } = await carts.addToCart(formData);
        toast.success(message);
        await get().fetchCart(); // refresh setelah add
      } catch (err) {
        toast.error(err.message);
      }
    },

    // ✅ Update item quantity
    updateItem: async (itemId, data) => {
      try {
        const { message } = await carts.updateCartItem(itemId, data);
        toast.success(message);
        await get().fetchCart();
      } catch (err) {
        toast.error(err.message);
      }
    },

    // ✅ Remove single item
    removeItem: async (itemId) => {
      try {
        const { message } = await carts.removeCartItem(itemId);
        toast.success(message);
        set((state) => {
          state.items = state.items.filter((i) => i.id !== itemId);
        });
      } catch (err) {
        toast.error(err.message);
      }
    },

    // ✅ Clear entire cart
    clearCart: async () => {
      try {
        const { message } = await carts.clearCart();
        toast.success(message);
        set((state) => {
          state.items = [];
        });
      } catch (err) {
        toast.error(err.message);
      }
    },
  }))
);
