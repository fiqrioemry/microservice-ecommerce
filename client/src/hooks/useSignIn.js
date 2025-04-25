import { useAuthStore } from "@/store/useAuthStore";

export const useSignIn = () => {
  const { login, loading } = useAuthStore();

  const onSubmit = async ({ rememberMe, ...data }) => {
    if (rememberMe) {
      localStorage.setItem("rememberedEmail", data.email);
    } else {
      localStorage.removeItem("rememberedEmail");
    }

    await login(data);
  };

  return { onSubmit, loading };
};
