import { useEffect } from "react";
import { useProductsQuery } from "../hooks/useProductsQuery";

const Home = () => {
  const { data: products = [], isLoading, isError } = useProductsQuery();

  if (isLoading) return <div className="text-center py-10">Loading...</div>;
  if (isError)
    return (
      <div className="text-center py-10 text-red-500">Gagal memuat produk.</div>
    );

  return (
    <section className="container h-screen mx-auto">
      <div className="px-2 space-y-4 py-3 md:py-6">
        <div className="grid grid-cols-4 gap-4">
          {products.results.map((product) => (
            <div className="col-span-1 rounded-lg" key={product.id}>
              <img src={product.images[0]} alt={product.name} />
              <h2 className="text-lg font-semibold">{product.name}</h2>
              <p className="text-gray-600">{product.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

export default Home;
