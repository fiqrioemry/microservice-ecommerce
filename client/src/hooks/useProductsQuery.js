import product from "@/services/products";
import banner from "@/services/banners";
import category from "@/services/categories";
import { useQuery } from "@tanstack/react-query";

export const useProductsQuery = (params = {}) =>
  useQuery({
    queryKey: ["products", params],
    queryFn: () => product.getAllProducts(params),
    keepPreviousData: true,
  });

export const useProductDetailQuery = (slug) =>
  useQuery({
    queryKey: ["product", slug],
    queryFn: () => product.getProductBySlug(slug),
    enabled: !!slug,
  });

export const useSearchProductsQuery = (queryParams) =>
  useQuery({
    queryKey: ["products", "search", queryParams],
    queryFn: () => product.searchProducts(queryParams),
    enabled: !!queryParams,
  });

export const useCategoriesQuery = () =>
  useQuery({
    queryKey: ["categories"],
    queryFn: category.getAllCategories,
  });

export const useGetAllBannersQuery = () =>
  useQuery({
    queryKey: ["banners"],
    queryFn: banner.getAllBanners,
  });

export const useGetBannerByPositionQuery = (position) =>
  useQuery({
    queryKey: ["banners", position],
    queryFn: () => banner.getBannersByPosition(position),
    enabled: !!position,
  });
