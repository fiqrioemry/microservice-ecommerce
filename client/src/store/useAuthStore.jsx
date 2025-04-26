import { create } from "zustand";
import { toast } from "sonner";
import auth from "@/services/auth";
import { persist } from "zustand/middleware";

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
          if (formData.rememberMe) {
            localStorage.setItem("rememberme", formData.email);
          } else {
            localStorage.removeItem("rememberme");
          }
          set({ user });
          toast.success(message);
          window.location.href = "/";
        } catch (error) {
          toast.error(error.response.data.message);
        } finally {
          set({ loading: false });
        }
      },

      logout: async () => {
        try {
          await auth.logout();
          set({ user: null });
        } catch (err) {
          console.warn("Failed to logout", err);
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
    }),
    {
      name: "auth-storage",
      partialize: (state) => ({ user: state.user }), // only persist user
    }
  )
);
