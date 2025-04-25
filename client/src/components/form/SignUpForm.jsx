import Input from "@/components/ui/input";
import { useForm } from "react-hook-form";
import Button from "@/components/ui/button";
import { registerSchema } from "@/lib/schema";
import { useSignIn } from "@/hooks/useSignIn";
import { zodResolver } from "@hookform/resolvers/zod";

const SignUpForm = () => {
  const { onSubmit, loading } = useSignIn();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(registerSchema),
  });

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
        name="fullname"
        label="Fullname"
        register={register}
        error={errors.fullname}
        placeholder="Masukkan nama anda"
      />

      <Input
        name="password"
        label="Password"
        type="password"
        register={register}
        error={errors.password}
        placeholder="*********"
      />

      <Button
        variant="primary"
        className="w-full"
        isLoading={loading}
        disabled={loading}
      >
        Create Account
      </Button>
    </form>
  );
};

export default SignUpForm;
