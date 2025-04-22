import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import { useProductsQuery } from "@/hooks/useProductsQuery";
import ProductCard from "@/components/products/ProductCard";

const Home = () => {
  const {
    isError,
    refetch,
    isLoading,
    data: products = { results: [] },
  } = useProductsQuery();

  if (isLoading) return <FetchLoading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="container min-h-screen mx-auto">
      <div className="px-2 space-y-4 py-3 md:py-6">
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          {products.results.map((product) => (
            <ProductCard key={product.id} product={product} />
          ))}
        </div>
      </div>
    </section>
  );
};

export default Home;
