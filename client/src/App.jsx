import Home from "./pages/Home";
import SignIn from "./pages/SignIn";
import SignUp from "./pages/SignUp";
import NotFound from "./pages/NotFound";
import CartPage from "./pages/CartPage";
import Checkout from "./pages/Checkout";
import ProductDetail from "./pages/ProductDetail";
import ProductResults from "./pages/ProductResults";

import { Toaster } from "sonner";
import { AuthRoute, NonAuthRoute } from "./middleware";
import Layout from "./components/layout/Layout";
import { Route, Routes } from "react-router-dom";

function App() {
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
            path="/cart"
            element={
              <AuthRoute>
                <CartPage />
              </AuthRoute>
            }
          />
          <Route
            path="/cart/checkout"
            element={
              <AuthRoute>
                <Checkout />
              </AuthRoute>
            }
          />
          <Route path="/products" element={<ProductResults />} />
          <Route path="/products/:slug" element={<ProductDetail />} />
        </Route>

        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
}

export default App;
