import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/navigation";
import {
  useProductsQuery,
  useCategoriesQuery,
  useGetAllBannersQuery,
} from "@/hooks/useProductsQuery";
import { Swiper, SwiperSlide } from "swiper/react";
import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import ProductCard from "@/components/product-results/ProductCard";
import { Autoplay, Pagination, Navigation } from "swiper/modules";

const Home = () => {
  const {
    isError,
    refetch,
    isLoading,
    data: banners = { results: [] },
  } = useGetAllBannersQuery();
  const { data: products = { results: [] } } = useProductsQuery({ limit: 50 });
  const { data: { categories } = { results: [] } } = useCategoriesQuery({
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
    <section className="section">
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
      {/* Browse by Category */}
      {categories.length > 0 && (
        <section className="mb-8">
          <h2 className="text-lg font-semibold mb-4">🛍️ Browse by Category</h2>
          <div className="flex gap-4 overflow-x-auto pb-2 no-scrollbar">
            {categories.map((cat) => (
              <div
                key={cat.ID}
                onClick={() =>
                  (window.location.href = `/products?category=${cat.slug}`)
                }
                className="min-w-[140px] cursor-pointer rounded-xl border shadow hover:shadow-md transition bg-background"
              >
                <img
                  src={cat.image}
                  alt={cat.name}
                  className="h-24 w-full object-cover rounded-t-xl"
                />
                <div className="p-2 text-center text-sm font-medium">
                  {cat.name}
                </div>
              </div>
            ))}
          </div>
        </section>
      )}

      {/* Side Banner 1 (mobile) */}
      {side1Banners.length > 0 && (
        <div className="md:hidden mb-6">
          <Swiper
            loop
            spaceBetween={10}
            autoplay={{ delay: 4000 }}
            pagination={{ clickable: true }}
            modules={[Autoplay, Pagination]}
            className="rounded-xl overflow-hidden shadow-md"
          >
            {side1Banners.map((banner) => (
              <SwiperSlide key={banner.id}>
                <img
                  src={banner.imageUrl}
                  alt="side1-mobile"
                  className="h-48 w-full object-cover"
                />
              </SwiperSlide>
            ))}
          </Swiper>
        </div>
      )}

      {/* Grid for Side1 - Products - Side2 */}
      <div className="grid grid-cols-1 md:grid-cols-12 gap-6 mb-8">
        {/* Side2 desktop */}
        {side2Banners.length > 0 && (
          <div className="hidden md:block md:col-span-2 space-y-4">
            {side2Banners.map((banner) => (
              <img
                key={banner.id}
                src={banner.imageUrl}
                alt="side2"
                className="rounded-xl object-cover w-full shadow"
              />
            ))}
          </div>
        )}
        {/* Main Content: Featured & Discount */}
        <div className="md:col-span-8 space-y-8">
          {/* Featured Products */}
          <section>
            <h2 className="text-xl font-semibold mb-2">✨ Featured Products</h2>
            <p className="text-sm text-muted-foreground mb-4">
              Do not miss the featured deals! It Limited time sales.
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
          <section>
            <h2 className="text-xl font-semibold mb-2">
              🔥 Discounted Products
            </h2>
            <p className="text-sm text-muted-foreground mb-4">
              Limited time discounts! Get our product with most valuable price.
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
        </div>

        {side1Banners.length > 0 && (
          <div className="hidden md:block md:col-span-2 space-y-4">
            {side1Banners.map((banner) => (
              <img
                key={banner.id}
                src={banner.imageUrl}
                alt="side1"
                className="rounded-xl object-cover w-full shadow"
              />
            ))}
          </div>
        )}
      </div>

      {/* Side Banner 2 (mobile) */}
      {side2Banners.length > 0 && (
        <div className="md:hidden mb-6">
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
                  alt="side2-mobile"
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
            <img
              key={banner.id}
              src={banner.imageUrl}
              alt="bottom-banner"
              className="rounded-xl object-cover h-56 md:h-64 w-full shadow"
            />
          ))}
        </div>
      )}
    </section>
  );
};

export default Home;
