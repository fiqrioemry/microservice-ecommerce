// 1. hooks/useCategoriesQuery.jsx
import { useQuery } from "@tanstack/react-query";
import categoryApi from "@/api/category";

export function useCategoriesQuery() {
  return useQuery({
    queryKey: ["categories"],
    queryFn: categoryApi.getAllCategories,
    staleTime: 1000 * 60 * 5,
    retry: 1,
  });
}
