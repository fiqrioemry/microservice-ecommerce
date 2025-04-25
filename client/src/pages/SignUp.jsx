import { Link } from "react-router-dom";
import SignUpForm from "@/components/form/SignUpForm";

const SignUp = () => {
  return (
    <section className="flex justify-center items-center min-h-screen bg-gray-100">
      <div className="grid grid-cols-1 md:grid-cols-2 bg-white rounded-xl shadow-lg overflow-hidden max-w-4xl w-full">
        <div className="hidden md:block bg-blue-600  p-8 text-white text-center">
          <h2 className="text-3xl font-bold mb-4">Join Us!</h2>
          <p className="text-sm">Create your account to start shopping now</p>
          <img
            src="/signup-wallpaper.webp"
            alt="Register Illustration"
            className="mt-6 w-full h-auto"
          />
        </div>

        <div className="p-8">
          <Link to="/">
            <h1 className="text-center text-primary">Ecommerce</h1>
          </Link>
          <h2 className="text-center">Register</h2>

          <SignUpForm />
          <p className="text-sm text-center mt-6 text-gray-600">
            Sudah punya akun?{" "}
            <Link
              to="/signin"
              className="text-indigo-600 hover:underline font-medium"
            >
              Masuk sekarang
            </Link>
          </p>
        </div>
      </div>
    </section>
  );
};

export default SignUp;
