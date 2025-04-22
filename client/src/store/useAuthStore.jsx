// src/store/useAuthStore.js
import { create } from "zustand";
import { toast } from "sonner";
import auth from "@/api/auth";

export const useAuthStore = create((set) => ({
  user: null,
  users: [],
  loading: false,
  checkingAuth: true,

  // Untuk cek apakah user masih login
  authCheck: async () => {
    try {
      const { user } = await auth.me();
      set({ user });
    } catch {
      set({ user: null });
    } finally {
      set({ checkingAuth: false });
    }
  },

  login: async (formData) => {
    set({ loading: true });
    try {
      const { message, user } = await auth.login(formData);
      set({ user });
      toast.success(message);
      window.location.href = "/";
    } catch (error) {
      toast.error(error.message);
    } finally {
      set({ loading: false });
    }
  },

  logout: async () => {
    try {
      const { message } = await auth.logout();
      toast.success(message);
      set({ user: null });
      window.location.href = "/signin";
    } catch (error) {
      console.warn(error.message);
    }
  },

  register: async (formData) => {
    set({ loading: true });
    try {
      const { message } = await auth.register(formData);
      toast.success(message);
      window.location.href = "/signin";
    } catch (error) {
      toast.error(error.message);
    } finally {
      set({ loading: false });
    }
  },

  forgotPassword: async (formData) => {
    set({ loading: true });
    try {
      const { message } = await auth.forgotPassword(formData);
      toast.success(message);
    } catch (error) {
      toast.error(error.message);
    } finally {
      set({ loading: false });
    }
  },

  resetPassword: async (formData) => {
    set({ loading: true });
    try {
      const { message } = await auth.resetPassword(formData);
      toast.success(message);
      window.location.href = "/signin";
    } catch (error) {
      toast.error(error.message);
    } finally {
      set({ loading: false });
    }
  },

  changePassword: async (formData) => {
    set({ loading: true });
    try {
      const { message } = await auth.changePassword(formData);
      toast.success(message);
    } catch (error) {
      toast.error(error.message);
    } finally {
      set({ loading: false });
    }
  },

  getAllUsers: async () => {
    try {
      const users = await auth.getAllUsers();
      set({ users });
    } catch (error) {
      toast.error("Gagal memuat data user");
    }
  },

  getUserByIdAdmin: async (id) => {
    try {
      const user = await auth.getUserByIdAdmin(id);
      return user; // langsung return data (tidak perlu simpan ke state)
    } catch (error) {
      toast.error("Gagal memuat data user");
    }
  },
}));
