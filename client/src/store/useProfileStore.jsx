// src/store/useProfileStore.jsx
import { toast } from "sonner";
import { create } from "zustand";
import users from "@/services/users";

export const useProfileStore = create((set) => ({
  loading: false,

  addAddress: async (data) => {
    set({ loading: true });
    try {
      const { message } = await users.addAddress(data);
      toast.success(message);
    } catch (error) {
      console.log(error);
      toast.error(error?.response?.data?.message);
    } finally {
      set({ loading: false });
    }
  },

  updateAddress: async (id, data) => {
    set({ loading: true });
    try {
      const { message } = await users.updateAddress(id, data);
      toast.success(message);
    } catch (error) {
      toast.error(error?.response?.data?.message);
    } finally {
      set({ loading: false });
    }
  },

  deleteAddress: async (id) => {
    set({ loading: true });
    try {
      const { message } = await users.deleteAddress(id);
      toast.success(message);
    } catch (error) {
      toast.error(error?.response?.data?.message);
    } finally {
      set({ loading: false });
    }
  },

  setMainAddress: async (id) => {
    set({ loading: true });
    try {
      await users.setMainAddress(id);
      toast.success("Main address set");
    } catch (error) {
      toast.error(error?.response?.data?.message);
    } finally {
      set({ loading: false });
    }
  },
}));
