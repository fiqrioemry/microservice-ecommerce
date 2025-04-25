import axios from "axios";
import { toast } from "sonner";
import { useAuthStore } from "@/store/useAuthStore";

export const productInstance = axios.create({
  baseURL: import.meta.env.VITE_PRODUCT_SERVICES,
  withCredentials: true,
  headers: {
    "X-API-Key": import.meta.env.VITE_API_KEY,
  },
});

export const userInstance = axios.create({
  baseURL: import.meta.env.VITE_USER_SERVICES,
  withCredentials: true,
  headers: {
    "X-API-Key": import.meta.env.VITE_API_KEY,
  },
});

export const cartInstance = axios.create({
  baseURL: import.meta.env.VITE_CART_SERVICES,
  withCredentials: true,
  headers: {
    "X-API-Key": import.meta.env.VITE_API_KEY,
  },
});

userInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      const { logout, user } = useAuthStore.getState();
      if (user) {
        toast.error("Session expired, please login again.");
        logout();
      }
    }
    return Promise.reject(error);
  }
);

cartInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      const { logout, user } = useAuthStore.getState();
      if (user) {
        toast.error("Session expired, please login again.");
        logout();
      }
    }
    return Promise.reject(error);
  }
);
