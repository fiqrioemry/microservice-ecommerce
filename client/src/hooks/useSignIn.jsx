import { useAuthStore } from "@/store/useAuthStore";

export const useSignIn = () => {
  const { login, loading } = useAuthStore();

  const onSubmit = async (data) => {
    await login(data);
  };

  return { onSubmit, loading };
};
