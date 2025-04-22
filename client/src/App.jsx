import Home from "./pages/Home";
import NotFound from "./pages/NotFound";

// import { Toaster } from "react-hot-toast";
import { Toaster } from "sonner";
import Layout from "./components/layout/Layout";
import { Route, Routes } from "react-router-dom";
import ProductDetail from "./pages/ProductDetail";

function App() {
  return (
    <>
      <Toaster />
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="/products/:slug" element={<ProductDetail />} />
        </Route>

        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
}

export default App;
