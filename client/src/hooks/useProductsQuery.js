import p from "@/services/products";
import c from "@/services/categories";
import b from "@/services/banners";
import { useQuery } from "@tanstack/react-query";

export const useProductsQuery = (params = {}) =>
  useQuery({
    queryKey: ["products", params],
    queryFn: () => p.getAllProducts(params),
  });

export const useProductDetailQuery = (slug) =>
  useQuery({
    queryKey: ["product", slug],
    queryFn: () => p.getProductBySlug(slug),
    enabled: !!slug,
  });

export const useSearchProductsQuery = (queryParams) =>
  useQuery({
    queryKey: ["products", "search", queryParams],
    queryFn: () => p.searchProducts(queryParams),
    enabled: !!queryParams,
  });

export const useCategoriesQuery = () =>
  useQuery({
    queryKey: ["categories"],
    queryFn: c.getAllCategories,
  });

export const useBannersQuery = (position) => console.log("get banner");
useQuery({
  queryKey: ["banners", position],
  queryFn: () => b.getBanners(position),
  enabled: !!position,
});
