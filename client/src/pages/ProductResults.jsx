import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import { useProductsQuery } from "@/hooks/useProductsQuery";
import ProductCard from "@/components/products/ProductCard";

const ProductResults = () => {
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
      <div className="px-2 space-y-4 py-3 md:py-6"></div>
    </section>
  );
};

export default ProductResults;
