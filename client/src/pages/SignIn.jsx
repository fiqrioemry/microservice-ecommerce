import SignInForm from "@/components/auth/SignInForm";

const SignIn = () => {
  return (
    <section className="flex justify-center items-center h-screen bg-gray-50">
      <div className="max-w-md w-full bg-white p-8 rounded shadow">
        <h1 className="text-2xl font-bold mb-4 text-center">Sign In</h1>
        <SignInForm />
      </div>
    </section>
  );
};

export default SignIn;
