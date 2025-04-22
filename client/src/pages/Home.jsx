import { useEffect } from "react";
import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import { useProductsQuery } from "@/hooks/useProductsQuery";

const Home = () => {
  const {
    isError,
    refetch,
    isLoading,
    data: products = { results: [] },
  } = useProductsQuery();
  console.log(products.results);

  if (isLoading) return <FetchLoading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

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
