import FormInput from "@/components/form/FormInput";
import { useAuthStore } from "@/store/useAuthStore";
import { loginFormSchema } from "@/components/form/constant";
import { loginSchema } from "@/components/form/schema";

const Home = () => {
  const { submitTest } = useAuthStore();

  return (
    <section className="container h-screen mx-auto">
      <div className="px-2 space-y-4 py-3 md:py-6">
        <FormInput
          action={submitTest}
          submitText="Login"
          schema={loginFormSchema}
          validation={loginSchema}
        />
      </div>
    </section>
  );
};

export default Home;
