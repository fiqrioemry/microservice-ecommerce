import { Swiper, SwiperSlide } from "swiper/react";
import { Autoplay, Pagination, Navigation } from "swiper/modules";
import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/navigation";

import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import {
  useGetAllBannersQuery,
  useProductsQuery,
} from "@/hooks/useProductsQuery";

const Banners = () => {
  const {
    isError,
    refetch,
    isLoading,
    data: banners = { results: [] },
  } = useGetAllBannersQuery();

  if (isLoading) return <FetchLoading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className=" mx-auto py-3 md:py-6">
      <div className="container mx-auto">
        <Swiper
          loop={true}
          spaceBetween={10}
          centeredSlides={true}
          autoplay={{
            delay: 4000,
            disableOnInteraction: false,
          }}
          pagination={{ clickable: true }}
          navigation={true}
          modules={[Autoplay, Pagination, Navigation]}
          className="rounded-xl overflow-hidden shadow-lg"
        >
          {banners?.map((banner) => (
            <SwiperSlide key={banner.id}>
              <div>
                <img
                  src={banner.imageUrl}
                  alt={banner.position}
                  className="h-96 w-full object-center rounded-xl"
                />
              </div>
            </SwiperSlide>
          ))}
        </Swiper>
        <div className="grid grid-cols-5 mt-4 gap-4">
          <div className="col-span-1 h-60 bg-blue-500 rounded-lg"></div>
          <div className="col-span-1 h-60 bg-blue-500 rounded-lg"></div>
          <div className="col-span-1 h-60 bg-blue-500 rounded-lg"></div>
          <div className="col-span-1 h-60 bg-blue-500 rounded-lg"></div>
          <div className="col-span-1 h-60 bg-blue-500 rounded-lg"></div>
        </div>

        <div className="grid grid-cols-5 mt-4 gap-4">
          <div className="col-span-1">
            <div className="h-80 w-full bg-red-500 border rounded-lg mb-4">
              sidebanner
            </div>
            <div className="h-80 w-full bg-blue-500 border rounded-lg mb-4">
              sidebanner
            </div>
          </div>
          <div className="col-span-4">
            <div className="h-80 mb-4">
              <h2>Popular Products</h2>
              <p> Do not miss the current offers until the end of March.</p>
              <div>product to show</div>
            </div>
            <div className="h-80 mb-4">
              <h2>Featured Products</h2>
              <p> Do not miss the current offers until the end of March.</p>
              <div className="h-80">product to show</div>
            </div>
          </div>
        </div>

        <div className="grid grid-cols-5 mt-4 mb-4 gap-4">
          <div className="col-span-1 h-60 bg-blue-500 rounded-lg"></div>
          <div className="col-span-1 h-60 bg-blue-500 rounded-lg"></div>
          <div className="col-span-1 h-60 bg-blue-500 rounded-lg"></div>
          <div className="col-span-1 h-60 bg-blue-500 rounded-lg"></div>
          <div className="col-span-1 h-60 bg-blue-500 rounded-lg"></div>
        </div>
      </div>
      <div className="h-80 bg-blue-500 rounded-lg"></div>
    </section>
  );
};

export default Banners;
