import products from "@/services/products";
import { useQuery } from "@tanstack/react-query";

export const useProductsQuery = (params = {}) =>
  useQuery({
    queryKey: ["products", params],
    queryFn: () => products.getAllProducts(params),
    staleTime: 1000 * 60 * 5,
  });

export const useProductDetailQuery = (slug) =>
  useQuery({
    queryKey: ["product", slug],
    queryFn: () => products.getProductBySlug(slug),
    enabled: !!slug,
  });

export const useSearchProductsQuery = (queryParams) =>
  useQuery({
    queryKey: ["products", "search", queryParams],
    queryFn: () => products.searchProducts(queryParams),
    enabled: !!queryParams,
  });
