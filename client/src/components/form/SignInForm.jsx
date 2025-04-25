import { useEffect } from "react";
import Input from "@/components/ui/input";
import { useForm } from "react-hook-form";
import { signInSchema } from "@/lib/schema";
import Button from "@//components/ui/button";
import { useSignIn } from "@/hooks/useSignIn";
import { zodResolver } from "@hookform/resolvers/zod";

const SignInForm = () => {
  const { onSubmit, loading } = useSignIn();
  const {
    register,
    setValue,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(signInSchema),
    defaultValues: {
      email: "",
      password: "",
      rememberMe: false,
    },
  });

  useEffect(() => {
    const rememberedEmail = localStorage.getItem("rememberedEmail");
    if (rememberedEmail) {
      setValue("email", rememberedEmail);
      setValue("rememberMe", true);
    }
  }, [setValue]);

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
      <Input
        name="email"
        label="Email"
        register={register}
        error={errors.email}
        placeholder="Masukkan email anda"
      />
      <Input
        name="password"
        label="Password"
        type="password"
        register={register}
        error={errors.password}
        placeholder="*********"
      />

      <div className="flex items-center justify-between">
        <label className="inline-flex items-center text-sm">
          <input
            type="checkbox"
            {...register("rememberMe")}
            className="rounded border-gray-300 text-blue-600 shadow-sm focus:ring-blue-500"
          />
          <span className="ml-2">Remember Me</span>
        </label>
        <a
          href="/forgot-password"
          className="text-sm text-blue-600 hover:underline"
        >
          Forgot Password?
        </a>
      </div>
      <Button
        type="submit"
        variant="primary"
        className="w-full"
        isLoading={loading}
        disabled={loading}
      >
        Sign In
      </Button>
    </form>
  );
};

export default SignInForm;
