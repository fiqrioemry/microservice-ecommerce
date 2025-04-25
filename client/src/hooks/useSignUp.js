import { useAuthStore } from "@/store/useAuthStore";

export const useSignUp = () => {
  const { register, loading } = useAuthStore();

  const onSubmit = async (formData) => {
    await register(formData);
  };

  return { onSubmit, loading };
};
