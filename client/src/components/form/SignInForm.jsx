import { useForm } from "react-hook-form";
import { signInSchema } from "@/lib/schema";
import { useSignIn } from "@/hooks/useSignIn";
import { zodResolver } from "@hookform/resolvers/zod";

const SignInForm = () => {
  const { onSubmit, loading } = useSignIn();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(signInSchema),
  });

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div>
        <label className="block text-sm font-medium">Email</label>
        <input
          type="email"
          {...register("email")}
          className="w-full border rounded p-2"
        />
        {errors.email && (
          <p className="text-red-500 text-xs">{errors.email.message}</p>
        )}
      </div>

      <div>
        <label className="block text-sm font-medium">Password</label>
        <input
          type="password"
          {...register("password")}
          className="w-full border rounded p-2"
        />
        {errors.password && (
          <p className="text-red-500 text-xs">{errors.password.message}</p>
        )}
      </div>

      <button
        type="submit"
        disabled={loading}
        className="btn btn-primary w-full"
      >
        {loading ? "Loading..." : "Sign In"}
      </button>
    </form>
  );
};

export default SignInForm;
