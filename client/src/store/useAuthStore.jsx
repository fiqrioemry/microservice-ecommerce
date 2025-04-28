// src/store/useAuthStore.ts
import { create } from "zustand";
import auth from "@/services/auth";
import { persist } from "zustand/middleware";

export const useAuthStore = create(
  persist(
    (set) => ({
      user: null,
      checkingAuth: true,

      setUser: (user) => set({ user }),

      clearUser: () => set({ user: null }),

      authCheck: async () => {
        try {
          const user = await auth.me();
          set({ user });
        } catch {
          set({ user: null });
        } finally {
          set({ checkingAuth: false });
        }
      },
    }),
    {
      name: "auth-storage",
      partialize: (state) => ({ user: state.user }),
    }
  )
);
