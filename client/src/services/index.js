import axios from "axios";
import { toast } from "sonner";

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

cartInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      const logout = useAuthStore.getState().logout;
      toast.error("Session expired, please login again.");
      logout();
    }
    return Promise.reject(error);
  }
);

userInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      const logout = useAuthStore.getState().logout;
      toast.error("Session expired, please login again.");
      logout();
    }
    return Promise.reject(error);
  }
);

productInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      const logout = useAuthStore.getState().logout;
      toast.error("Session expired, please login again.");
      logout();
    }
    return Promise.reject(error);
  }
);

cartInstance.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      const logout = useAuthStore.getState().logout;
      toast.error("Session expired, please login again.");
      logout();
    }
    return Promise.reject(error);
  }
);
