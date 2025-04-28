// src/App.jsx
import Home from "./pages/Home";
import SignIn from "./pages/SignIn";
import SignUp from "./pages/SignUp";
import NotFound from "./pages/NotFound";
import CartPage from "./pages/CartPage";
import Checkout from "./pages/Checkout";
import Orders from "./pages/customer/Orders";
import Profile from "./pages/customer/Profile";
import Address from "./pages/customer/Address";
import ProductDetail from "./pages/ProductDetail";
import ProductResults from "./pages/ProductResults";

import { Toaster } from "sonner";
import { useEffect } from "react";
import Loading from "@/components/ui/Loading";
import Layout from "./components/layout/Layout";
import { useAuthStore } from "@/store/useAuthStore";
import { AuthRoute, NonAuthRoute } from "./middleware";
import UserLayout from "./components/layout/UserLayout";
import { Navigate, Route, Routes } from "react-router-dom";

function App() {
  const { authCheck, checkingAuth } = useAuthStore();

  useEffect(() => {
    authCheck();
  }, [authCheck]);

  if (checkingAuth) return <Loading />;
  return (
    <>
      <Toaster />
      <Routes>
        <Route
          path="/signin"
          element={
            <NonAuthRoute>
              <SignIn />
            </NonAuthRoute>
          }
        />
        <Route
          path="/signup"
          element={
            <NonAuthRoute>
              <SignUp />
            </NonAuthRoute>
          }
        />

        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route
            path="cart"
            element={
              <AuthRoute>
                <CartPage />
              </AuthRoute>
            }
          />

          <Route
            path="user"
            element={
              <AuthRoute>
                <UserLayout />
              </AuthRoute>
            }
          >
            <Route path="orders" element={<Orders />} />
            <Route path="profile" element={<Profile />} />
            <Route path="address" element={<Address />} />
            <Route index element={<Navigate to="profile" replace />} />
          </Route>

          <Route
            path="cart/checkout"
            element={
              <AuthRoute>
                <Checkout />
              </AuthRoute>
            }
          />
          <Route path="products" element={<ProductResults />} />
          <Route path="products/:slug" element={<ProductDetail />} />
        </Route>

        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
}

export default App;
