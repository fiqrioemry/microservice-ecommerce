// src/hooks/useAuthMutation.ts
import { toast } from "sonner";
import authService from "@/services/auth";
import { useNavigate } from "react-router-dom";
import { useMutation } from "@tanstack/react-query";
import { useAuthStore } from "@/store/useAuthStore";

const getErrorMessage = (error) =>
  error?.response?.data?.message || "Something went wrong!";

export const useLogin = () => {
  const navigate = useNavigate();
  const setUser = useAuthStore((state) => state.setUser);

  return useMutation({
    mutationFn: authService.login,
    onSuccess: ({ user, message }) => {
      setUser(user);
      toast.success(message || "Login successful");
      navigate("/");
    },
    onError: (error) => {
      toast.error(getErrorMessage(error));
    },
  });
};

export const useRegister = () => {
  const navigate = useNavigate();

  return useMutation({
    mutationFn: authService.register,
    onSuccess: ({ message }) => {
      toast.success(message || "Register successful");
      navigate("/signin");
    },
    onError: (error) => {
      toast.error(getErrorMessage(error));
    },
  });
};

export const useLogout = () => {
  const navigate = useNavigate();
  const clearUser = useAuthStore((state) => state.clearUser);

  return useMutation({
    mutationFn: authService.logout,
    onSuccess: () => {
      clearUser();
      toast.success("Logged out successfully");
      navigate("/signin");
    },
    onError: (error) => {
      toast.error(getErrorMessage(error));
    },
  });
};
