// src/store/useCartStore.js
import { create } from "zustand";

export const useCartStore = create((set) => ({
  items: [],

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

  setCart: (items) => set({ items }),
  addItem: (item) =>
    set((state) => ({
      items: [...state.items, item],
    })),
  updateItem: (itemId, quantity) =>
    set((state) => ({
      items: state.items.map((item) =>
        item.id === itemId ? { ...item, quantity } : item
      ),
    })),
  removeItem: (itemId) =>
    set((state) => ({
      items: state.items.filter((item) => item.id !== itemId),
    })),
  clearCart: () => set({ items: [] }),
}));
