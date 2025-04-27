import { create } from "zustand";

export const useCartStore = create((set, get) => ({
  items: [],
  totalItems: 0,
  totalPrice: 0,

  getTotalItems: () => {
    const totalItems =
      get().items?.reduce((total, item) => total + item.quantity, 0) || 0;
    set({ totalItems });
  },

  getTotalPrice: () => {
    const totalPrice =
      get().items?.reduce(
        (total, item) => total + item.price * item.quantity,
        0
      ) || 0;
    set({ totalPrice });
  },

  setCart: (items) => {
    set({ items });
    get().getTotalItems();
    get().getTotalPrice();
  },

  addItem: (item) => {
    set((state) => ({
      items: [...state.items, item],
    }));
    get().getTotalItems();
    get().getTotalPrice();
  },

  updateItem: (itemId, quantity) => {
    set((state) => ({
      items: state.items.map((item) =>
        item.id === itemId ? { ...item, quantity } : item
      ),
    }));
    get().getTotalItems();
    get().getTotalPrice();
  },

  removeItem: (itemId) => {
    set((state) => ({
      items: state.items.filter((item) => item.id !== itemId),
    }));
    get().getTotalItems();
    get().getTotalPrice();
  },

  clearCart: () => {
    set({ items: [] });
    set({ totalItems: 0, totalPrice: 0 });
  },
}));
