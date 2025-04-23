import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import { useBannersQuery } from "@/hooks/useProductsQuery";

const Banners = () => {
  const {
    isError,
    refetch,
    isLoading,
    data: banners = { results: [] },
  } = useBannersQuery("top");

  if (isLoading) return <FetchLoading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="container mx-auto">
      <div className="px-2 space-y-4 py-3 md:py-6">
        <img src={banners[0].imageUrl} alt={banners[0].position} />
      </div>
    </section>
  );
};

export default Banners;
