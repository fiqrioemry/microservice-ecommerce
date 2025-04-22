import { create } from "zustand";
import { persist } from "zustand/middleware";
import { toast } from "sonner";
import auth from "@/services/auth";

export const useAuthStore = create(
  persist(
    (set) => ({
      user: null,
      loading: false,
      checkingAuth: true,

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
          await auth.logout(); // optional: you can skip this if token in cookies
        } catch (err) {
          console.warn("Failed to logout", err);
        }
        set({ user: null });
        window.location.href = "/signin";
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
    }),
    {
      name: "auth-storage", // localStorage key
      partialize: (state) => ({ user: state.user }), // only persist user
    }
  )
);
