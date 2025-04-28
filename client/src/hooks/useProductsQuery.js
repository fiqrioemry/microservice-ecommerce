import product from "@/services/products";
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
