import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/navigation";
import {
  useProductsQuery,
  useCategoriesQuery,
  useGetAllBannersQuery,
} from "@/hooks/useProductsQuery";
import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import ProductCard from "@/components/products/ProductCard";
import { Swiper, SwiperSlide } from "swiper/react";
import { Autoplay, Pagination, Navigation } from "swiper/modules";

const Home = () => {
  const {
    isError,
    refetch,
    isLoading,
    data: banners = { results: [] },
  } = useGetAllBannersQuery();
  const { data: products = { results: [] } } = useProductsQuery({ limit: 50 });
  const { data: categories = { results: [] } } = useCategoriesQuery({
    limit: 50,
  });

  if (isLoading) return <FetchLoading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  const topBanners = banners?.filter((b) => b.position === "top");
  const side1Banners = banners?.filter((b) => b.position === "side1");
  const side2Banners = banners?.filter((b) => b.position === "side2");
  const bottomBanners = banners?.filter((b) => b.position === "bottom");

  const featuredProducts = products.results.filter((p) => p.isFeatured);
  const discountProducts = products.results.filter((p) => p.discount > 0.05);
  const gadgetProducts = products.results.filter(
    (p) => p.category?.name === "Gadget & Electronics"
  );

  return (
    <section className="py-6 px-4 md:px-6 max-w-screen-xl mx-auto">
      {/* Top Banner */}
      {topBanners.length > 0 && (
        <div className="mb-6">
          <Swiper
            loop
            spaceBetween={10}
            centeredSlides
            autoplay={{ delay: 4000, disableOnInteraction: false }}
            pagination={{ clickable: true }}
            navigation
            modules={[Autoplay, Pagination, Navigation]}
            className="rounded-xl overflow-hidden shadow-lg"
          >
            {topBanners.map((banner) => (
              <SwiperSlide key={banner.id}>
                <img
                  src={banner.imageUrl}
                  alt="top-banner"
                  className="h-64 md:h-96 w-full object-cover"
                />
              </SwiperSlide>
            ))}
          </Swiper>
        </div>
      )}

      {/* Side Banner 1 (Mobile First) */}
      {side1Banners.length > 0 && (
        <div className="mb-6 md:hidden">
          <Swiper
            loop
            spaceBetween={10}
            autoplay={{ delay: 4000, disableOnInteraction: false }}
            pagination={{ clickable: true }}
            modules={[Autoplay, Pagination]}
            className="rounded-xl overflow-hidden shadow-md"
          >
            {side1Banners.map((banner) => (
              <SwiperSlide key={banner.id}>
                <img
                  src={banner.imageUrl}
                  alt="side1-banner"
                  className="h-48 w-full object-cover"
                />
              </SwiperSlide>
            ))}
          </Swiper>
        </div>
      )}

      {/* Featured Products */}
      <section className="mb-8">
        <h2 className="text-xl font-semibold mb-2">✨ Featured Products</h2>
        <p className="text-sm text-muted-foreground mb-4">
          Don’t miss out on the latest trends and deals!
        </p>
        <Swiper
          spaceBetween={16}
          navigation
          modules={[Navigation]}
          breakpoints={{
            0: { slidesPerView: 1 },
            640: { slidesPerView: 2 },
            768: { slidesPerView: 3 },
            1024: { slidesPerView: 4 },
          }}
        >
          {featuredProducts.map((product) => (
            <SwiperSlide key={product.id}>
              <ProductCard product={product} />
            </SwiperSlide>
          ))}
        </Swiper>
      </section>

      {/* Discount Products */}
      <section className="mb-8">
        <h2 className="text-xl font-semibold mb-2">🔥 Discounted Products</h2>
        <p className="text-sm text-muted-foreground mb-4">
          Grab these deals while they last!
        </p>
        <Swiper
          spaceBetween={16}
          navigation
          modules={[Navigation]}
          breakpoints={{
            0: { slidesPerView: 1 },
            640: { slidesPerView: 2 },
            768: { slidesPerView: 3 },
            1024: { slidesPerView: 4 },
          }}
        >
          {discountProducts.map((product) => (
            <SwiperSlide key={product.id}>
              <ProductCard product={product} />
            </SwiperSlide>
          ))}
        </Swiper>
      </section>

      {/* Side Banner 2 (Mobile) */}
      {side2Banners.length > 0 && (
        <div className="mb-6 md:hidden">
          <Swiper
            loop
            spaceBetween={10}
            autoplay={{ delay: 4000 }}
            pagination={{ clickable: true }}
            modules={[Autoplay, Pagination]}
            className="rounded-xl overflow-hidden shadow-md"
          >
            {side2Banners.map((banner) => (
              <SwiperSlide key={banner.id}>
                <img
                  src={banner.imageUrl}
                  alt="side2-banner"
                  className="h-48 w-full object-cover"
                />
              </SwiperSlide>
            ))}
          </Swiper>
        </div>
      )}

      {/* Gadget & Electronics */}
      <section className="mb-8">
        <h2 className="text-xl font-semibold mb-2">📱 Gadget & Electronics</h2>
        <p className="text-sm text-muted-foreground mb-4">
          Discover the latest tech at great prices.
        </p>
        <Swiper
          spaceBetween={16}
          navigation
          modules={[Navigation]}
          breakpoints={{
            0: { slidesPerView: 1 },
            640: { slidesPerView: 2 },
            768: { slidesPerView: 3 },
            1024: { slidesPerView: 4 },
          }}
        >
          {gadgetProducts.map((product) => (
            <SwiperSlide key={product.id}>
              <ProductCard product={product} />
            </SwiperSlide>
          ))}
        </Swiper>
      </section>

      {/* Bottom Banners */}
      {bottomBanners.length > 0 && (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {bottomBanners.map((banner) => (
            <div key={banner.id} className="rounded-xl overflow-hidden">
              <img
                src={banner.imageUrl}
                alt="bottom-banner"
                className="h-56 md:h-64 w-full object-cover"
              />
            </div>
          ))}
        </div>
      )}
    </section>
  );
};

export default Home;
