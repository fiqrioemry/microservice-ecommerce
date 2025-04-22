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
    <section className="container min-h-screen mx-auto">
      <div className="px-2 space-y-4 py-3 md:py-6">
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          {products.results.map((product) => (
            <div
              key={product.id}
              className="bg-white rounded-xl shadow-sm border hover:shadow-md transition duration-300"
            >
              <div className="relative w-full aspect-square overflow-hidden rounded-t-xl">
                <img
                  src={product.images?.[0]}
                  alt={product.name}
                  className="w-full h-full object-cover"
                />
                {product.isFeatured && (
                  <span className="absolute top-2 left-2 text-xs font-semibold bg-yellow-400 text-black px-2 py-1 rounded">
                    ⭐ Featured
                  </span>
                )}
              </div>

              <div className="p-4 space-y-1">
                <h2 className="text-base font-semibold text-gray-900">
                  {product.name}
                </h2>
                <p className="text-xs text-muted-foreground">
                  {product.subcategory?.name} — {product.category?.name}
                </p>
                <p className="text-sm text-gray-500 line-clamp-2">
                  {product.description}
                </p>
                <div className="flex justify-between items-center mt-3">
                  <span className="text-primary font-bold text-sm">
                    ${product.price.toFixed(2)}
                  </span>
                  <button className="text-sm font-medium text-blue-600 hover:underline">
                    Lihat Detail
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

export default Home;
